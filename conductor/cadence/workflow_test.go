package cadence

import (
	"fmt"
	"testing"
)

func TestTransferToGJsonFormat(t *testing.T) {
	sstr := transferToGJsonFormat("get_skeleton.output.response.body.data[0].result..way_name")
	fmt.Sprintf(sstr)
}
