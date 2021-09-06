package logOperation

import (
	"encoding/json"
	"github.com/gin-gonic/gin"
	"github.com/zjswh/go-tool/utils"
	"time"
)

const LogPushUrl = "http://consoleapi.guangdianyun.tv/v1/log/operation"

func Push(c *gin.Context, XCaStage string, apiMap map[string][]string, data interface{})  {
	path := c.Request.URL.Path
	method := c.Request.Method
	token := c.Request.Header.Get("token")
	ip := c.ClientIP()

	if _, ok := apiMap[path]; !ok {
		return
	}

	reqMap := map[string]interface{}{}
	c.Request.ParseForm()
	for k, v := range c.Request.PostForm {
		reqMap[k] = v[0]
	}
	res, _ := json.Marshal(data)
	req, _ := json.Marshal(reqMap)
	param := map[string]interface{} {
		"token" : token,
		"operate": apiMap[path][0],
		"module": apiMap[path][1],
		"path": path,
		"method": method,
		"req": string(req),
		"res": string(res),
		"type": 2,
		"ip": ip,
		"opTime": time.Now().Unix(),
	}
	utils.Request(LogPushUrl, param, map[string]interface{}{
		"X-CA-STAGE" :XCaStage,
	}, "POST", "form")
}
