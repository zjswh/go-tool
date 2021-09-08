package gen

import (
	"fmt"
	"github.com/tal-tech/go-zero/tools/goctl/api/gogen"
	"github.com/tal-tech/go-zero/tools/goctl/api/spec"
	"github.com/tal-tech/go-zero/tools/goctl/util"
	"io/ioutil"
	"os"
	"path"
	"strings"
)

func GenApi(api *spec.ApiSpec, projectName, destPath string) error {
	apiDir := destPath + "/api/" + dealVersion(api.Syntax.Version)
	err := util.MkdirIfNotExist(apiDir)
	if err != nil {
		return err
	}
	apiFunc := ""
	for _, r := range api.Service.Groups {
		for _, v := range r.Routes {
			funcInfo := functionTemplate
			funcName := path.Base(v.Path)
			funcInfo = strings.ReplaceAll(funcInfo, "FUNC_NAME", v.Handler)
			funcInfo = strings.ReplaceAll(funcInfo, "VAR_STRUCT", funcName+"Request")
			funcInfo = strings.ReplaceAll(funcInfo, "STRUCT_E", v.RequestType.Name())
			apiFunc += funcInfo
		}
	}
	apiTemp = strings.ReplaceAll(apiTemp, "FUNC_LIST", apiFunc)
	apiTemp = strings.ReplaceAll(apiTemp, "TEMPLATE", projectName)
	err = ioutil.WriteFile(apiDir+"/"+api.Service.Name+".go", []byte(apiTemp), os.ModePerm)
	return err
}

func GenService(api *spec.ApiSpec, projectName, destPath string) error {
	serviceName := api.Service.Name + "Service"
	servicePath := destPath + "/service/" + serviceName
	err := util.MkdirIfNotExist(servicePath)
	if err != nil {
		return err
	}
	serviceFunc := ""
	serviceTemp = strings.ReplaceAll(serviceTemp, "SERVICE_NAME", serviceName)
	for _, r := range api.Service.Groups {
		for _, v := range r.Routes {
			funcInfo := serviceFunctionTemplate
			funcInfo = strings.ReplaceAll(funcInfo, "FUNC_NAME", v.Handler)
			funcInfo = strings.ReplaceAll(funcInfo, "STRUCT_E", v.RequestType.Name())
			serviceFunc += funcInfo
		}
	}
	serviceTemp = strings.ReplaceAll(serviceTemp, "FUNC_LIST", serviceFunc)
	serviceTemp = strings.ReplaceAll(serviceTemp, "TEMPLATE", projectName)
	err = ioutil.WriteFile(servicePath+"/"+serviceName+".go", []byte(serviceTemp), os.ModePerm)
	return nil
}

func GenTypes(api *spec.ApiSpec, destPath string) error {
	//创建文件夹
	requestDir := destPath + "/types"
	err := util.MkdirIfNotExist(requestDir)
	if err != nil {
		return err
	}
	typesContent, err := gogen.BuildTypes(api.Types)
	if err != nil {
		return err
	}
	err = ioutil.WriteFile( requestDir+"/types.go", []byte("package types\n\n"+typesContent), os.ModePerm)
	return nil
}

func GenRoutes(api *spec.ApiSpec, projectName, destPath string) error {
	//创建文件夹
	routerDir := destPath + "/router"
	err := util.MkdirIfNotExist(routerDir)
	useMiddleImport := false
	if err != nil {
		return err
	}
	routerContent := ""
	for _, r := range api.Service.Groups {
		middleware := r.Annotation.Properties["middleware"]
		serviceName := api.Service.Name
		routerName := serviceName + middleware
		router := fmt.Sprintf("\t%sRouter := Router.Group(\"\")", routerName)
		if middleware != "" {
			router = fmt.Sprintf("%s.\n\t\tUse(middleware.%s())", router, middleware)
			useMiddleImport = true
		}

		router += "\n\t{\n"
		for _, v := range r.Routes {
			router += fmt.Sprintf("\t\t%sRouter.%s(\"%s\", v1.%s)\n", routerName, strings.ToUpper(v.Method), v.Path[1:], v.Handler)
		}
		router += "\t}\n\n"
		routerContent += router
	}

	//判断是否使用了中间件
	if useMiddleImport == true {
		routerTemplate = strings.ReplaceAll(routerTemplate, "MIDDLEWARE_IMPORT", "\""+projectName+"/middleware\"")
	}
	routerTemplate = strings.ReplaceAll(routerTemplate, "ROUTER_TEMP", routerContent)
	routerTemplate = strings.ReplaceAll(routerTemplate, "TEMPLATE", projectName)

	err = ioutil.WriteFile(routerDir+"/router.go", []byte(routerTemplate), os.ModePerm)
	return err
}

func dealVersion(version string) string {
	return strings.ReplaceAll(version, "\"", "")
}
