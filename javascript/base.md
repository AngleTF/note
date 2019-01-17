### 原型链
![](/assets/prototype.png)
上面这张图是整个原型链的大概结构

首先我声明了一个Demo构造函数 , 并且将它进行了实例化
```js
function Demo() {}
var demo = new Demo;
```
根据上面的图我们可以看到Demo的实例化对象的 __proto__ 和 Demo构造函数的 prototype 指向同一个地方 那就是Object构造函数的实例化对象, 而Object构造函数的实例化对象的constructor指向Demo构造函数

```js
//true
console.log(Demo.prototype === demo.__proto__);
//true
console.log(Demo.prototype.constructor === Demo);
```

不难发现Demo的实例化对象的原型是Demo构造函数的原型,而他是个Object的对象, 那么Object的对象是否也有原型呢?根据图片我们可以发现 , Object对象的原型也是Object构造函数的原型
```js
//true
console.log(Object.prototype === demo.__proto__.__proto__);
//true
console.log(Object.prototype === {}.__proto__);
```

我的疑问Object构造的原型是否还有原型,Object构造的原型是谁的实例化对象?

的确 Object构造的原型确实还存在原型而且指向的是null,Object构造的原型是不是某个实例化对象我就不得而知了

```js
//true
console.log(Object.prototype.__proto__ === null);
//true
console.log(Object.prototype.__proto__ === Object.prototype === demo.__proto__.__proto__.__proto__);
```

很神奇的是null是一个对象, 但是却没有任何的方法和属性
```js
//object
console.log(typeof null);
```
原理是这样的，不同的对象在底层都表示为二进制，在 JavaScript 中二进制前三位都为 0 的话会被判断为 object 类型，null 的二进制表示是全 0，自然前三位也是 0，所以执行 typeof 时会返回“object”

Demo构造函数和Object构造函数是Function的实例化对象
```js
//function
console.log(typeof Demo);
//function
console.log(typeof Object);
//true
console.log(Object instanceof Function);
//true
console.log(Demo instanceof Function);
```

既然是实例化那么又进入了上面的循环, 对象的原型等于构造的原型....
```js
//true
console.log(Demo.__proto__ === Function.prototype);
//true
console.log(Object.__proto__ === Function.prototype);
```

函数构造的原型的原型等于对象的原型!!
```js
console.log(Function.prototype.__proto__ === Object.prototype);
```

### 函数式作用域和Catch作用域
说说作用域吧, 在ES6之前没有块级作用域的概念,只有函数作用域和catch作用域
```js
{
    var a = 'a';
}
//能输出 , 说明不存在块级作用域
//a
console.log(a);


function foo() {
    var b = 'b'
}
//Uncaught ReferenceError: b is not defined
console.log(b);


try {

}catch (e) {
    var c = 'c';
}
//undefined
console.log(c);
```
function和catch稍有不同,function作用域是直接报错,而cache是会隐式声明,注意是声明而不是赋值,google兼容ES5以下的块级作用域就是通过try抛出数据 , catch接收数据并且声明赋值进行向下兼容

**作用域的原理**
当函数进行调用时,CPU会在开辟一块内存空间当做栈,使用段寄存器SS存放栈顶的段地址,用寄存器SP存放偏移地址, 操作栈无非是两种情况,一种是push(入栈),一种是pop(出栈),每种操作都会对SP进行修改,来改变SS:SP的栈顶地址,函数中的变量声明,就是一种入栈操作,当一个函数运行完后这些数据也就没用了,就会出栈(一般语言都会自动出栈),出栈后的数据就不存在了,所有函数执行完后,函数内的数据就消失了,这就是为什么会有作用域这种东西,那有没有什么办法保存这个函数作用域呢?答案是有的,这个就是后面要说的闭包

### JS的优化之预解析导致的变量提升原则
说说js的预解析吧, js在执行时会将代码交给编译器进行编译, 在编译之前编译器会做一些优化, 它会将变量提升至自身作用域的最顶端

比如以下代码
```js
console.log(a);
var variable = 'a';
```
//这个时候你觉得他会打出什么东西?
//是ReferenceError or undefined ?
//调用在前, 声明在后
//按照正常思维必然是内部抛出ReferenceError的异常
//但是!! js会将代码进行优化成下面的样子 , 这就是预解析

```js
var variable;
//undefined
console.log(a);
variable = 'a';
```

在看个例子
```js
var variable = 'a';
+ function () {
    console.log(variable);
    var variable = 'b';
}()
```

**js优化后**
```js
var variable = 'a';
+ function () {
    var variable
    //undefined
    console.log(variable);
    variable = 'b';
}()
```

### 静态成员和对象成员
```js
function Demo() {
    //添加对象成员
    this.title = 'tao'
}
//添加静态成员可以理解为后端语言的 public static TITLE = 'feng'
Demo.title = 'feng';

var demo = new Demo;

//tao
console.log(demo.title);
//feng
console.log(Demo.title);
```

### 闭包的原理, 闭包有哪些好处

什么是闭包? 很多人都说是函数嵌套函数, 其实这是一个模糊的概念, 因为你根本不知道它为什么要弄成闭包

其实我个人认为闭包是一个作用域的映射,对函数内部的作用域映射, 之前我讲过作用域的原理中函数执行完,数据会丢失,数据会被pop,如何保存这个数据, 闭包就是一个很好的机制. 它会对js底层说 你TM别把这个给我pop了,不然我跟你拼命, ^_^好家伙.

至于闭包的好处我就不多说了吧?

你执行N个函数需要创建N*函数内部的变量

而你可以用闭包映射(通常是return一个函数)的作用域直接对函数内部结构进行修改,而不需要创建.你说哪个更加省时省力?

### JS严格模式与非严格模式有哪些区别
严格模式: `'use strict'`

严格模式的作用主要体现在禁止隐式声明和禁止this指向window

**禁止隐式声明**
```js
variable = 'hhh';
//hhh
console.log(variable);


'use strict';
variable = 'hhh';
//大兄弟报错啦
console.log(variable);
代码variable = 'hhh' 是一种隐式声明
```

**禁止this指向window**
```js
var variable = 'aaa';
+ function () {
    //aaa
    console.log(this.variable);
}()


'use strict';
var variable = 'aaa';
+ function () {
    //大兄弟报错啦
    console.log(this.variable);
}()
```

不允许删除变量或对象, 不允许删除函数, 不允许变量重名, 不允许使用8进制, 不允许使用转移符等...
### 什么是词法作用域? 和动态作用域有什么区别?
首先js是词法作用域, 那么什么是词法作用域呢?直接上代码

词法作用域:书写代码或者说定义时确定的作用域
动态作用域:运行时确定的作用域

```js
var test = 1;

function t1() {
	console.log(test);
}
function t2() {
	var test = 2;
	t1();
}

t2();//1
//这就是词法作用域 , 而动态作用域则会输出2
```

### 上下文函数调用call和apply的区别
call和apply两个函数在框架封装上用的都是特别多,他们两个的作用都是改变this指向,主要区别在于第二个参数
```
call(this[,param1,param2....])

apply(this[,[param1,param2...]])
```
具体使用还得看需求, 比如获取最大值
```js
var arr = [1,2,3,4,5];
//这个时候你就不能用call了
console.log(Math.max.apply(Math, arr));
```