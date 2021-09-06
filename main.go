package main

import (
	"bytes"
	"encoding/json"
	"errors"
	"flag"
	"fmt"
	"io/ioutil"
	"net/http"
	"os"
	"path"
	"path/filepath"
	"regexp"
	"runtime"
	"strings"
)

var sqlAddr = flag.String("sql", "", "sql file path")
var action = flag.String("a", "", "action")
var name = flag.String("name", "", "project name")
var targetAddr = flag.String("dir", "", "code produced path")

func main() {
	flag.Parse()
	switch *action {
	case "model":
		model()
	case "template":
		template()
	default:
		fmt.Println("action miss")
	}
}

func getCurrentPath() string {
	var abPath string
	_, filename, _, ok := runtime.Caller(0)
	if ok {
		abPath = path.Dir(filename)
	}
	return abPath + "/"
}

func template() {
	CopyDir(getCurrentPath() +  "template", *targetAddr, *name)
	fmt.Println("项目创建成功")
}

func model() {
	if *sqlAddr == "" || !checkFileIsExist(*sqlAddr) {
		fmt.Println("sql路径不存在")
		return
	}

	if *targetAddr == "" {
		fmt.Println("请输入文件生成地址")
		return
	}

	//判断文件生成地址是否存在
	if b, _ := pathExists(*targetAddr); !b {
		err := os.Mkdir(*targetAddr, os.ModePerm)
		if err != nil {
			fmt.Println(err)
			return
		}
	}

	bt, _ := os.ReadFile(*sqlAddr)
	arr := strings.Split(string(bt), ";")

	funcTemplateByte, err  := ioutil.ReadFile(getCurrentPath() + "model.tpl")
	funcTemplateStr := string(funcTemplateByte)
	if err != nil {
		fmt.Println(err)
		return
	}
	for _, v := range arr {
		if v != "" && v != "\n" && v != "\r\n" {
			sqlName := regexpData(v, "CREATE TABLE `(.*?)` ")
			dir := *targetAddr
			dir = strings.TrimRight(dir, "/")
			filename := dir + "/" + Case2Camel(sqlName) + ".go"
			if !checkFileIsExist(filename) { //如果文件存在
				var result struct {
					Data  string `json:"data"`
					Error string `json:"error"`
				}
				res, _ := Request("https://www.printlove.cn/api/sql2gorm", map[string]interface{}{
					"ddl": v + ";",
				}, map[string]interface{}{}, "POST", "form")
				json.Unmarshal(res, &result)
				structName := regexpData(result.Data, "type (.*?) struct")
				funcTemplateStrCopy := funcTemplateStr
				funcTemplateStrCopy = strings.Replace(funcTemplateStrCopy, "Template", Case2Camel(structName), -1)
				arr1 := strings.Split(result.Data, "\n")
				importIndex := 2
				hasImport := false
				hasTableNameFunc := false
				for k, c := range arr1 {
					if strings.Contains(c, "import (") {
						importIndex = k + 1
						hasImport = true
					}
					if strings.Contains(c, "TableName()") {
						hasTableNameFunc = true
					}
				}

				var insertArr []string
				if hasImport == true {
					insertArr = []string{"\"gorm.io/gorm\"", ""}
				} else {
					insertArr = []string{"import \"gorm.io/gorm\"", ""}
				}
				arr1 = append(arr1[:importIndex], append(insertArr, arr1[importIndex:]...)...)
				if hasTableNameFunc == false {
					insertArr = []string{fmt.Sprintf("func (m *%s) TableName() string {", Case2Camel(sqlName)), fmt.Sprintf("\treturn \"%s\"", sqlName), "}", ""}
					arr1 = append(arr1, insertArr...)
				}
				result.Data = strings.Join(arr1, "\n")
				err := ioutil.WriteFile(filename, []byte(result.Data+"\n"+funcTemplateStrCopy), 0644)
				if err != nil {
					fmt.Println("model文件生成失败，原因:" + err.Error())
				}
				fmt.Println("文件生成成功")
			} else {
				fmt.Println("文件已存在")
			}
		}
	}
}

func regexpData(str string, pattern string) string {
	reg2 := regexp.MustCompile(pattern)
	result2 := reg2.FindAllStringSubmatch(str, -1)
	return result2[0][1]
}

// 下划线写法转为驼峰写法
func Case2Camel(name string) string {
	name = strings.Replace(name, "_", " ", -1)
	name = strings.Title(name)
	return strings.Replace(name, " ", "", -1)
}

func checkFileIsExist(filename string) bool {
	if _, err := os.Stat(filename); os.IsNotExist(err) {
		return false
	}
	return true
}

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

/**
 * 拷贝文件夹,同时拷贝文件夹中的文件
 * @param srcPath  		需要拷贝的文件夹路径: D:/test
 * @param destPath		拷贝到的位置: D:/backup/
 */
func CopyDir(srcPath, destPath, name string) error {
	//检测目录正确性
	if srcInfo, err := os.Stat(srcPath); err != nil {
		fmt.Println(err.Error())
		return err
	} else {
		if !srcInfo.IsDir() {
			e := errors.New("srcPath不是一个正确的目录！")
			fmt.Println(e.Error())
			return e
		}
	}

	if b, _ := pathExists(destPath); !b {
		err := os.Mkdir(destPath, os.ModePerm)
		if err != nil {
			return err
		}
	}

	if destInfo, err := os.Stat(destPath); err != nil {
		fmt.Println(err.Error())
		return err
	} else {
		if !destInfo.IsDir() {
			e := errors.New("destInfo不是一个正确的目录！")
			fmt.Println(e.Error())
			return e
		}
	}

	//加上拷贝时间:不用可以去掉
	err := filepath.Walk(srcPath, func(path string, f os.FileInfo, err error) error {
		if f == nil {
			return err
		}
		if !f.IsDir() {
			path = strings.Replace(path, "\\", "/", -1)
			destNewPath := strings.Replace(path, srcPath, destPath, -1)
			copyFile(path, destNewPath, name)
		}
		return nil
	})
	if err != nil {
		fmt.Printf(err.Error())
	}
	return err
}

//生成目录并拷贝文件
func copyFile(src, dest, name string) error {
	srcFile, err := os.Open(src)
	if err != nil {
		fmt.Println(err.Error())
		return err
	}
	defer srcFile.Close()
	//分割path目录
	destSplitPathDirs := strings.Split(dest, "/")

	//检测时候存在目录
	destSplitPath := ""
	for index, dir := range destSplitPathDirs {
		if index < len(destSplitPathDirs)-1 {
			destSplitPath = destSplitPath + dir + "/"
			b, _ := pathExists(destSplitPath)
			if b == false {
				fmt.Println("创建目录:" + destSplitPath)
				//创建目录
				err := os.Mkdir(destSplitPath, os.ModePerm)
				if err != nil {
					fmt.Println(err)
				}
			}
		}
	}

	fileContentByte, _ := os.ReadFile(src)
	fileContent := strings.ReplaceAll(string(fileContentByte), "TEMPLATE", name)

	return ioutil.WriteFile(dest, []byte(fileContent), 0644)
}

//检测文件夹路径时候存在
func pathExists(path string) (bool, error) {
	_, err := os.Stat(path)
	if err == nil {
		return true, nil
	}
	if os.IsNotExist(err) {
		return false, nil
	}
	return false, err
}
