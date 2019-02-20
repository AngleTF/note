### go get golang.org/x下载失败
国内无法通过 go get 去请求国外的网站进行下载, 但是在Github上有相应的代码

### download
```
mkdir -p $GOPATH/src/golang.org/x
cd $GOPATH/src/golang.org/x
git clone https://github.com/golang/net.git
```

### package url
```
https://github.com/golang/[PACKAGE_NAME].git
```