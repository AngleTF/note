### 什么是设计模式

设计模式（Design Pattern）是一套被反复使用、多数人知晓的、经过分类的、代码设计经验的总结。使用设计模式的目的：为了代码可重用性、让代码更容易被他人理解、保证代码可靠性。 设计模式使代码编写真正工程化；设计模式是软件工程的基石脉络，如同大厦的结构一样。

### 三种设计模式

| 模式分类 | 模式名称 |
| --- | --- |
| 创建型 | 单例 , 原型 , 建造者 , 工厂 , 抽象工厂 ... |
| 结构型 | 桥接 , 外观 , 组合 , 装饰 , 适配器 , 代理 , 享元 ... |
| 行为型 | 迭代器 , 解释器 , 观察者 , 中介者 , 访问者 , 备忘录 , 状态 , 策略 , 模板方法 , 命令 , 职责链 ... |

下面简单的介绍下创造型的5大设计模式

部分代码来自[实验楼16个PHP设计模式](https://www.shiyanlou.com/courses/699)详解

### 工厂模式

简单工厂模式最大的优点在于实现对象的创建和对象的使用分离，  
将对象的创建交给专门的工厂类负责，但是其最大的缺点在于工厂类不够灵活，  
增加新的具体产品需要修改工厂类的判断逻辑代码，而且产品较多时，工厂方法代码将会非常复杂。

```php
class duck{
    function __construct()
    {
        echo "i am duck";
    }
}

class elephant{
    function __construct()
    {
        echo "i am elephant";
    }
}

class plant{
    public static function create($name){
        if(class_exists($name)){
            return new $name;
        }else{
            die("not a class");
          }
    }
}

plant::create('elephant'); //i am elephant
plant::create('duck');   //i am duck
plant::create('1');     //not a class
```

### 抽象工厂

```php
interface zoology{
    public function run();
}

class duck implements zoology {
    public $name = 'duck';
    function __construct()
    {
        echo "i am {$this->name}";
    }
    function run()
    {
                // TODO: Implement run() method.
                echo "{$this->name} running...";
    }
}

class elephant implements zoology{
    public $name = 'elephant';
    function __construct()
    {
        echo "i am {$this->name}";
    }
    function run()
    {
        // TODO: Implement run() method.
        echo "{$this->name} running...";
    }
}

abstract class factory{
    abstract function create($name);
}

//工厂所要做的就是创建实例
class plant extends factory{
    public function create($name){
        if(class_exists($name)){
            return new $name;
        }else{
            echo("not a class");
        }
    }
}
$plant = new plant();
$plant->create('elephant'); //i am elephant
$plant->create('duck');   //i am duck
$plant->create('1');     //not a class
```

### 单例

三私一公, 一个类只创建一个对象  
私有化 构造方法, 私有化 克隆, 私有化 静态 存储的对象, 公开 获取对象的入口

```php
final class Single{

    //私有可变实例化对象数组
    private static $var = [];
    //私有构造
    private function __construct()
    {

    }
    //私有克隆
    private function __clone()
    {
        // TODO: Implement __clone() method.
    }

    //一个类只创建一个对象 , 并将其保存至上下文环境
    public static function create($key){

        if(!isset(Single::$var[$key])){
                echo "\n实例化\n";
                Single::$var[$key] = new $key;
        }
        echo "返回数据\n";
        return Single::$var[$key];
    }
}

$a = Single::create('worker');
$b = Single::create('worker');
var_dump($a , $b);
```

### 建造者

```php
abstract class Builder
{
  protected $car;
  abstract public function buildPartA();
  abstract public function buildPartB();
  abstract public function buildPartC();
  abstract public function getResult();
}

class CarBuilder extends Builder
{
  function __construct()
  {
    $this->car = new Car();
  }
  public function buildPartA(){
    $this->car->setPartA('发动机');
  }

  public function buildPartB(){
    $this->car->setPartB('轮子');
  }

  public function buildPartC(){
    $this->car->setPartC('其他零件');
  }

  public function getResult(){
    return $this->car;
  }
}

class Car
{
  protected $partA;
  protected $partB;
  protected $partC;

  public function setPartA($str){
    $this->partA = $str;
  }

  public function setPartB($str){
    $this->partB = $str;
  }

  public function setPartC($str){
    $this->partC = $str;
  }

  public function show()
  {
    echo "这辆车由：".$this->partA.','.$this->partB.',和'.$this->partC.'组成';
  }
}

class Director
{
  public $myBuilder;

  public function startBuild()
  {
    $this->myBuilder->buildPartA();
    $this->myBuilder->buildPartB();
    $this->myBuilder->buildPartC();
    return $this->myBuilder->getResult();
  }

  public function setBuilder(Builder $builder)
  {
    $this->myBuilder = $builder;
  }
}

$carBuilder = new CarBuilder();
$director = new Director();
$director->setBuilder($carBuilder);
$newCar = $director->startBuild();
$newCar->show();
```



