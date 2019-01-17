### BFC
Block Formatting Context
![](/assets/bfc-1.jpg)
```html
<style type="text/css">
#box div:first-child{
     width: 100px;
     height: 100px;
     background-color: deepskyblue;
     float: left;
}
#box div:last-child{
     width: 500px;
}
#box{

}
</style>
	
<div id="box">
	<div>

	</div>
	<div>
		Lorem ipsum dolor sit amet, consectetur adipisicing elit. 
		Cumque doloribus enim explicabo fugiat harum
            Cumque doloribus enim explicabo fugiat harum
            Cumque doloribus enim explicabo fugiat harum
		Lorem ipsum dolor sit amet, consectetur adipisicing elit. 
		Cumque doloribus enim explicabo fugiat harum
            Cumque doloribus enim explicabo fugiat harum
            Cumque doloribus enim explicabo fugiat harum
	</div>
</div>
```
在文字多的情况下 , 左侧的100*100盒子会被右侧的文字包裹 , 这个时候可以使用 `overflow:auto` 给右侧的文字盒子 使其左右分离

### 给右侧的文字盒子 加上 `overflow:auto` 后
![](/assets/bfc-2.jpg)
