package activity

func Common_http(input interface{}, taskRefName string) ([]byte, error) {
	if taskRefName == "get_skeleton" {
		result := ` {"response": {
              "reasonPhrase": "OK",
              "body": {
                 "data": [
                    {
                       "result": [
                          {
                             "extended_geom": "116.36660841886894,39.86871386309301;116.3668336,39.8687313;116.366841,39.8687319;116.366841,39.8687319;116.3676093,39.8687914;116.3694519,39.868907;116.3752482,39.8693138;116.3768951,39.8694301;116.3797232,39.8696368;116.3813213,39.8696911;116.3813863,39.8696955;116.3813865,39.8696955;116.3813865,39.8696955;116.38161934659664,39.86971137404944",
                             "geom": "116.366841,39.8687319;116.366841,39.8687319;116.3676093,39.8687914;116.3694519,39.868907;116.3752482,39.8693138;116.3768951,39.8694301;116.3797232,39.8696368;116.3813213,39.8696911;116.3813863,39.8696955;116.3813865,39.8696955",
                             "way_id": "364178",
                             "way_name": "Guanqumen Nanbinhe Lu",
                             "way_type": "trunk"
                          }
                       ],
                       "type": "normal"
                    }
                 ]
              },
              "statusCode": 200
           }}`
		return []byte(result), nil
	}
	if taskRefName == "save_update_info" {
		result := `{
            "response": {
               "headers": {
                  "Server": [
                     "nginx/1.13.12"
                  ],
                  "Access-Control-Allow-Origin": [
                     "*"
                  ],
                  "Access-Control-Allow-Methods": [
                     "POST, OPTIONS, GET, PUT, DELETE"
                  ],
                  "Access-Control-Allow-Credentials": [
                     "true"
                  ],
                  "Connection": [
                     "keep-alive"
                  ],
                  "Access-Control-Allow-Headers": [
                     "Content-Type, Content-Length, Accept-Encoding, X-CSRF-Token, Authorization, accept, origin, Cache-Control, X-Requested-With"
                  ],
                  "Date": [
                     "Mon, 24 Jun 2019 05:27:58 GMT"
                  ]
               },
               "reasonPhrase": "No Content",
               "body": null,
               "statusCode": 204
            }
         }`
		return []byte(result), nil
	}
	return nil, nil
}
