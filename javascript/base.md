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