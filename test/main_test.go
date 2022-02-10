package test

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net"
	"net/http"
	"time"
)

type header struct {
	Key   string
	Value string
}

var (
	host = "http://localhost:8080"
)

func PerformRequest(method, path string, req, res interface{}, headers ...header) (*http.Response, error) {
	body, err := json.Marshal(req)
	if err != nil {
		return nil, err
	}

	tr := &http.Transport{
		DialContext: (&net.Dialer{
			Timeout:   10 * time.Second,
			KeepAlive: 10 * time.Second,
		}).DialContext,
		TLSHandshakeTimeout: 10 * time.Second,

		ExpectContinueTimeout: 10 * time.Second,
		ResponseHeaderTimeout: 10 * time.Second,
		MaxIdleConns:          10,
		MaxConnsPerHost:       10,
	}

	client := &http.Client{
		Transport: tr,
		Timeout:   10 * time.Second,
	}

	request, err := http.NewRequest(method, fmt.Sprintf("%s%s", host, path), bytes.NewBuffer(body))
	if err != nil {
		return nil, err
	}

	for _, h := range headers {
		request.Header.Add(h.Key, h.Value)
	}

	request.Header.Add("Accept", "application/json")
	resp, err := client.Do(request)
	if err != nil {
		return resp, err
	}

	defer resp.Body.Close()
	resp_body, _ := ioutil.ReadAll(resp.Body)

	err = json.Unmarshal(resp_body, res)
	if err != nil {
		return resp, err
	}

	return resp, nil
}
