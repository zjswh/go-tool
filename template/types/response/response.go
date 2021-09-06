package response

import (
	"TEMPLATE/config"
	"TEMPLATE/types/appConst"
	"github.com/gin-gonic/gin"
	"github.com/zjswh/go-tool/logOperation"
	"net/http"
)

type Response struct {
	Code         int         `json:"code"`
	Data         interface{} `json:"data"`
	ErrorMessage string      `json:"errorMessage"`
	ErrorCode    int         `json:"errorCode"`
}

type GdyAPiResult struct {
	Code         int                   `json:"code"`
	ErrorCode    int                   `json:"errorCode"`
	ErrorMessage string                `json:"errorMessage"`
	Data         map[string][]LiveList `json:"data"`
}

type LiveList struct {
	Id         int    `json:"id"`
	Name       string `json:"name"`
	Logo       string `json:"logo"`
	CreateTime string `json:"createTime"`
	Type       string `json:"type"`
	WatchNum   int    `json:"watchNum"`
	IsSelect   int    `json:"isSelect"`
}

func Success(data interface{}, c *gin.Context) {
	Result(0, data, "", c)
}

func Result(code int, data interface{}, msg string, c *gin.Context) {
	result := Response{
		200,
		data,
		msg,
		code,
	}
	go logOperation.Push(c, config.GVA_CONFIG.Param.XCaStage, appConst.ApiMap, result)
	c.JSON(http.StatusOK, result)
}

func SystemError(msg string, c *gin.Context) {
	// 开始时间
	c.JSON(http.StatusInternalServerError, Response{
		http.StatusInternalServerError,
		"",
		msg,
		0,
	})
	c.Abort()
}

func ParamError(message string, c *gin.Context) {
	Result(2, "", message, c)
}

func DbError(message string, c *gin.Context) {
	Result(3, "", message, c)
}

