### 效果图
![](/assets/tb.jpg)

###代码
```html
<!DOCTYPE html>
<html lang="en">
<head>
    <meta charset="UTF-8">
    <title>Title</title>
    <style>
        .box{
            margin:100px 0 0 100px;
        }
        .left-box{
            width: 400px;
            height: 400px;
            position: relative;
            border: 1px solid #000;
        }
        .left-box img{
            width: 100%;
            height: 100%;
        }
        .mask{
            position: absolute;
            top:0;
            left:0;
            width: 200px;
            height: 200px;
            background: rgba(255,10,10,.5);
            cursor: move;
            display: none;
        }
        .right-box{
            position: absolute;
            left: 550px;
            top: 100px;
            width: 400px;
            height: 400px;
            overflow: hidden;
            border: 1px solid #000;
        }
        .right-box img{
            position: absolute;
            left:0;
            top:0;
        }
        html{
            height:700px;
        }
    </style>
</head>
<body>
<div class="box">
    <div class="left-box">
        <div class="mask"></div>
        <img src="https://img.alicdn.com/bao/uploaded/i3/1749001929/TB1He7Gd_lYBeNjSszcXXbwhFXa_!!0-item_pic.jpg_430x430q90.jpg" alt="">
    </div>
    <div class="right-box">
        <img src="https://img.alicdn.com/bao/uploaded/i3/1749001929/TB1He7Gd_lYBeNjSszcXXbwhFXa_!!0-item_pic.jpg" alt="">
    </div>
    <div class="test"></div>
</div>
<script>
    var box = document.querySelector('.box'),
        leftBox = document.querySelector('.left-box'),
        rightBox = document.querySelector('.right-box'),
        mask = document.querySelector('.mask');
    var x,y;

    leftBox.onmouseleave = function(){this.children[0].style.display = 'none'};
    leftBox.onmouseenter = function(){this.children[0].style.display = 'block'};
    leftBox.onmousemove = function (e) {
        //计算元素偏移量
        [x,y] = [e.pageX  || e.clientX + document.body.scrollLeft,e.pageY  || e.clientY + document.body.scrollTop];
        x -= box.offsetLeft + mask.offsetWidth / 2;
        y -= box.offsetTop + mask.offsetHeight / 2;
        [x,y] = [0 >= x ? 0 : x,0 >= y ? 0 : y];
        x = this.clientWidth - mask.offsetWidth <= x ? this.clientWidth - mask.offsetWidth : x;
        y = this.clientWidth - mask.offsetHeight <= y ? this.clientWidth - mask.offsetHeight : y;
        [mask.style.left,mask.style.top] = [x + 'px',y + 'px'];
        //计算移动比例倍数
        [wbl,hbl] = [rightBox.children[0].offsetWidth / this.clientWidth, rightBox.children[0].offsetHeight / this.clientHeight];
        [rightBox.children[0].style.top,rightBox.children[0].style.left] = [-y * wbl + 'px',-x * hbl + 'px'];
    };



</script>
</body>
</html>
```