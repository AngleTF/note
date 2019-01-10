CODE
```go
package main

import (
	"net/http"
	"net/url"
	"io/ioutil"
	"fmt"
	"log"
)


func main() {
	var (
		u *url.URL
		err error
		tr *http.Transport
		client http.Client
		req *http.Request
		rps *http.Response
		ioData []byte
	)

	//解析代理地址
	if u, err = url.ParseRequestURI("http://58.53.128.83:3128"); err != nil{
		log.Fatal(err)
	}

	tr = &http.Transport{
		//设置http代理地址
		Proxy: http.ProxyURL(u),
	}

	client = http.Client{
		Transport:tr,
	}

	//请求地址
	if req, err = http.NewRequest("GET", "http://47.98.50.193:8888/test.php",nil); err != nil{
		log.Fatal(err)
	}

	if rps, err = client.Do(req); err != nil{
		log.Fatal(err)
	}

	if ioData, err = ioutil.ReadAll(rps.Body); err != nil{
		log.Fatal(err)
	}

	fmt.Println(string(ioData))
}
```

> ##请求地址, 已经变成了代理地址

![](/assets/proxy.png)



