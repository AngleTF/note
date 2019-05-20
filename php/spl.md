### 什么是SPL?
SPL 标准库(Standard PHP Library)是用于解决典型问题(standard problems)的一组接口与类的集合, 从PHP5.3开始SPL不能被禁掉.

### 数据结构类

双向链表`SplDoublyLinkedList`, 实现了迭代器,及数组的访问 可使用`foreach`进行遍历

![](/assets/78200a9bda1da528ae53bd6a77d4b7a6.jpg)

```php
$spl = new SplDoublyLinkedList();

//推送数据
$spl->push('tao');
$spl->push('li');
$spl->push('feng');
$spl->push('123');

//删除了123
$spl->pop();

//重置到开头, 调用指针操作第一次必须rewind指向节点头
$spl->rewind();

$spl->next();
$spl->prev();
$spl->next();

//检测是否溢出(有效)
if(!$spl->valid()){
    echo "invalid";die;
}

//如果溢出current会返回NULL
//如果指定的节点被pop, 在返回current, 同样会返回NULL
//current, bottom, top代表不同方向的指针, 分别代表当前, 头部, 尾部
var_dump(
    $spl->current(), count($spl), $spl, $spl->bottom(), $spl->top()
);
```

---
队列`SplQueue`继承了`SplDoublyLinkedList`, 可以使用 `SplDoublyLinkedList`的所有方法, 同时新增`enqueue` 和 `dequeue`两个方法, 用于出队和入队.
![](/assets/c6f77fa3f8635f2f070e10607b3a02b0.jpg)
```php
//队列 先进先出
$spl = new SplQueue();

//入队
$spl->enqueue('a');
$spl->enqueue('b');
$spl->enqueue('c');

$spl->rewind();
$spl->next();

//b
echo $spl->current() . "\n";

//删除了a, 出队列操作
$spl->dequeue();

var_dump($spl);
```

---
堆栈`SplQueue`继承了`SplDoublyLinkedList`, 可以使用 `SplDoublyLinkedList`的所有方法.
![](/assets/a61b4edfed878927252207f9571ca00f.jpg)
```php
//堆栈 先进先出 后进后出, 和正常的操作结果相反
$spl = new SplStack();
$spl->push('a');
$spl->push('b');
$spl->push('c');


//c
$spl->rewind();
var_dump($spl->current());

//b
$spl->next();
var_dump($spl->current());

//删除c
$spl->pop();

//a
var_dump($spl->bottom());

//b
var_dump($spl->top());
```
