package gen


var (
	routerTemplate = `
package router

import (
	"github.com/gin-gonic/gin"
	v1 "TEMPLATE/api/v1"
	MIDDLEWARE_IMPORT
)

func InitRouter(Router *gin.RouterGroup) {
ROUTER_TEMP
}
`

	apiTemp = `package v1

import (
	"TEMPLATE/service/couponService"
	"TEMPLATE/types"
	"TEMPLATE/types/response"
	"github.com/gin-gonic/gin"
)
FUNC_LIST
`
	functionTemplate = `
func FUNC_NAME(c *gin.Context) {
	var VAR_STRUCT types.STRUCT_E
	err := c.ShouldBind(&VAR_STRUCT)
	if err != nil {
		response.ParamError("参数缺失", c)
		return
	}

	err = couponService.FUNC_NAME(VAR_STRUCT)
	if err != nil {
		response.DbError(err.Error(), c)
		return
	}

	response.Success("", c)
	return
}
`

	serviceTemp = `package SERVICE_NAME

import (
	"TEMPLATE/types"
)
FUNC_LIST
`

	serviceFunctionTemplate = `
func FUNC_NAME(req types.STRUCT_E) (error) {
	//add your code ...
	return nil
}
`
)
