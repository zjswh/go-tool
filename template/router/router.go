package router

import (
	v1 "TEMPLATE/api/v1"
	"TEMPLATE/service"
	"github.com/gin-gonic/gin"
)

func InitRouter(Router *gin.RouterGroup) {
	userRouter := Router.Group("User").
		Use(service.CheckLogin())
	{
		userRouter.POST("create", v1.CreateCoupon)
		userRouter.POST("update", v1.UpdateCoupon)
		userRouter.POST("delete", v1.DeleteCoupon)
		userRouter.GET("info", v1.CouponInfo)
		userRouter.GET("getList", v1.GetCouponList)
	}
}
