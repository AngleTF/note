### meta标签的应用

meta标签的属性content需要配合name或http-equiv使用, 当content配合name使用时, 属于SEO和网页样式的知识, 当content配合http-equiv使用时, 属于HTTP的知识

### 网页关键字
```html
<meta name="keyword" content="苹果,橘子"/>
```	

### 网页描述
```html
<meta name="description" content="这是一个购物网站"/>
```

### 网页重定向 5s后跳转至百度
属性http-equiv代表Http请求报文的首部字段值 refresh: 5;https://baidu.com
```html
<meta http-equiv="refresh" content="5;https://baidu.com" />
```	

### Author
```html
<meta name="Author" content="you name">
```
### 该文件编写的工具
```html
<meta name="generator" content="webStorm">
```
		
### Robots爬虫的收录
```html
<meta name="Robots" content="all|none|index|noindex|follow|nofollow">
```

|content|desc|
|---|---|
|all |文件将被索引,页面链接可以被查询|
|none |文件将不被索引,切页面上的链接不可被查询|
|follow |页面链接可被查询|
|index|文件将被索引|
|noindex|页面将不被查询,但链接可被查询|


### 字符集
```html
<meta charset="ISO-8859-1">
```
	
### 定义最新版本
```html
<meta name="revised" content="David, 2008/8/8/" />
```
	
		
### X-UA-Compatible 
IE8的一个专有<meta>属性，它告诉IE8采用何种IE版本去渲染网页

告诉IE浏览器，无论是否用DTD声明文档标准，IE8/9都会以IE7引擎来渲染页面

```html
<meta http-equiv="X-UA-Compatible" content="IE=7" >   
```

告诉IE浏览器，IE8/9都会以IE8引擎来渲染页面。 
```html 
<meta http-equiv="X-UA-Compatible" content="IE=8" >   
```	


告诉IE浏览器，IE8/9及以后的版本都会以Edge引擎来渲染页面 
```html
<meta http-equiv="X-UA-Compatible" content="IE=edge" >   
```
	
告诉IE浏览器 , 使用chrome内核渲染页面
```html
<meta http-equiv="X-UA-Compatible" content="chrome=1">
```
	
### 告诉360(双核浏览器)使用哪种内核
```html
<meta name="renderer" content="webkit">
			
			webkit			webkit内核
			ie-comp 		IE兼容核
			ie-stand		IE标准核

```		
					
### 移动端视口
```html
<meta name="viewport"
	content="
	width=device-width, 
	user-scalable=no, 
	initial-scale=1.0, 
	maximum-scale=1.0, 
	minimum-scale=1.0">
```
|field|desc|
|---|---|
|viewport|视口属性 , 只适用与手机端|
|width|网页宽度 , 设置layout viewport 的宽度，为一个正整数，使用字符串”width-device”表示设备宽度|
|height|网页宽度 , 设置layout viewport 的高度，为一个正整数|
|user-scalable|是否支持用户缩放 yes或者no|
|initial-scale|设置页面的初始缩放值|
|maximum-scale|允许用户的最大缩放值|
|minimum-scale|允许用户的最小缩放值|


### iPhone手机是否解析手机号码,邮箱号码,地址 , 如果解析会变成拨号连接,跳转链接
```html
<meta name="format-detection" content="telephone=no,email=no,adress=no"/>
```
### iPhone手机是否显示默认的苹果工具栏和菜单栏。
```html
<meta name="apple-mobile-web-app-capable" content="no" >
```
	
### 控制iPhone控制栏颜色
```html
<meta name="apple-mobile-web-app-status-bar-style" content="your color"> 
```
		
### 全屏显示网页, 部分移动端浏览器的属性
```html		
<meta name="full-screen" content="yes"/>
```	

### 强制坚屏显示, 部分移动端浏览器的属性
```html
<meta name="browsermode" content="application"/>
```	