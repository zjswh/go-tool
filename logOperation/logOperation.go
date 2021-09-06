package logOperation

import (
	"encoding/json"
	"github.com/gin-gonic/gin"
	"github.com/zjswh/go-tool/utils"
	"time"
)

var ApiMap = map[string][]string{
	"/v1/coupon/market/Discount/create":         []string{"创建优惠券", "活动管理-优惠券"},
	"/v1/coupon/market/Discount/update":         []string{"修改优惠券", "活动管理-优惠券"},
	"/v1/coupon/market/Discount/delete":         []string{"删除优惠券", "活动管理-优惠券"},
	"/v1/coupon/market/Discount/saveInclude":    []string{"配置优惠券", "活动管理-优惠券"},
	"/v1/coupon/market/DiscountDetail/writeOff": []string{"核销优惠券", "活动管理-优惠券"},
}

const LogPushUrl = "http://consoleapi.guangdianyun.tv/v1/log/operation"

func Push(c *gin.Context,XCaStage string, data interface{})  {
	path := c.Request.URL.Path
	method := c.Request.Method
	token := c.Request.Header.Get("token")
	ip := c.ClientIP()

	if _, ok := ApiMap[path]; !ok {
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
		"operate": ApiMap[path][0],
		"module": ApiMap[path][1],
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
