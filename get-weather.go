// image: hojjat12/get-weather-prediction-from-open-meteo:lambda
package main

import (
	"context"
	"encoding/json"
	"fmt"
	"io"
	"io/ioutil"
	"net/http"
	"time"
)

type Event struct {
	Body           map[string]interface{} `json:"body"`
	ResultEndpoint string                 `json:"RESULT_ENDPOINT"`
	Authorization  string                 `json:"AUTHORIZATION"`
}

type Response struct {
	Successfull bool                   `json:"successfull"`
	Status      string                 `json:"status"`
	ReturnValue map[string]interface{} `json:"return_value"`
}

func HandleLambdaEvent(event Event) (Response, error) {
	fmt.Println("event.Body:", event.Body)
	resp := Response{}
	resp.Successfull = true
	singleInput := event.Body
	latitude := singleInput["latitude"].(string)
	longitude := singleInput["longitude"].(string)

	url := fmt.Sprintf("https://api.open-meteo.com/v1/forecast?latitude=%s&longitude=%s&timezone=auto&current_weather=true&daily=temperature_2m_max,temperature_2m_min,precipitation_sum,sunrise,sunset", latitude, longitude)
	out, err, statusCode, _ := httpRequest(http.MethodGet, url, nil, nil, 0)
	fmt.Println("response of get weather prediction:", string(out))
	if err != nil || statusCode != http.StatusOK {
		if err != nil {
			fmt.Println(err)
		}
		resp.Successfull = false
	}

	var httpCallresp map[string]interface{}
	json.Unmarshal(out, &httpCallresp)
	resp.ReturnValue = map[string]interface{}{
		"outputs": httpCallresp,
	}
	if resp.Successfull {
		resp.Status = "completed"
		fmt.Println("Getting weather prediction was successfull")
	} else {
		resp.Status = "failed"
		fmt.Println("Getting weather prediction wasn't successfull")
	}
	return resp, nil
}

type Header struct {
	Key   string
	Value string
}

func httpRequest(method string, url string, body io.Reader, headers []Header, timeout time.Duration) (out []byte, err error, statusCode int, header *http.Header) {

	var req *http.Request
	if timeout > 0 {
		ctx, cancel := context.WithTimeout(context.Background(), timeout)
		defer cancel()
		req, err = http.NewRequestWithContext(ctx, method, url, body)
	} else {
		req, err = http.NewRequest(method, url, body)
	}
	if err != nil {
		return
	}

	for _, header := range headers {
		req.Header.Add(header.Key, header.Value)
	}

	c := &http.Client{}
	res, err := c.Do(req)
	if err != nil {
		return
	}
	defer res.Body.Close()

	statusCode = res.StatusCode
	out, err = ioutil.ReadAll(res.Body)
	header = &res.Header
	return
}
