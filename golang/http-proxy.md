### Http/s代理
一般我们在编写爬虫软件的时候都需要不停的切换UA和IP, 因为一般服务器都会对过渡的请求进行屏蔽, 这就导致我们不能通过同一个IP去请求对方服务器, 而代理能让我们很好的去不停的切换IP, 让对方服务器认为我不是"我", 下面是一段Go的代理请求, 支持http和https.
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



