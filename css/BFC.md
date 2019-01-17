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

我们把文字减少, 会发现另外一个问题, 左侧蓝色盒子不在 #box 高度内
![](/assets/bfc-3.png)
这是由于浮动原因导致父元素坍塌
### 给\#box 加上 `overflow:auto` 后
![](/assets/bfc-4.jpg)
溢出的问题也消失了

### BFC作用
一旦BFC创建，它就会阻止它内部的元素逃离并从盒子里伸出来 , 没BFC的margin-top会逃出盒子 , 有BFC的margin-top不会逃出盒子
![](/assets/bfc-5.jpg)
```html
<style type="text/css">
.outer {
      background-color: #ccc;
}

.outer div {
      padding: 0;
      margin-top: 50px;
      background-color:deepskyblue;
      color: #fff;
      height: 100px;
}

.overflow {
      overflow: auto;
}
</style>
	
<h2>no BFC</h2>
<div class="outer">
	<div></div>
</div>

<h2>Have BFC</h2>

<div class="outer overflow">
	<div></div>
</div>
```

### 总结
将 overflow 属性 设置为非默认值visible的值 , 其实就是添加一个BFC区块