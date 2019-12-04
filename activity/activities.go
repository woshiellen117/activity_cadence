package activity

func Save_to_grid_store(input interface{}) (interface{}, error) {
	result := `{"response": {
               "headers": {
                  "Server": [
                     "nginx/1.13.12"
                  ],
                  "Access-Control-Allow-Origin": [
                     "*"
                  ],
                  "Connection": [
                     "keep-alive"
                  ],
                  "Content-Length": [
                     "11"
                  ],
                  "Date": [
                     "Mon, 24 Jun 2019 05:27:57 GMT"
                  ],
                  "Content-Type": [
                     "text/html; charset=utf-8"
                  ]
               },
               "reasonPhrase": "OK",
               "body": {
                  "code": 0
               },
               "statusCode": 200
            }}`
	return []byte(result), nil
}
