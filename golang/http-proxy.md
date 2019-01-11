### Http/s代理
一般我们在编写爬虫软件的时候都需要不停的切换UA和IP, 因为一般服务器都会对过渡的请求进行屏蔽, 这就导致我们不能通过同一个IP去请求对方服务器, 而代理能让我们很好的去不停的切换IP, 让对方服务器认为我不是"我", 下面是一段Go的代理请求, 支持http和https.
```go
package main

import (
	"fmt"
	"net/url"
	"net/http"
	"io/ioutil"
	"log"
)

func main() {
	var (
		proxyUrl *url.URL
		rep *http.Response
		req *http.Request
		err error
		data []byte
	)

	if req, err = http.NewRequest("GET", "https://www.baidu.com", nil); err != nil{
		log.Fatal(err)
	}
	
	//填入对应的信息
	if proxyUrl, err = url.Parse("http://username:password@ip:port"); err != nil{
		log.Fatal(err)
	}

	client := &http.Client{
		Transport: &http.Transport{
			Proxy:           http.ProxyURL(proxyUrl),
		},
	}

	if rep, err = client.Do(req); err != nil{
		log.Fatal(err)
	}

	if data, err = ioutil.ReadAll(rep.Body); err != nil{
		log.Fatal(err)
	}

	fmt.Println(string(data))
}
```

请求地址, 已经变成了代理地址

![](/assets/proxy.png)



