package middleware

import (
	"TEMPLATE/config"
	"TEMPLATE/types/response"
	"encoding/json"
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/zjswh/go-tool/utils"
)

type UserInfo struct {
	AccountId int    `json:"accountId"`
	Aid       int    `json:"aid"`
	Uin       int    `json:"uin"`
	Name      string `json:"name"`
}

type CUserInfo struct {
	Uin         int    `json:"uin"`
	Id          int    `json:"id"`
	IsSafe      int    `json:"isSafe"`
	Phone       string `json:"phone"`
	UserNick    string `json:"userNick"`
	LoginType   string `json:"loginType"`
	UserHeadImg string `json:"userHeadImg"`
	UserIp      string `json:"userIp"`
}

func CheckLogin() gin.HandlerFunc {
	return func(c *gin.Context) {
		token := c.GetHeader("token")
		if token == "" {
			token = c.Query("token")
		}

		if token == "" {
			response.Result(1, "", "未登录", c)
			c.Abort()
			return
		}
		realToken := JwtAuth(token)
		if realToken == "" {
			response.Result(1, "", "登录过期", c)
			c.Abort()
			return
		}
		url := config.GVA_CONFIG.Param.BGatewayHost + "/v1/Passport/Index/getLoginInfo"
		url = fmt.Sprintf("%s?token=%s&path=/%s&method=%s", url, token, c.FullPath(), c.Request.Method)
		res, _ := utils.Request(url, map[string]interface{}{}, map[string]interface{}{
			"X-CA-STAGE": config.GVA_CONFIG.Param.XCaStage,
		},"GET", "")
		var result response.Response
		json.Unmarshal(res, &result)
		if result.Code != 200 || result.ErrorCode != 0 {
			response.Result(result.ErrorCode, "", result.ErrorMessage, c)
			c.Abort()
			return
		}

		userInfo, _ := json.Marshal(result.Data)
		c.Set("userInfo", string(userInfo))
		c.Next()
	}
}

func GetCUserInfo(c *gin.Context) CUserInfo {
	userInfo := CUserInfo{}
	token := c.GetHeader("token")
	if token == "" {
		return userInfo
	}

 	realToken := JwtAuth(token)
	if realToken == "" {
		return userInfo
	}

	info, _ := config.GVA_REDIS.Get(realToken).Result()
	json.Unmarshal([]byte(info), &userInfo)
	return userInfo
}
