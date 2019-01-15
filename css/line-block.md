### Line-block 自带高度和宽度
![](/assets/line-block.jpg)
```html
<!DOCTYPE html>
<html lang="en">
<head>
    <meta charset="UTF-8">
    <title>Title</title>
    <style type="text/css">
        *{
            box-sizing: border-box;
            margin:0;
            padding: 0;
        }
        div{
            width: 100%;
            text-align: center;
            border: 1px solid #000;
        }
        div span{
            display: inline-block;
            width: 5px;
            height: 5px;
            border-radius: 50%;
            background-color: red;
        }
    </style>
</head>
<body>
<div>
   <span></span>
   <span></span>
   <span></span>
   <span></span>
</div>
</body>
</html>
```

### 解决宽度
红点和红点之间的间距是由于空白符造成的 , 这个好解决 可以通过`margin-left:-5px` 或者去掉空白的位置即可 ,但是高度的问题还是没有解决
```css
div span{
    display: inline-block;
    width: 5px;
    height: 5px;
    border-radius: 50%;
    background-color: red;
    margin-left: -5px;
}
div span:nth-child(1){
    margin-left: 0;
}
```
![](/assets/Line-block-height.jpg)

### 解决高度
经过长时间的尝试 , 得出结论 可能是转化 Line-block 时会在div内span外产生额外的空白符 , 或者是BFC造成的
![](/assets/Line-block-height-2.jpg)

最终得出两种解决方案
方法一: 父容器的 `font-size` 设置为0
```css
div{
  width: 100%;
  text-align: center;
  border: 1px solid #000;
  font-size:0;
}
```
方法二: 父容器 `line-height` 设置为0 , 并且子元素于中线对齐
```css
div{
  width: 100%;
  text-align: center;
  border: 1px solid #000;
  line-height: 0;
}
div span{
  display: inline-block;
  width: 5px;
  height: 5px;
  border-radius: 50%;
  background-color: red;
  vertical-align: middle;
}
```



