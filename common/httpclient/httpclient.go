package httpclient

import (
	"bytes"
	"fmt"
	"github.com/prometheus/common/log"
	"io/ioutil"
	"net/http"
	"net/url"
	"strings"
)

// HTTPClient the encapsulation of http
type HTTPClient struct {
	BaseURL   string
	Headers   map[string]string
	PrintLogs bool
}

// NewHTTPClient create a http client
func NewHTTPClient(baseURL string, headers map[string]string, printLogs bool) *HTTPClient {
	httpClient := new(HTTPClient)
	httpClient.BaseURL = baseURL
	httpClient.Headers = headers
	httpClient.PrintLogs = printLogs
	return httpClient
}

func (c *HTTPClient) logSendRequest(url string, requestType string, body string) {
	log.Info("Sending [", requestType, "] request to Server (", url, "):")
	log.Info("Body:", body)
}

func (c *HTTPClient) logResponse(statusCode string, response string) {
	log.Info("Received response from Server (", c.BaseURL, "):")
	log.Info("Status", statusCode,
		"Response", response)
}

func genParamString(paramMap map[string]string) string {
	if paramMap == nil || len(paramMap) == 0 {
		return ""
	}

	output := "?"
	for key, value := range paramMap {
		output = fmt.Sprintf("%s%s=%s&", output, key, value)
	}
	return output
}

func (c *HTTPClient) httpRequest(url string, requestType string, headers map[string]string, body string) (string, error) {
	var (
		req *http.Request
		err error
	)
	if requestType == "GET" {
		req, err = http.NewRequest(requestType, url, nil)
	} else {
		bodyStr := []byte(body)
		req, err = http.NewRequest(requestType, url, bytes.NewBuffer(bodyStr))
	}
	if err != nil {
		return "", nil
	}

	// Default header
	for key, value := range c.Headers {
		req.Header.Set(key, value)
	}
	// Custom header
	for key, value := range headers {
		req.Header.Set(key, value)
	}

	if c.PrintLogs {
		c.logSendRequest(url, requestType, body)
	}

	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		return "", err
	}

	defer resp.Body.Close()
	response, err := ioutil.ReadAll(resp.Body)
	responseString := string(response)
	if err != nil {
		log.Error("ERROR reading response for URL: ")
	}

	// judge whether response code is larger than 400
	if resp.StatusCode >= 400 {
		c.logResponse(resp.Status, responseString)
		return "", fmt.Errorf("return status code is larger than 400, the response String is: %s", responseString)
	}
	// print response infomation
	if c.PrintLogs {
		c.logResponse(resp.Status, responseString)
	}

	return responseString, nil
}

// getEncodeURL 对url和params进行url编码
func getEncodeURL(urlString string, queryParamsMap map[string]string) (string, error) {
	baseURL, err := url.Parse(urlString)
	if err != nil {
		return "", err
	}
	params := url.Values{}
	for key, value := range queryParamsMap {
		params.Add(key, value)
	}
	baseURL.RawQuery = params.Encode()
	return baseURL.String(), nil
}

// Get HTTP GET method
func (c *HTTPClient) Get(url string, queryParamsMap map[string]string, headers map[string]string) (string, error) {
	urlString, err := getEncodeURL(url, queryParamsMap)
	if err != nil {
		log.Error(fmt.Sprintf("Encode baseURL error, baseurl: %s, reason: %s", url, err.Error()))
		return "", err
	}
	resp, err := c.httpRequest(urlString, "GET", headers, "")
	if err != nil {
		log.Error(fmt.Sprintf("Http Get Error for URL %s", urlString))
		return "", err
	}
	return resp, nil
}

// Post HTTP POST method
func (c *HTTPClient) Post(url string, queryParamsMap map[string]string, headers map[string]string, body string) (string, error) {
	urlString := url + genParamString(queryParamsMap)
	resp, err := c.httpRequest(urlString, "POST", headers, body)
	if err != nil {
		log.Error(fmt.Sprintf("HTTP POST Error for URL %s", urlString))
		return "", err
	}
	return resp, nil
}

// Put HTTP POST method
func (c *HTTPClient) Put(url string, queryParamsMap map[string]string, headers map[string]string, body string) (string, error) {
	urlString := url + genParamString(queryParamsMap)
	resp, err := c.httpRequest(urlString, "PUT", headers, body)
	if err != nil {
		log.Error(fmt.Sprintf("HTTP PUT Error for URL %s", urlString))
		return "", err
	}
	return resp, nil
}

// MakeURL Make the url using given paramaters
func (c *HTTPClient) MakeURL(path string, args ...string) string {
	url := c.BaseURL
	r := strings.NewReplacer(args...)
	return url + r.Replace(path)
}
