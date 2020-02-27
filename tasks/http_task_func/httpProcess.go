package http_task_func

import (
	"encoding/json"
	"fmt"
	"go.momenta.works/activity_cadence/common/httpclient"
	"strings"
)

type HttpRequestStruct struct {
	HttpRequest HttpRequest `json:"http_request"`
}
type HttpRequest struct {
	Body        interface{} `json:"body"`
	ContentType string      `json:"contentType"`
	Method      string      `json:"method"`
	Uri         string      `json:"uri"`
}

func HttpProcess(input []byte) ([]byte, error) {

	var httpData HttpRequestStruct
	fmt.Print(string(input))
	if err := json.Unmarshal(input, &httpData); err != nil {
		return nil, err
	}
	url := httpData.HttpRequest.Uri
	headers := map[string]string{"Content-Type": httpData.HttpRequest.ContentType}
	httpClient := httpclient.NewHTTPClient(url, headers, true)
	if httpData.HttpRequest.Method == "PUT" {
		httpBody, err := json.Marshal(httpData.HttpRequest.Body)
		if err != nil {
			return nil, err
		}
		response, err := httpClient.Put(url, nil, headers, string(httpBody))
		if err != nil {
			return nil, err
		}
		if response == "" {
			return nil, nil
		}
		returnJSON, err := packageResult(response)
		if err != nil {
			return nil, err
		}
		return returnJSON, nil
	}
	if httpData.HttpRequest.Method == "POST" {
		httpBody, err := json.Marshal(httpData.HttpRequest.Body)
		if err != nil {
			return nil, err
		}
		response, err := httpClient.Post(url, nil, headers, string(httpBody))
		if err != nil {
			return nil, err
		}
		if response == "" {
			return nil, nil
		}
		returnJSON, err := packageResult(response)
		if err != nil {
			return nil, err
		}
		return returnJSON, nil
	}
	if httpData.HttpRequest.Method == "GET" {
		if strings.Contains(url, "?") {
			urlAndParamArray := strings.Split(url, "?")
			baseURL := urlAndParamArray[0]
			paramArray := strings.Split(urlAndParamArray[1], "&")
			params := make(map[string]string)
			for _, param := range paramArray {
				kv := strings.Split(param, "=")
				params[kv[0]] = kv[1]
			}
			response, err := httpClient.Get(baseURL, params, nil)
			if err != nil {
				return nil, err
			}
			if response == "" {
				return nil, nil
			}
			returnJSON, err := packageResult(response)
			if err != nil {
				return nil, err
			}
			return returnJSON, nil
		} else {
			response, err := httpClient.Get(url, nil, headers)
			if err != nil {
				return nil, err
			}
			if response == "" {
				return nil, nil
			}
			returnJSON, err := packageResult(response)
			if err != nil {
				return nil, err
			}
			return returnJSON, nil
		}
	}
	return nil, nil
}
func packageResult(response string) ([]byte, error) {
	var responseJSON interface{}
	if err := json.Unmarshal([]byte(response), &responseJSON); err != nil {
		return nil, nil
	}
	bodyMap := make(map[string]interface{})
	bodyMap["body"] = responseJSON
	responseMap := make(map[string]interface{})
	responseMap["response"] = bodyMap

	returnJSON, err := json.Marshal(responseMap)
	if err != nil {
		return nil, err
	}
	return returnJSON, nil
}
