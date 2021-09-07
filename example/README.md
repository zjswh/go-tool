# go-tool工具使用

### 安装依赖  

    go get github.com/zjswh/go-tool

### 生成gin项目基础框架代码

    go-tool -a template -name [project name] -dir [dirname]  
    cd [dirname]
    go mod init [project name]
    go mod tidy

### 生成基础gorm的model代码

    go-tool -a model -sql [sql file] -dir [model path]  
