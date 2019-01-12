### jQuerAjax
| Key| Value|
| ------------- | ------------- |
| type| 请求类型 |
| url| 请求地址 |
| data | 请求的数据, obj格式 |



```js
$.ajax({
	type:'get',
	url:'www.taolifeng',
	data:{"in_val":'input_val'},
	dateType:json,
	success:function(data){
		
	}
})
```
如果是jsonp 那么需要指定dateType:jsonp
还需要指定jsonpcallback:回调函数名