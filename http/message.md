### 请求报文的组成
```http
POST /test/test.php HTTP/1.1	
Host:localhost								
Content-type:application/x-www-form-urlencoded
Content-length:20

act=query&name=ghost
```

| 请求组成部分| 简介|
| ------------- | ------------- |
| 请求行| 请求方式, 请求URL, HTTP版本组成 |
| 请求头| 字段名和字段值组成  key:value 形式|
| 请求体 | 请求的数据, 例如上面的 act=query&name=ghost |


### 响应报文的组成
```http
HTTP/1.1 200 OK	
Date:Sun 28 Aug 2016 03.08.56
Server:Apache/2.4.10
X-Powered-By:PHP/5.3.29
Content-Length:9				
Content-Type:text/html
<html>
...........
</html>
```

| 请求组成部分| 简介|
| ------------- | ------------- |
| 响应行| HTTP版本, 状态码和描述 |
| 请求头| 字段名和字段值组成  key:value 形式|
| 请求体 | 响应的数据, 例如上面的html代码 |


