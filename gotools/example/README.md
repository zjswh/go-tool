# go-tool工具使用

### 安装依赖  

    go get github.com/zjswh/go-tool/gotools

### 生成gin项目基础框架代码

    gotools template -api [api file] -name [project name] -dir [dirname]  
    cd [dirname]
    go mod tidy

### 生成基础gorm的model代码

    gotools model -sql [sql file] -dir [model path]  
