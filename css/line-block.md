### 问题代码
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
            position: absolute;
            top: 0;
            left: 0;
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