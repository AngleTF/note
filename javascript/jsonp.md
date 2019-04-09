### 什么是JSONP

JSONP\(JSON with Padding\)是JSON的一种“使用模式”，可用于解决主流浏览器的跨域数据访问的问题。由于同源策略，一般来说位于 server1.example.com 的网页无法与不是 server1.example.com的服务器沟通，而 HTML 的 `<script>` 元素是一个例外。利用 `<script>` 元素的这个开放策略，网页可以得到从其他来源动态产生的 JSON 资料，而这种使用模式就是所谓的 JSONP

### 原理是什么?

通过script的可远程获取的策略, 通过后台返回调用JS函数的形式, 并且将传入数据到函数形参, 比如服务器返回了字符串 `callbackFunction('TestData')`, 这本来是没有任何意义的字符串 但是通过script标签获取 则会进行JS方式的解析, 而在JS的眼中 `callbackFunction('TestData')` 则是一个函数调用, 我称它为 **远程执行回调函数**, 而我们可以以GET方式传递给服务器需要回调的函数名, 那么就可以由前端来定义和操作这个回调函数, 而服务器则需要根据需求拼接数据.

### JS实现

```html
<!DOCTYPE html>
<html lang="en">
<head>
    <meta charset="UTF-8">
    <title>Title</title>
</head>
<body>
<div id="divCustomers"></div>
<script type="text/javascript">
    function callbackFunction(result)
    {
        var html = '<ul>';
        for(var i = 0; i < result.length; i++)
        {
            html += '<li>' + result[i] + '</li>';
        }
        html += '</ul>';
        document.getElementById('divCustomers').innerHTML = html;
    }
</script>
<script type="text/javascript" src="http://www.mvc.com/?jsoncallback=callbackFunction"></script>
</body>
</html>
```

### PHP实现

```php
<?php
header('Content-type: application/json');
//获取回调函数名
$jsoncallback = htmlspecialchars($_REQUEST ['jsoncallback']);
//json数据
$json_data = '["customername1","customername2"]';
//输出jsonp格式的数据
echo $jsoncallback . "(" . $json_data . ")";
?>
```



