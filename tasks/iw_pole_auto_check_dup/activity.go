package main

import "encoding/json"

func iw_pole_auto_check_dup(input []byte) ([]byte, error) {
	mapReturn := make(map[string]interface{})
	mapReturn["status"] = "YES"
	strJSON, err := json.Marshal(mapReturn)
	if err != nil {
		return nil, err
	}
	return strJSON, nil
}
