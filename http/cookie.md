### Cookie
在请求和响应报文中写入叫做Cookie信息来控制客户端的状态, cookie会根据从服务端发送的响应报文的一个叫做
Set-Cookie的首部字段信息,通知客户端保存信息

第一次请求报文(无cookie)
```http
GET /reader/ HTTP/1.1
Host:hackr.jp
```

第一次响应报文, 服务端要求客户端存储这个cookie(设置cookie)
```http
HTTP/1.1 200 OK
Date:Sun 28 Aug 2016 03.08.56
Server:Apache/2.4.10
X-Powered-By:PHP/5.3.29
Content-Length:9				
Content-Type:text/html; charset=UTF-8
Set-Cookie:PHPSESSID=86be21bb32bb7ede2dfa8860ef5735f9; path=/
```

第二次请求报文(有cookie)
```http
GET /reader/ HTTP/1.1
Host:hackr.jp
Cookie:PHPSESSID=86be21bb32bb7ede2dfa8860ef5735f9
```

