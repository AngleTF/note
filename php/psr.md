### psr规范
```
PSR-0 --- 自动加载标准(弃用)
PSR-1 --- 基础编码标准
PSR-2 --- 代码风格向导
PSR-3 --- 日志接口
PSR-4 --- 自动加载增强版, 替换PSR-0
``` 

[psr规范文档](https://www.kancloud.cn/thinkphp/php-fig-psr/3139)


### psr4自动加载机制

目录结构
```
psr
├─app                    
│  ├─controller            
│      ├─Index.php            
|
├─psr4.php
```

psr4.php
```php
<?php
namespace psr;

define('BASE_DIR', __DIR__);

class AutoLoader{

    public static function register($class_name){

        $psr4 = [
            'app' => BASE_DIR . '/app'
        ];

        $sub_space_list = explode('\\', $class_name);

        $first_sub_space = array_shift($sub_space_list);
        if($first_sub_space === null){
            return;
        }

        if(!isset($psr4[$first_sub_space])){
            return;
        }

        $base = $psr4[$first_sub_space];

        //类名第一个字母大写
        $last_sub_space = array_pop($sub_space_list);
        if($last_sub_space === null){
            return;
        }

        array_push($sub_space_list, ucfirst($last_sub_space));

        $target = join('/', $sub_space_list);

        include_once "{$base}/{$target}.php";
    }
}

spl_autoload_register(__NAMESPACE__ . "\\AutoLoader::register");


var_dump(new \app\controller\Index());
```
其核心原理就是通过解析命名空间, 拼接路径并且引入