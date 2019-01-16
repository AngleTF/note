### 主进程和线程如何实现通讯


worker生成线程(master.html)
```js
<script defer>

//主进程
let worker = new Worker('thread.js');

//监听线程发送的数据
worker.addEventListener('message',function (e) {
console.log("线程对象" , e);
console.log("线程数据:" + e.data);
});

//向thread.js线程发送数据
worker.postMessage("hello i am master");

</script>
```

线程(thread.js)
```js
self.addEventListener('message',function (e) {
console.log("主进程对象:" , e);
console.log("线程收到主进程信息:" , e.data);
self.postMessage('hello i am thread');
});
```