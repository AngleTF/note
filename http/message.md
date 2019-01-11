### 请求报文的组成
```
POST /test/test.php HTTP/1.1	
Host:localhost								
Content-type:application/x-www-form-urlencoded
Content-length:20

act=query&name=ghost
```
| 请求组成部分| 简介|
| ------------- | ------------- |
| 请求行| 规定统一的格式如 http, ftp  |
| Resource  | 可标识的任何东西如 图片,电影  |
| Identifier| 可标识的对象 ,也称标识符  |


主要由请求行[请求方式, 请求URL, HTTP版本组成
