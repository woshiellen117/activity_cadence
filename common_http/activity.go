package main

import (
	"encoding/json"
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

func CommonHttp(input []byte) ([]byte, error) {
	var httpData HttpRequestStruct
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
		return []byte(response), nil
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
		return []byte(response), nil
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
			response, err := httpClient.Get(baseURL, params, headers)
			if err != nil {
				return nil, err
			}
			return []byte(response), nil
		} else {
			response, err := httpClient.Get(url, nil, headers)
			if err != nil {
				return nil, err
			}
			return []byte(response), nil
		}
	}
	return nil, nil
}
