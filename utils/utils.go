package utils

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"strings"
)

const MessageHost = "http://pubdev.guangdianyun.tv"

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

func SendDms(topic string, cmd string, uin int, data interface{}, withUin bool)  {
	extra, _ := json.Marshal(data)
	url := MessageHost + "/v1/message/Index/send"
	if !withUin {
		url = MessageHost + "/v1/message/Index/sendDms"
	}
	res, err := Request(url, map[string]interface{}{
		"topic": topic,
		"cmd" : cmd,
		"uin" : uin,
		"extra": string(extra),
	}, map[string]interface{}{}, "POST", "form")
	if err != nil {
		fmt.Println(string(res))
	}
}

func StructToMap(data interface{}) map[string]interface{} {
	dataMap := map[string]interface{}{}
	dataBytes, _ := json.Marshal(data)

	//防止底层在输出的时候会进行格式化防止出现科学计数法
	d := json.NewDecoder(bytes.NewReader(dataBytes))
	d.UseNumber()
	_ = d.Decode(&dataMap)
	return dataMap
}

func Reverse(s string) string {
	a := []rune(s)
	for i, j := 0, len(a)-1; i < j; i, j = i+1, j-1 {
		a[i], a[j] = a[j], a[i]
	}
	return string(a)
}
