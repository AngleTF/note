### 什么是SPL?
SPL 标准库(Standard PHP Library)是用于解决典型问题(standard problems)的一组接口与类的集合, 从PHP5.3开始SPL不能被禁掉.

### 数据结构

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

### 迭代器
`ArrayIterator`是针对数组的迭代器, 可以使用很多高级的操作
```
//直接实例化ArrayIterator
$arrayIter1 = new ArrayIterator([3, 1, 2]);

//更具value排序, 此外还有很多的排序方式
$arrayIter1->asort();

//实例化
foreach ($arrayIter1 as $k => $v){
    printf("key:%s\nvalue:%s\n", $k , $v);
}
print_r("\n\n\n");


//实例化ArrayObject
$obj = new ArrayObject(['name'=>'tao', 'age' => 12]);

//获取ArrayIterator类
$arrayIter2 = $obj->getIterator();
foreach ($arrayIter2 as $k => $v){
    printf("key:%s\nvalue:%s\n", $k , $v);
}
print_r("\n\n\n");
```

---
`AppendIterator`合并多个迭代器, 类似于`array_merge`方法
```php
//合并两个ArrayIterator
$append = new AppendIterator();
$append->append($arrayIter1);
$append->append($arrayIter2);

foreach ($append as $k => $v){
printf("key:%s\nvalue:%s\n", $k , $v);
}
```

---
`MultipleIterator`组合多个迭代器
```php
$i1 = new ArrayIterator([1, 2, 3]);
$i2 = new ArrayIterator(['taolifeng', 'zhangshan', 'lisi']);
$i3 = new ArrayIterator([22, 23, 34]);


$mult = new MultipleIterator(MultipleIterator::MIT_KEYS_ASSOC);

$mult->attachIterator($i1, 'id');
$mult->attachIterator($i2, 'name');
$mult->attachIterator($i3, 'age');

foreach ($mult as $v){
    print_r($v);
}
```

---
`RecursiveArrayIterator`迭代器, 一般用于树形结构的遍历, 该类实现了RecursiveIterator, 同时新增hasChildren, 和getChildren两个方法
```php

$myArray = array(
    0 => 'a',
    1 => array('subA', 'subB', array(0 => 'subsubA', 1 => 'subsubB', 2 => array(0 => 'deepA', 1 => 'deepB'))),
    2 => 'b',
    3 => array('subA', 'subB', 'subC'),
    4 => 'c'
);

//实例化 RecursiveArrayIterator类
$iterator = new RecursiveArrayIterator($myArray);


//树结构的遍历
function traverseStructure($iterator)
{
    while ($iterator->valid()) {
        if ($iterator->hasChildren()) {
            traverseStructure($iterator->getChildren());
        } else {
            echo $iterator->key() . ' : ' . $iterator->current() . PHP_EOL;
        }
        $iterator->next();
    }
}

traverseStructure($iterator);
```


---
`FilesystemIterator`和`DirectoryIterator`是针对文件的迭代器, 迭代返回的FileInfo类, 里面封装了很多方法, 值得注意的是资源的释放.

```php
//文件系统迭代, 继承DirectoryIterator
$fileIterator = new FilesystemIterator('.');

foreach ($fileIterator as $fileInfo) {
    //返回 SplFileInfo类
    print_r($fileInfo->getFileName() . "\n");
}
print_r("\n\n\n");


//DirectoryIterator会打印 . 和 ..
$dirIterator = new DirectoryIterator('.');
foreach ($dirIterator as $fileInfo) {
    //返回 SplFileInfo类
    print_r($fileInfo->getFileName() . "\n");
}
print_r("\n\n\n");

//单独使用SplFileInfo
$fileInfo = new SplFileInfo('ArrayAccess.php');
print_r($fileInfo->getFilename());


//关闭文件
print "Declaring file object\n";
$file = new SplFileObject('example.txt');

print "Trying to delete file...\n";
//unset($file) 销毁变量, 及时释放资源, 防止无法删除文件
unlink('example.txt');
```

### 接口
`ArrayAccess`接口用于对象像数组的形式访问
```php
class myObj implements ArrayAccess
{


    public $arr = [

        'one' => 'hello',
        'two' => 'word',
        'three' => 'this',
        'four' => 'is',
        'five' => 'blog',

    ];

    function offsetExists($offset)
    {
        // TODO: Implement offsetExists() method.
        return isset($this->arr[$offset]);
    }


    function offsetGet($offset)
    {
        // TODO: Implement offsetGet() method.
        return $this->offsetExists($offset) ? $this->arr[$offset] : null;

    }

    function offsetSet($offset, $value)
    {
        // TODO: Implement offsetSet() method.

        if (is_null($offset)) {
            $this->arr[] = $value;
        } else {
            $this->arr[$offset] = $value;
        }
    }

    public function offsetUnset($offset)
    {
        // TODO: Implement offsetUnset() method.
        unset($this->arr[$offset]);

    }

}


$obj = new myObj;

echo $obj['one'];
```
---
`Countable`接口用于count一个对象
```
class Demo implements Countable{

    public $arr = [];

    function __construct($arr)
    {
        $this->arr = $arr;
    }


    //实现count方法
    public function count()
    {
        // TODO: Implement count() method.

        return count($this->arr);
    }
}

echo count(new Demo([1,2,3]));
```

---
`IteratorIterator`接口自己写迭代规则
```php
//如果在迭代中想进行一些处理, 可以继承IteratorIterator它是OuterIterator的实现
class Demo extends IteratorIterator
{

    //覆写key方法
    public function key()
    {
        return "header--" . parent::key(); // TODO: Change the autogenerated stub


    }

    //覆写current方法
    public function current()
    {
        return "tail--" . parent::current(); // TODO: Change the autogenerated stub
    }
}

$demo = new Demo(new ArrayIterator(["a", "b", "c"]));

foreach ($demo as $k => $v){
    print_r($k . $v . "\n");
}
```


