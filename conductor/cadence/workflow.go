package cadence

import (
	"encoding/json"
	"errors"
	"fmt"
	"github.com/tidwall/gjson"
	"go.momenta.works/activity_cadence/conductor/model"
	"go.uber.org/cadence/workflow"
	"go.uber.org/zap"
	"strconv"
	"strings"
)

// ApplicationName is the task list for this sample
//const ApplicationName = "myDslGroup"

type (
	// Workflow is the type used to express the workflow definition. Variables are a map of valuables. Variables can be
	// used as input to Activity.
	Workflow struct {
		Variables map[string]interface{}
		Root      Statement
	}

	// Statement is the building block of dsl workflow. A Statement can be a simple ActivityInvocation or it
	// could be a Sequence or Parallel.
	Statement struct {
		Activity *ActivityInvocation
		Sequence *Sequence
		Parallel *Parallel
		Choice   *Choice
	}

	// Sequence consist of a collection of Statements that runs in sequential.
	Sequence struct {
		Elements []*Statement
	}

	// Parallel can be a collection of Statements that runs in parallel.
	Parallel struct {
		Branches []*Statement
	}

	// Choice can make choice using the last output
	Choice struct {
		Keybase string

		Name        string
		TaskRefName string
		Arguments   map[string]interface{}
		Branches    map[string]Statement
	}

	// SingleChoice
	//SingleChoice struct{
	//	SingleChoice map[string]Statement
	//}

	//SingleChoiceStruct struct{
	//	Key string
	//	Value *Statement
	//}

	// ActivityInvocation is used to express invoking an Activity. The Arguments defined expected arguments as input to
	// the Activity, the result specify the name of variable that it will store the result as which can then be used as
	// arguments to subsequent ActivityInvocation.
	ActivityInvocation struct {
		Name        string
		TaskRefName string
		Arguments   map[string]interface{}
		Result      string
	}

	Executable interface {
		Execute(ctx workflow.Context, ao model.ActivityOptions, bindings map[string]interface{}) error
	}
)

func (b *Statement) Execute(ctx workflow.Context, ao model.ActivityOptions, bindings map[string]interface{}) error {
	if b.Parallel != nil {
		fmt.Printf(">>>>>>>Parallel")
		err := b.Parallel.Execute(ctx, ao, bindings)
		if err != nil {
			return err
		}
	}
	if b.Sequence != nil {
		fmt.Printf(">>>>>>>Sequence")
		err := b.Sequence.Execute(ctx, ao, bindings)
		if err != nil {
			return err
		}
	}
	if b.Activity != nil {
		fmt.Printf(">>>>>>>Activity")
		err := b.Activity.Execute(ctx, ao, bindings)
		if err != nil {
			return err
		}
	}
	if b.Choice != nil {
		fmt.Printf(">>>>>>>Choice")
		err := b.Choice.Execute(ctx, ao, bindings)
		if err != nil {
			return err
		}
	}
	return nil
}

func (p Parallel) Execute(ctx workflow.Context, ao model.ActivityOptions, bindings map[string]interface{}) error {
	//
	// You can use the context passed in to activity as a way to cancel the activity like standard GO way.
	// Cancelling a parent context will cancel all the derived contexts as well.
	//

	// In the parallel block, we want to execute all of them in parallel and wait for all of them.
	// if one activity fails then we want to cancel all the rest of them as well.
	childCtx, cancelHandler := workflow.WithCancel(ctx)
	selector := workflow.NewSelector(ctx)
	var activityErr error
	for _, s := range p.Branches {
		f := ExecuteAsync(s, childCtx, ao, bindings)
		selector.AddFuture(f, func(f workflow.Future) {
			err := f.Get(ctx, nil)
			if err != nil {
				// cancel all pending activities
				cancelHandler()
				activityErr = err
			}
		})
	}

	for i := 0; i < len(p.Branches); i++ {
		selector.Select(ctx) // this will wait for one branch
		if activityErr != nil {
			return activityErr
		}
	}
	return nil
}

func ExecuteAsync(exe Executable, ctx workflow.Context, ao model.ActivityOptions, bindings map[string]interface{}) workflow.Future {
	future, settable := workflow.NewFuture(ctx)
	workflow.Go(ctx, func(ctx workflow.Context) {
		err := exe.Execute(ctx, ao, bindings)
		settable.Set(nil, err)
	})
	return future
}

func (s Sequence) Execute(ctx workflow.Context, ao model.ActivityOptions, bindings map[string]interface{}) error {
	for _, a := range s.Elements {
		err := a.Execute(ctx, ao, bindings)
		if err != nil {
			return err
		}
	}
	return nil
}

func (a ActivityInvocation) Execute(ctx workflow.Context, ao model.ActivityOptions, bindings map[string]interface{}) error {
	inputParam, err := makeInput(a.Arguments, bindings)
	if err != nil {
		return err
	}
	var result []byte
	activityOptions := workflow.ActivityOptions{
		ScheduleToStartTimeout: ao.ScheduleToStartTimeout,
		StartToCloseTimeout:    ao.StartToCloseTimeout,
		HeartbeatTimeout:       ao.HeartbeatTimeout,
	}
	activityOptions.TaskList = a.Name
	ctx = workflow.WithActivityOptions(ctx, activityOptions)

	if a.Name == "wait_complete" {
		var signalVal string
		signalChan := workflow.GetSignalChannel(ctx, a.TaskRefName)

		s := workflow.NewSelector(ctx)
		s.AddReceive(signalChan, func(c workflow.Channel, more bool) {
			c.Receive(ctx, &signalVal)
			workflow.GetLogger(ctx).Info("Received signal!", zap.String("signal", a.TaskRefName), zap.String("value", signalVal))
		})
		s.Select(ctx)

		if len(signalVal) > 0 && signalVal != "SOME_VALUE" {
			return errors.New("signalVal")
		}
		return nil
	} else {
		err = workflow.ExecuteActivity(ctx, a.Name, inputParam).Get(ctx, &result)
		if err != nil {
			return err
		}
		if len(result) == 0 {
			return nil
		}
		fmt.Printf(">>>After activity result:%s", result)
		var returnData interface{}
		if err := json.Unmarshal(result, &returnData); err != nil {
			return err
		}

		output := make(map[string]interface{})
		output["output"] = returnData

		bindings[a.TaskRefName] = output
		return nil
	}

}

func (c Choice) Execute(ctx workflow.Context, ao model.ActivityOptions, bindings map[string]interface{}) error {
	inputParam, err := makeInput(c.Arguments, bindings)
	if err != nil {
		return err
	}
	var inputJSON interface{}
	err = json.Unmarshal(inputParam, &inputJSON)
	if err != nil {
		return err
	}
	var key string
	mapValue, mapOk := inputJSON.(map[string]interface{})
	if mapOk {
		for k, v := range mapValue {
			if k == c.Keybase {
				strValue, strOk := v.(string)
				if strOk {
					key = strValue
				}
			}
		}
	}
	for k, v := range c.Branches {
		if k == key {
			err := v.Execute(ctx, ao, bindings)
			if err != nil {
				return err
			}
		}
	}
	return nil
}

// 输入workflow Def文件中inputParameters的内容，替换成bindings中存的实际值
// 1、判断 interface是不是 string
//      1.1 是string 判断是否存在$符号
//			1.1.1 存在$：转换成bindings里相应的数据，使用gjson取得
//          1.1.2 不存在$：直接赋值返回
//      1.2 不是string 递归判断
func makeInput(inputParameters map[string]interface{}, bindings map[string]interface{}) ([]byte, error) {
	args, err := transferInputsToParams(inputParameters, bindings)
	if err != nil {
		return nil, err
	}
	argsByte, err := json.Marshal(args)
	return argsByte, nil
}

//args 返回的数据
//mapString inputParameters
//bindings bindings
func transferInputsToParams(mapString map[string]interface{}, bindings map[string]interface{}) (interface{}, error) {
	bindingsJsonStr, err := json.Marshal(bindings)
	if err != nil {
		return nil, err
	}
	args := make(map[string]interface{})
	for k, v := range mapString {
		stringValue, stringOk := v.(string)
		if stringOk { //string
			if strings.Contains(stringValue, "$") { //存在变量占位替换符$
				dollarCount := strings.Count(stringValue, "$")
				if dollarCount > 1 { //如果某字符串中存在两个$，默认不为 interface
					for i := 0; i < dollarCount; i++ {
						dollarIndex := strings.Index(stringValue, "$")

						before := stringValue[0:dollarIndex]
						middle := stringValue[dollarIndex+2 : strings.Index(stringValue, "}")]
						after := stringValue[strings.Index(stringValue, "}")+1:]

						bindingSearchingString := transferToGJsonFormat(middle)
						bindingValue := gjson.GetBytes(bindingsJsonStr, bindingSearchingString)

						if bindingValue.Str != "" {
							stringValue = before + bindingValue.Str + after
						} else if bindingValue.Num != 0 {
							floatString := strconv.FormatFloat(bindingValue.Num, 'f', -1, 64)
							stringValue = before + floatString + after
						}
					}
					args[k] = stringValue
				} else {
					dollarIndex := strings.IndexAny(stringValue, "$")
					//考虑占位前后有其他字符的情况
					before := stringValue[0:dollarIndex]
					middle := stringValue[dollarIndex+2 : strings.IndexAny(stringValue, "}")]
					after := stringValue[strings.IndexAny(stringValue, "}")+1:]
					bindingSearchString := transferToGJsonFormat(middle)
					bindingValue := gjson.GetBytes(bindingsJsonStr, bindingSearchString)

					if bindingValue.Type == gjson.String {
						args[k] = before + bindingValue.Str + after
					} else if bindingValue.Type == gjson.Number {
						if before == "" && after == "" {
							args[k] = bindingValue.Num
						} else {
							floatValue := strconv.FormatFloat(bindingValue.Num, 'f', -1, 64)
							args[k] = before + floatValue + after
						}

					} else if bindingValue.Type == gjson.JSON {
						var jsonRawValue interface{}
						if err := json.Unmarshal([]byte(bindingValue.Raw), &jsonRawValue); err != nil {
							return nil, err
						}
						args[k] = jsonRawValue
					} else if bindingValue.Type == gjson.True {
						args[k] = true
					} else if bindingValue.Type == gjson.False {
						args[k] = false
					}
				}
			} else {
				args[k] = v
			}
		}

		floatValue, floatOk := v.(float64)
		if floatOk { //float
			args[k] = floatValue
		}

		interfaceValue, interfaceOk := v.(map[string]interface{})
		if interfaceOk { //interface
			argsChild, err := transferInputsToParams(interfaceValue, bindings)
			if err != nil {
				return nil, err
			}
			args[k] = argsChild
		}

	}

	return args, nil
}

//1、将slice形式的result[0]--->result.0.
//2、将循环取用的..--->.#.
func transferToGJsonFormat(str string) string {
	strArray := strings.Split(str, ".")
	var newStringArray []string
	for _, strItem := range strArray {
		var newStrItem string
		if strItem == "" {
			newStrItem = "#"
		} else if strings.Contains(strItem, "[") {
			leftBracket := strings.IndexAny(strItem, "[")
			rightBracket := strings.IndexAny(strItem, "]")
			newStrItem = fmt.Sprintf("%s.%s", strItem[:leftBracket], strItem[leftBracket+1:rightBracket])
		} else {
			newStrItem = strItem
		}
		newStringArray = append(newStringArray, newStrItem)
	}
	return strings.Join(newStringArray, ".")
}
