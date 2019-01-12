

### 如何使用理解PHP的命名空间
从下面这个例子可以看出 n1.php 声明了一个命名空间 , 然后 n2.php 对 n1.php 进行引入 , 但是结果却报错了 , 那这就是 namespace 导致的 , 那如何对 namespace 进行理解呢?

n1命名空间(n1.php)
```php
namespace n1;

//声明常量 NAME 为 n1 命名空间的
define(__NAMESPACE__ . '\\' . 'NAME','tao');

function nPrint(){
    echo __NAMESPACE__;
}

class test{
    function printf(){
        echo __NAMESPACE__;
    }
}
```

不声明命名空(n2.php)
```php
include "n1.php";

//报错
echo NAME;

//报错
nPrint();

//报错
(new test())->printf();
```

### 个人理解

我个人对 namespace 的理解其实很简单 , 那就是给类 , 函数 , 常量 添加前缀 , 也有人理解为目录的层级关系 , 当然这样理解也是可以的 . 为什么我会怎么说呢? 我们来看看下面的几个例子 .

### 完全限定名称
不声明命名空间(n2.php)
```php
include "n1.php";

//tao
echo \n1\NAME;

//n1
\n1\nPrint();

//n1
(new \n2\test())->printf();
```

我只对 n2.php 进行了修改 , 通过对比上一次 n2.php 文件 , 不难发现 我其实就是在 类 , 函数 , 常量 前面加了一个前缀 \n1

添加 n1 可以理解 , 但为什么还要添加 \ 这个呢? 当然你可以理解一旦声明命名空间 , 命名空间都是从 \ 根开始的 所有我们使用添加前缀的 \n1 , 向这种从根开始找命名空间的我们称之为 "完全限定名称"

### 限定名称
n2\branch 命名空间(n1.php)
```php
namespace n2\branch;

//声明常量 NAME 为 n2 命名空间的
define(__NAMESPACE__ . '\\' . 'NAME','tao');

function nPrint(){
    echo __NAMESPACE__;
}

class test{
    function printf(){
        echo __NAMESPACE__;
    }
}
```

n2命名空间(n2.php)
```php
namespace n2;
include "n1.php";

//tao
echo branch\NAME;

//n1\branch
branch\nPrint();

//n1\branch
(new branch\test())->printf();
```

对 n1.php 和 n2.php 同时进行了修改 , n1.php 的命名空间是 n1\branch , 同时 n2.php 添加了一个命名空间为 n1 , 按照目录的原理 n1.php 的命名空间 比 n2.php 的命名空间更加深入(目录层级结构) , 而n2.php 只需要添加前缀 branch 就能够使用 , 这就是限定名称

### 非限定名称
n2\branch命名空间(n2.php)
```php
namespace n2\branch;
include "n1.php";

//tao
echo NAME;

//n2\branch
nPrint();

//n2\branch
(new test())->printf();
```

对n2.php 再次进行修改 , 将命名空间 声明和 n1.php 一样 , 那么就不需要再前面添加任何前缀 , 这就想等同于 同级目录的概念, 这种声明我们称之为 非限定名称 .


### USE
如果你觉得每次调用 声明类 , 调用函数 , 调用常量 都需要写入命名空间的前缀 , 那么你可以使用 关键字 USE 进行全局命名


n2\branch命名空间(n2.php)
```php
include "n1.php";
use \n1\branch as b;


//tao
echo b\NAME;

//n2\branch
b\nPrint();

//n2\branch
(new b\test())->printf();
```

又或者单独导入
```php
use \n1\branch\test;
use function n1\branch\nPrint;
use const n1\branch\NAME;

//tao
echo NAME;

//n2\branch
nPrint();

//n2\branch
(new test())->printf();
```