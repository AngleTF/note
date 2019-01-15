### 什么是Less
Less （Leaner Style Sheets 的缩写） 是一门向后兼容的 CSS 扩展语言。这里呈现的是 Less 的官方文档（中文版），包含了 Less 语言以及利用 JavaScript 开发的用于将 Less 样式转换成 CSS 样式的 Less.js 工具。

### 如何将Less编译成css?
打开 Kaola , 选择一个项目 , 如果less文件在less文件夹内 , Kaola则会自动编译生成 css文件夹和文件
![](/assets/kaola.png)

### Less的基本使用
定义一个变量为Test 值为rgb(0,0,0)

```less
@Test:rgb(0,0,0)

/*使用变量*/
root{
	color:@Test;	
}
```

嵌套使用
```less
.box{
  width:@100 ;
  height:@100 ;
  border:1px solid red ;
    .son{
    background-color: #fff;
    /*伪元素*/﻿
    &::before{
      content: "";
      clear: both;
      display: none;
    }
  }
}
```

使用, 定义类
```less
.Test{
	background:red;
	color:red;
}
	
body{
	.Test;
}
```

定义ID函数 并且可以传入参数 , 默认值为red
```less
﻿#Tag(@color:red){
  background:@color;
  color:@color;
  border: 1px solid @color;
}

/*使用ID函数*/
body{
  #Tag(yellow);
}
```

### Less的好处
1. 编写可维护代码
2. 结构清晰
3. 可自定义变量和函数
4. 减少代码量