### JS非多线程语言
JavaScript语言的一大特点就是单线程，也就是说，同一个时间只能做一件事。那么，为什么JavaScript不能有多个线程呢？这样能提高效率啊。

JavaScript的单线程，与它的用途有关。作为浏览器脚本语言，JavaScript的主要用途是与用户互动，以及操作DOM。这决定了它只能是单线程，否则会带来很复杂的同步问题。比如，假定JavaScript同时有两个线程，一个线程在某个DOM节点上添加内容，另一个线程删除了这个节点，这时浏览器应该以哪个线程为准？

所以，为了避免复杂性，从一诞生，JavaScript就是单线程，这已经成了这门语言的核心特征，将来也不会改变。

为了利用多核CPU的计算能力，HTML5提出Web Worker标准，允许JavaScript脚本创建多个线程，但是子线程完全受主线程控制，且不得操作DOM。所以，这个新标准并没有改变JavaScript单线程的本质。 ---阮一峰

### 多线程的浏览器
它可以帮助JS实现异步操作, 比如Ajax

### 定时器是否异步?
那么我们日常使用的 setTimeout() setInterval() 是否是异步呢?
例子一
```js
for (var i = 1; i <= 3; i++) {
	setTimeout(function(){
		console.log(i);
	}, 0);
}
```
//结果输出 4 , 4 , 4

例子二
```js
var start = new Date;

setTimeout(function(){
	var end = new Date;
	console.log('Time elapsed:', end - start, 'ms');
}, 0);
	
while (new Date - start < 5000) {}
```
//输出 Time elapsed: 5000 ms


按照案例二的逻辑 , 当 JS 运行时 应该是马上就运行 定时器内的 console.log 函数 , 但是最后的结果却是输出 5000ms 或者 5000ms以上

这是因为while循环造成了 5000ms 的阻塞 , setTimeout 和 setInterval 并非异步 , 而是将其注册在队列中 当JS脚本完全运行完后 , 才会执行队列 , 所以 while 将JS的运行时间拉长 导致阻塞队列内的定时


### 定时器为什么让我们看上去像是多线程

在调用 setTimeout 的时候，会有一个延时事件排入队列。然后setTimeout 调用之后的那行代码运行，接着是再下一行代码，直到再也没有任何代码。这时 JavaScript 虚拟机才会问："队列里都有谁啊？", 如果队列中至少有一个事件适合于"触发"（就像 1000 毫秒之前设定好的那个为期 500 毫秒的延时事件），则虚拟机会挑选一个事件，并调用此事件的处理器（譬如传给 setTimeout 的那个函数）。事件处理器返回后，我们又回到队列处。输入事件的工作方式完全一样：用户单击一个已附加有单击事件处理器的 DOM（Document Object Model，文档对象模型）元素时，会有一个单击事件排入队列。但是，该单击事件处理器要等到当前所有正在运行的代码均已结束后（可能还要等其他此前已排队的事件也依次结束）才会执行。因此，使用 JavaScript 的那些网页一不小心就会变得毫无反应。