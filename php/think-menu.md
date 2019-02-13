 ### tp5.0手册地址

[https://www.kancloud.cn/manual/thinkphp5/118003](https://www.kancloud.cn/manual/thinkphp5/118003)

### tp5.1手册地址

[https://www.kancloud.cn/manual/thinkphp5\_1/353946](https://www.kancloud.cn/manual/thinkphp5_1/353946)

### 先决条件

1. 了解命名空间
2. 了解composer自动加载机制
3. 了解psr4规范
4. 熟悉PDO,MYSQLI等
5. 了解smarty模板引擎机制

### 5.0和5.1差异

1. 目录结构调整
2. 命名空间调整\(think\facade....\)
3. 配置文件获取调整为二级获取方式`config('app.app_debug')`
4. 取消内置常量, 通过APP类获取`|THINK_VERSION| App::version()|`
5. 取消配置项`app_namespace`, 改成环境变量

### 规范

1. 目录使用小写 + 下划线命名
2. 类库、函数文件统一以.php为后缀
3. 类的文件名均以命名空间定义，并且命名空间的路径和类库文件所在路径一致
4. 类文件采用驼峰法命名\(首字母大写\), 其它文件采用小写+下划线命名
5. 类名和类文件名保持一致，统一采用驼峰法命名
6. 类的命名采用大驼峰法 \(User , UserType\), 不需要添加后缀
7. 函数的命名使用小写字母和下划线\(小写字母开头\)的方式, 例如 get\_client\_ip
8. 方法的命名使用驼峰法\(首字母小写\) , 例如 getUserName
9. 属性的命名使用驼峰法\(首字母小写\), 例如 tableName
10. 常量以大写字母和下划线命名, 例如 APP\_PATH
11. 配置参数以小写字母和下划线命名, 例如 url\_route\_on
12. 数据表和字段采用小写加下划线方式命名, 并注意字段名不要以下划线开头, 例如 think\_user 表和 user\_name字段

### 入口文件

```
项目根目录/public/index.php/modelName/controllerName/actionName
```

### 视图

可以使用助手函数 view 或者 控制器方法 fetch,

fetch

```php
return $this->fetch('admin@member/edit',['think' => 'php']);
```

view助手

```php
return view('admin@member/edit',['think' => 'php'])
```

不带任何参数 自动定位当前模块 **view/类名/操作.html** 的模板文件, 以上模板路径是 ** admin/view/member/edit.html **, 并且将模板对应的think变量替换成php

### 模型

实例化模型之前先要将 ** 账户/密码/库名 ** 填入配置文件的 database.php 文件

实例化模型调用

```php
$p = new Person();
$p->getUserInfo();
```

静态方式调用

```php
Person::getUserInfo();
```

助手函数调用

```php
model('Person')->getUserInfo();
```

### 配置

**模块优先级**  
`惯例配置->应用配置->模块配置->动态配置`

**惯例配置**  
框架核心配置, 无需更改, 位于 thinkphp/convention.php

**应用配置**  
应用初始化的时候首先加载的公共配置文件，默认位于application/config.php\(5.0\), 5.01 版本位于 config 目录下的所有配置文件

**模块配置**  
每个模块会自动加载自己的配置文件位于 application/当前模块名/config/\*.php

**动态配置**  
通过控制器设置配置时动态的配置

**配置文件**

| 文件名 | 描述 |
| --- | --- |
| app.php | 应用配置 |
| cache.php | 缓存配置 |
| cookie.php | cookie配置 |
| database.php | 数据库配置 |
| log.php | 日志配置 |
| session.php | session配置 |
| template.php | 模板引擎配置 |
| trace.php | 页面trace配置 |
| paginate.php | 分页配置 |

**环境变量配置**  
在开发过程中，可以在应用根目录下面的.env来模拟环境变量配置，.env文件中的配置参数定义格式采用ini方式, 例如:

```ini
[database]
username = 123
password = 456
```

**读取.env的配置**

```php
Env::get('database.username', 'default_username');
```

\(.\)的目的是获取二级配置, 也就是获取 database 项的username, 第二个参数是默认值

**获取配置信息**

```php
Config::get('参数名');

//或者使用助手函数
 config('参数名')
```

**获取所有参数**

```php
Config::get();

//或者
 config();
```

**判断参数是否存在**

```php
Config::has('配置参数2');

// 或者 
config('?配置参数2');
```

**设置参数**

```php
Config::set('配置参数','配置值');

// 或者使用助手函数
config('配置参数','配置值');
```

**设置多个参数**

```php
Config::set([
    '配置参数1'=>'配置值',
    '配置参数2'=>'配置值'
]);

// 或者使用助手函数
config([
    '配置参数1'=>'配置值',
    '配置参数2'=>'配置值'
]);
```

### 路由

**路由规则文件**

```
├─route                 路由定义目录
│  ├─route.php          路由定义
│  ├─api.php            路由定义
│  └─...                更多路由定义
```

通过修改路由文件来注册相应的规则

**注册路由规则**

```php
Route::rule('路由表达式','路由地址','请求类型','路由参数','变量规则');
```

**数组形式注册路由**

```php
return [
    '路由规则' => ['路由地址', ['路由选项'], ['变量规则']]
]
```

**路由规则**

```
路由标识符/:变量名1/:变量名2/[:可选变量名1]
```

**路由地址**  
通过匹配相应的路由规则, 映射到对应的路由地址  
![](/assets/thinkphp-rule-1.png)

**路由选项**  
Route::rule方法的第三个参数, 有以下选项  
![](/assets/thinkphp-rule.png)

**路由参数效验**  
上面例子中 \['id' =&gt; '\d+'\], 则是对路由参数的正则验证

**例子**

```php
// 注册路由到index模块的News控制器的read操作
Route::rule('new/:id','index/News/read', ['ext' => 'html'],['id' => '\d+']);
```

通过访问 `http://serverName/new/5.html`, 会将其映射到`http://serverName/index/news/read/id/5` 并且原来的访问地址会自动失效。

**强制路由**

```
url_route_must 开启后无法使用pathinfo模式访问, 只能通过路由规则访问
```

**路由参数传递方式**

假如路由为

```php
Route::get('getUserInfo/:name/:age', 'index/test/index');
```

`url_param_type` 为 0, 则成对解析  
[http://serverName/getUserInfo/name/tao/age/22](http://serverName/getUserInfo/name/tao/age/22)  
getUserInfo是路由的映射, 而 name/tao/age/22 是参数名和参数值, 以 key/value 形式展现

`url_param_type` 为1, 则顺序解析  
[http://serverName/getUserInfo/tao/22](http://serverName/getUserInfo/tao/22)

打印

```php
dump(Request::param());
```

```
array (size=2)
  'name' => string 'name' (length=4)
  'age' => string '22' (length=2)
```

**变量映射**

```php
//将路由匹配到的version的值映射到跳转地址中
Route::rule('api/:version','index/:version/index');
```

**变量的全局规则定义**

```php
Route::pattern([
    'age' => '\d+'
]);
```

**路由调用控制器方法**

```php
Reute::rule('demo', '[@模块名][/控制器][/方法]')
```

**路由调用类方法**

```php
Reute::rule('demo', '命名空间@方法')
```

**路由调用回调函数\(小型方法\)**

```php
Reute::rule('demo', function(){
    return 'demo'
})
```


**路由分组**

格式
```php
Route::group('分组名或者分组参数','分组路由规则',['路由选项]',['变量规则']);
```

例子    
```php
//数组直接定义规则
Route::group('blog', [
    ':id'   => 'Blog/read',
    ':name' => 'Blog/read',
])->ext('html')->pattern(['id' => '\d+']);

//闭包方式定义规则
Route::group('blog', function () {
    Route::rule(':id', 'blog/read');
    Route::rule(':name', 'blog/read');
})->ext('html')->pattern(['id' => '\d+', 'name' => '\w+']);
```

**路由别名**
格式
```php
Route::alias('路由别名', '路由地址', ['路由选项']);
```
**示例**
```php
Route::alias('math', 'index/index');
```

index控制器,index模块, add方法
```php
function add($n1 = 0, $n2 = 0)
{
    return $n1 + $n2;
}
```

通过 `http://localhost/math/add/1/2` 进行访问


路由别名的黑白名单
```php
Route::alias('math', 'index/index', [
    //允许的方法
    'allow' => 'index,read',
    
    //禁止的方法
    'except' => 'edit,delete'
]);
```

**miss路由**
当路由没有匹配后执行的路由
```php
Route::miss('路由地址或者闭包')
```


### 控制器
在父类控制器当中，定义了相关的操作方法，以及加载了不同功能需要使用的业务类，能够一定意义的节省在子类当中的代码书写量

只有public的方法才能具有访问权限

**控制器名后缀**
修改 config/app.php 中的, 开启后文件和类名需要加后缀Controller 
```
'controller_suffix' => true
```


**多层控制器**
```
├─index                     index模块
│  ├─controller             
|      ├─service            控制器的service目录
|         ├─IndexController service目录下的Index控制器
```
通过 http://www.tp51.com/index/service.index/dump 进行访问service层的dump方法, 使用 . 来访问多层控制器

如果想把 . 转换成 / 使用, 那么只要将`'controller_auto_search' => true`即可

**分层控制器**
分层控制器是url无法直接访问到的, 直接在模块下定义新的控制器文件夹, 并且创建新文件, 由于遵循psr4规范 可直接利用命名空间实例化类, 也可以直接使用controller方法进行实例化
```php
controller('Mail', 'service')
```
实例化 service控制器下的Mail类


**空操作 / 空控制器**
```php
//控制器内, 错误的操作名调用的函数
public function _empty(){
    return '空操作';
}
```

定义一个Error的控制器, 并且定义index方法, 访问错误的控制器就能调用此方法
```php
class Error
{
    public function index(){
        return "空控制器";
    }
}
```


**控制器前置操作**
```php
protected $beforeActionList = [
    'beforeActions',        //所有方法的前置操作
    'beforeActions' => ['except' => 'index'],    //除了index外其他都能调用的前置操作
    'beforeActions' => ['only' => 'index']        //只有index方法才能调用的前置操作
];
```
覆写父类 beforeActionList  属性, 数组`key`是调用的前置方法名


**重定向 / 跳转**
系统的\think\Controller类内置了两个跳转方法success和error，用于页面跳转提示
```php
$this->success('新增成功', 'User/list', [], 3, []);

$this->error('新增失败', 'User/list', [], 3, []);
```

跳转对应的模板文件
```
//默认错误跳转对应的模板文件
'dispatch_error_tmpl' => 'public/error',
//默认成功跳转对应的模板文件
'dispatch_success_tmpl' => 'public/success',
```


重定向
```php
$this->redirect('模块/控制器/操作', ['参数'], 'httpcode', ['session data']);
```

助手函数支持跳转前记住当前url
```php
redirect('News/category')->remember();
```

使用上次记住的url
```php
redirect()->restore();
```


**url生成**
```php
Url::build('模块/控制器/操作',['参数'],['url后缀'],['域名']);

//助手函数
url('模块/控制器/操作',['参数'],['url后缀'],['域名']);
```
例子
```php
echo Url::build('index/index/index',['name' => 'tao'], 'html', true);
```
