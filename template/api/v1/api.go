package v1

import (
	"TEMPLATE/types/request"
	"TEMPLATE/types/response"
	"github.com/gin-gonic/gin"
	"github.com/zjswh/go-tool/utils"
)

func CreateCoupon(c *gin.Context) {
	var createStruct request.CreateCouponStruct
	err := c.ShouldBind(&createStruct)
	if err != nil {
		response.ParamError("参数缺失", c)
		return
	}
	response.Success(1, c)
	return
}

func UpdateCoupon(c *gin.Context) {
	id := utils.DefaultIntFormValue("couponId", 0, c)
	if id == 0 {
		response.ParamError("参数缺失", c)
		return
	}

	response.Success("修改成功", c)
	return
}

func DeleteCoupon(c *gin.Context) {
	id := utils.DefaultIntFormValue("couponId", 0, c)
	if id == 0 {
		response.ParamError("参数缺失", c)
		return
	}

	response.Success("删除成功", c)
	return
}

func CouponInfo(c *gin.Context) {
	id := utils.DefaultIntParam("id", 0, c)
	if id == 0 {
		response.ParamError("参数缺失", c)
		return
	}

	response.Success("", c)
	return
}

func GetCouponList(c *gin.Context) {
	//page := utils.DefaultIntParam("page", 1, c)
	//num := utils.DefaultIntParam("num", 10, c)
	//name := c.DefaultQuery("name", "")
	//status := utils.DefaultIntParam("status", -1, c)
	//startTime := utils.DefaultIntParam("startTime", 0, c)
	//endTime := utils.DefaultIntParam("endTime", 0, c)

	response.Success(gin.H{
		"list":  "",
		"count": 0,
	}, c)
	return
}
