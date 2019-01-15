### 原文转自

https://blog.csdn.net/panshiqu/article/details/53788067

Golang 支持交叉编译，在一个平台上生成另一个平台的可执行程序，最近使用了一下，非常好用，这里备忘一下。


### Mac编译Linux64位可执行程序
```
CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build main.go
```

### Mac编译Windows64位可执行程序
```
CGO_ENABLED=0 GOOS=windows GOARCH=amd64 go build main.go
```


### Linux编译Mac64位
```
CGO_ENABLED=0 GOOS=darwin GOARCH=amd64 go build main.go
```

### Linux编译Windows64位
```
CGO_ENABLED=0 GOOS=windows GOARCH=amd64 go build main.go
```

### Windows编译Mac64位
```
CGO_ENABLED=0 GOOS=darwin GOARCH=amd64 go build main.go
```

### Windows编译Linux64位
```
CGO_ENABLED=0 GOOS=windows GOARCH=amd64 go build main.go
```

### Windows编译 Mac64位
```
SET CGO_ENABLED=0
SET GOOS=darwin
SET GOARCH=amd64
go build main.go
```

### Windows编译 Linux64位
```
SET CGO_ENABLED=0
SET GOOS=linux
SET GOARCH=amd64
go build main.go
```

### 参数明细

|参数名|参数简介|
|----|-----|
|GOOS|目标平台的操作系统(darwin、freebsd、linux、windows)|
|GOARCH|目标平台的体系架构(386、amd64、arm) |


交叉编译不支持 CGO 所以要禁用它

上面的命令编译 64 位可执行程序，你当然应该也会使用 386 编译 32 位可执行程序 

很多博客都提到要先增加对其它平台的支持，但是我跳过那一步，上面所列的命令也都能成功，且得到我想要的结果，可见那一步应该是非必须的，或是我所使用的 Go 版本已默认支持所有平台。



很不错的Go commandLine 介绍

http://wiki.jikexueyuan.com/project/go-command-tutorial/0.1.html