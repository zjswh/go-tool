package utils

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"strings"
)

func Request(url string, data map[string]interface{}, header map[string]interface{}, method string, stype string) (body []byte, err error) {
	url = strings.ReplaceAll(strings.ReplaceAll(strings.ReplaceAll(url, "\n", ""), " ", ""), "\r", "")
	param := []byte("")
	if stype == "json" {
		param, _ = json.Marshal(data)
		header["Content-Type"] = "application/json"
	} else {
		s := ""
		for k, v := range data {
			s += fmt.Sprintf("%s=%v&", k, v)
		}
		header["Content-Type"] = "application/x-www-form-urlencoded"
		param = []byte(s)
	}

	client := &http.Client{}
	req, err := http.NewRequest(method, url, bytes.NewReader(param))
	if err != nil {
		err = fmt.Errorf("new request fail: %s", err.Error())
		return
	}

	for k, v := range header {
		req.Header.Add(k, fmt.Sprintf("%s", v))
	}

	res, err := client.Do(req)
	if err != nil {
		err = fmt.Errorf("do request fail: %s", err.Error())
		return
	}

	defer res.Body.Close()
	body, err = ioutil.ReadAll(res.Body)
	if err != nil {
		err = fmt.Errorf("read res body fail: %s", err.Error())
		return
	}
	return
}

func RequestGet(url string) (body []byte, err error) {
	return Request(url, map[string]interface{}{}, map[string]interface{}{}, "GET", "")
}
