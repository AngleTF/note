### jQuerAjax

参数一览表

| Key| Value|
| ------------- | ------------- |
| type| 请求类型 |
| url| 请求地址 |
| data | 请求的数据, obj格式 |
| dataType | 返回时被转化的类型 |
| success| 请求成功的回调函数 |
| error | 请求失败的回调函数 |
| timeout| 超时时间, 单位毫秒 |
| async| 是否异步, 默认是true |
| contentType| 请求时的http头部contentType信息 |

```js
$.ajax({
	type:'post',
	url:'www.taolifeng',
	data:{"in_val":'input_val'},
	dataType:json,
	contentType:"application/x-www-form-urlencoded"
	success:function(data){
		
	},
	error:function(xhr,status,error){
	
	}	
})
```
如果是jsonp 那么需要指定dateType:jsonp
还需要指定jsonpcallback:回调函数名