package main

import (
	"encoding/json"
)

func iw_mapping(input []byte) ([]byte, error) {
	var inputData interface{}
	if err := json.Unmarshal(input, &inputData); err != nil {
		return nil, err
	}
	jsonReturn := `{"coverages_ok":"NO"}`
	return []byte(jsonReturn), nil
}
