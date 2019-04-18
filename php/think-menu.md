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

不带任何参数 自动定位当前模块 **view/类名/操作.html** 的模板文件, 以上模板路径是 ** admin/view/member/edit.html **, 并且将模板对应的think变量替换成php

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

通过 [http://www.tp51.com/index/service.index/dump](http://www.tp51.com/index/service.index/dump) 进行访问service层的dump方法, 使用 . 来访问多层控制器

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
    public function _empty(){
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

### 请求

**实例化Request**

```php
//依赖注入方式, 使用注入需要 `use think\Request`
public function __construct(Request $request)
{
    $this->request = $request;
}


//继承think\Controller后直接使用
public function index()
{
    return $this->request->param('name');
}

//Facade调用
public function index()
{
    return think\facade\Request::param('name');
} 


//助手函数
public function index()
{
    return request()->param('name');
}
```

**请求信息获取的方法**

| 方法 | 含义 |
| --- | --- |
| host | 当前访问域名或者IP |
| scheme | 当前访问协议 |
| port | 当前访问的端口 |
| remotePort | 当前请求的REMOTE\_PORT |
| protocol | 当前请求的SERVER\_PROTOCOL |
| contentType | 当前请求的CONTENT\_TYPE |
| domain | 当前包含协议的域名 |
| subDomain | 当前访问的子域名 |
| panDomain | 当前访问的泛域名 |
| rootDomain | 当前访问的根域名（V5.1.6+） |
| url | 当前完整URL |
| baseUrl | 当前URL（不含QUERY\_STRING） |
| query | 当前请求的QUERY\_STRING参数 |
| baseFile | 当前执行的文件 |
| root | URL访问根地址 |
| rootUrl | URL访问根目录 |
| pathinfo | 当前请求URL的pathinfo信息（含URL后缀） |
| path | 请求URL的pathinfo信息\(不含URL后缀\) |
| ext | 当前URL的访问后缀 |
| time | 获取当前请求的时间 |
| type | 当前请求的资源类型 |
| method | 当前请求类型 |

**请求变量获取的方法**

| 方法 | 含义 |
| --- | --- |
| param | 获取当前请求的变量 包含 PUT/GET/POST/FEIL/SESSION/ROUTE... |
| get | 获取 $\_GET 变量 |
| post | 获取 $\_POST 变量 |
| put | 获取 PUT 变量 |
| delete | 获取 DELETE 变量 |
| session | 获取 $\_SESSION 变量 |
| cookie | 获取 $\_COOKIE 变量 |
| request | 获取 $\_REQUEST 变量 |
| server | 获取 $\_SERVER 变量 |
| env | 获取 $\_ENV 变量 |
| route | 获取 路由（包括PATHINFO） 变量 |
| file | 获取 $\_FILES 变量 |

**判断请求的类型**

| 方法 | 含义 |
| --- | --- |
| 获取当前请求类型 | method |
| 判断是否GET请求 | isGet |
| 判断是否POST请求 | isPost |
| 判断是否PUT请求 | isPut |
| 判断是否DELETE请求 | isDelete |
| 判断是否AJAX请求 | isAjax |
| 判断是否PJAX请求 | isPjax |
| 判断是否手机访问 | isMobile |
| 判断是否HEAD请求 | isHead |
| 判断是否PATCH请求 | isPatch |
| 判断是否OPTIONS请求 | isOptions |
| 判断是否为CLI执行 | isCli |
| 判断是否为CGI模式 | isCgi |

**检测变量**

```php
Request::has('id','get');
```

变量检测可以支持所有支持的系统变量, 包括get/post/put/request/cookie/server/session/env/file

**请求头信息**

```php
Request::header('user-agent');
```

### 验证器

**定义验证器**

```php
namespace app\index\validate;

use think\Validate;

class User extends Validate
{
    protected $rule =   [
        'name'  => 'require|max:25',
        'age'   => 'number|between:1,120',
        'email' => 'email',    
    ];

    protected $message  =   [
        'name.require' => '名称必须',
        'name.max'     => '名称最多不能超过25个字符',
        'age.number'   => '年龄必须是数字',
        'age.between'  => '年龄只能在1-120之间',
        'email'        => '邮箱格式错误',    
    ];

    protected $scene = [
        'edit'  =>  ['name','age'],
    ];

}
```

我们定义一个\app\index\validate\User验证器类用于User的验证  
rule属性是验证的规则  
message是验证错误的提示  
scene是验证场景, 如上是 edit 场景, 需要验证两个参数 \(name 和 age\)

**快速生成验证器**

```
php think make:validate index/User
```

**实例化验证器使用**

```php
namespace app\index\controller;

use think\Controller;

class Index extends Controller
{
    public function index()
    {
        $data = [
            'name'  => 'thinkphp',
            'email' => 'thinkphp@qq.com',
        ];

        $validate = new \app\index\validate\User;

        if (!$validate->check($data)) {
            dump($validate->getError());
        }

        /*
        如果使用场景
        if (!$validate->scene('edit')->check($data)) {
            dump($validate->getError());
        }
        */
    }
}
```

**控制器中使用验证器**

```php
namespace app\index\controller;

use think\Controller;

class Index extends Controller
{
    public function index()
    {
        $result = $this->validate(
            [
                'name'  => 'thinkphp',
                'email' => 'thinkphp@qq.com',
            ],
            'app\index\validate\User');

        /*
        如果使用场景
        $result = $this->validate(
            [
                'name'  => 'thinkphp',
                'email' => 'thinkphp@qq.com',
            ],
            'app\index\validate\User.edit');

        */


        if (true !== $result) {
            // 验证失败 输出错误信息
            dump($result);
        }
    }
}
```

### 响应

```php
<?php
namespace app\index\controller;

class Index
{
    public function hello($name='thinkphp')
    {
        return 'Hello,' . $name . '!';
    }
}
```

控制器返回数据后默认输出 `Html` 格式, 可以通过修改 `default_return_type`

```php
// 默认输出类型
'default_return_type'    => 'json',
```

**快捷输出方法**

```php
<?php
namespace app\index\controller;

class Index
{
    public function hello()
    {
        $data = ['name' => 'thinkphp', 'status' => '1'];
        return json($data);
    }
}
```

| 输出类型 | 快捷方法 | 对应Response类 |
| :--- | :--- | :--- |
| HTML输出 | response | \think\Response |
| 渲染模板输出 | view | \think\response\View |
| JSON输出 | json | \think\response\Json |
| JSONP输出 | jsonp | \think\response\Jsonp |
| XML输出 | xml | \think\response\Xml |
| 页面重定向 | redirect | \think\response\Redirect |
| 附件下载（V5.1.21+） | download | \think\response\Download |

### facade

门面为容器中的类提供了一个静态调用接口  
**common中的FacadeTest类**

```php
namespace app\common;

class FacadeTest
{
    public function say(){
        return "hello";
    }
}
```

**facade目录创建FacadeTest类的映射**

```php
namespace app\facade;

use think\facade;

class Test extends facade
{
    protected static function getFacadeClass()
    {
        return 'app\common\FacadeTest';
    }
}
```

1. 需要继承 think\Facade类
2. 需要实现 getFacadeClass 方法

### DB类

**静态Db连接**  
先要将 ** 账户/密码/库名 ** 填入配置文件的 database.php 文件

**动态Db连接**

```php
    public function dbConnect()
    {

        $config = [
            // 数据库类型
            'type' => 'mysql',
            // 服务器地址
            'hostname' => '127.0.0.1',
            // 数据库名
            'database' => 'backstage',
            // 用户名
            'username' => 'root',
            // 密码
            'password' => 'mysql',
            // 端口
            'hostport' => '3307',
        ];

        $link = Db::connect($config);

        dump($link->table('user')->select());
    }
```

**Db类, 预处理方式**

```
$res = Db::query("select * from book_user where id = ?", [ 1 ]);
```

### 模型

**生成模型文件**

```
php think make:model 模块名/模型名
```

**模型定义**

```php
<?php
namespace app\index\model;

use think\Model;

class User extends Model
{
    protected $pk = 'uid';        //定义主键
    protected $table = 'user'    //表名
    protected $visible = ['id'] //显示字段
    protected $hidden = ['id']  //隐藏字段
}
```

**实例化模型调用**

```php
$p = new Person();
$p->getUserInfo();
```

**静态方式模型调用**

```php
Person::getUserInfo();
```

**助手函数模型调用**

```php
model('Person')->getUserInfo();
```

**新增多个/一个**

```
$user = new User;
$user->save([
    'name'  =>  'thinkphp',
    'email' =>  'thinkphp@qq.com'
]);


$user = new User;
$list = [
    ['name'=>'thinkphp','email'=>'thinkphp@qq.com'],
    ['name'=>'onethink','email'=>'onethink@qq.com']
];
$user->saveAll($list);

//save方法返回影响的记录数
```

**更新多个/一个**

```
$stu = new Student();
$stu->save(['name' => 'tao'], ['id' => 1]);


//当数据中存在主键的时候会认为是更新操作
$user = new User;
$list = [
    ['id'=>1, 'name'=>'thinkphp', 'email'=>'thinkphp@qq.com'],
    ['id'=>2, 'name'=>'onethink', 'email'=>'onethink@qq.com'],
];
$user->saveAll($list, false);
```

**删除多个/一个**

```php
//根据主键删除
User::destroy(1);
// 支持批量删除多个数据
User::destroy('1,2,3');

//条件删除
User::where('id','>',10)->delete();
```

**获取器**

获取器模型方法命名, FieldName为数据表字段的驼峰转换

> getFieldNameAttr

例子

```
<?php
class User extends Model 
{
    public function getStatusAttr($value)
    {
        $status = [-1=>'删除',0=>'禁用',1=>'正常',2=>'待审核'];
        return $status[$value];
    }
}
```

我取出某条User数据时, 本来status在数据库中是 int类型, 但是经过的获取器, 拿到的数据, 经过上面的针对status的过滤, 拿到了字符类型

**修改器**

修改器模型方法命名, FieldName为数据表字段的驼峰转换

> setFieldNameAttr

例子

```
<?php
class User extends Model 
{
    public function setNameAttr($value)
    {
        return strtolower($value);
    }
}
```

我设置某条User数据时, 将Name的值进行小写后在插入数据库

**类型转换**

> 支持给字段设置类型自动转换，会在写入和读取的时候自动进行类型转换处理

```php
<?php
class User extends Model 
{
    protected $type = [
        'status'    =>  'integer',
        'score'     =>  'float',
        'birthday'  =>  'datetime',
        'info'      =>  'array',
    ];
}
```

| 类型 | 简介 |
| --- | --- |
| integer | 设置为integer（整型）后，该字段写入和输出的时候都会自动转换为整型。 |
| float | 该字段的值写入和输出的时候自动转换为浮点型。 |
| boolean | 该字段的值写入和输出的时候自动转换为布尔型。 |
| array | 如果设置为强制转换为array类型，系统会自动把数组编码为json格式字符串写入数据库，取出来的时候会自动解码。 |
| object | 该字段的值在写入的时候会自动编码为json字符串，输出的时候会自动转换为stdclass对象。 |
| serialize | 指定为序列化类型的话，数据会自动序列化写入，并且在读取的时候自动反序列化。 |
| json | 指定为json类型的话，数据会自动json\_encode写入，并且在读取的时候自动json\_decode处理。 |
| timestamp | 指定为时间戳字段类型的话，该字段的值在写入时候会自动使用strtotime生成对应的时间戳，输出的时候会自动转换为dateFormat属性定义的时间字符串格式，默认的格式为Y-m-d H:i:s |

**自动写入时间搓**

修改数据库配置

```
// 开启自动写入时间戳字段
'auto_timestamp' => true,
```

或者在模型里单独开启

```php
<?php
namespace app\index\model;

use think\Model;

class User extends Model
{
    protected $autoWriteTimestamp = true;
}
```

一旦配置开启的话，会自动写入create\_time和update\_time两个字段的值，默认为整型（int），如果你的时间字段不是int类型的话，可以直接使用：

```
// 开启自动写入时间戳字段
'auto_timestamp' => 'datetime'
```

系统创建新数据后会自动插入 create\_time 字段, 数据修改后会自动插入 update\_time

修改 修改/插入字段

```php
<?php
namespace app\index\model;

use think\Model;

class User extends Model 
{
    // 定义时间戳字段名
    protected $createTime = 'create_at';
    protected $updateTime = 'update_at';
}
```

如果你只需要使用create\_time字段而不需要自动写入update\_time

```php
class User extends Model 
{
    // 关闭自动写入update_time字段
    protected $updateTime = false;
}
```

**查询范围**

例子

```php
<?php
namespace app\index\model;

use think\Model;

class User extends Model
{

    public function scopeThinkphp($query)
    {
        $query->where('name','thinkphp')->field('id,name');
    }


    //方法名 按照 scope + 自定义名称进行拼接
    public function scopeAge($query)
    {
        $query->where('age','>',20)->limit(10);
    }    

}
```

控制器中使用, scope除了第一个参数, 其他都会传入 模型的scope自定义方法中, 第二个参数用来接收

```php
// 查找name为thinkphp的用户
User::scope('thinkphp')->find();
// 查找年龄大于20的10个用户
User::scope('age')->select();
// 查找name为thinkphp的用户并且年龄大于20的10个用户
User::scope('thinkphp,age')->select();
```

**模型输出**

| 操作 | 结果类型 |
| --- | --- |
| $user-&gt;toArray\(\) | 数组 |
| $user-&gt;toJson\(\) | json字符串 |

**一对一关联**  
定义单一关联

```php
<?php
namespace app\index\model;

use think\Model;

class User extends Model
{
    public function profile()
    {
        return $this->hasOne('Profile');
    }
}
```

> hasOne\('关联数据库模型 类名' \[,'外键','主键'\]\);

外键：默认的外键规则是当前模型名（不含命名空间，下同）+\_id ，例如user\_id  
主键：当前模型主键，默认会自动获取也可以指定传入

使用关联

```php
$user = User::get(1);
// 输出Profile关联模型的email属性
echo $user->profile->email;
```

关联的方法返回hasOne的对象是`关联模型的映射`, 方法调用时名字会转化成驼峰写法

定义多个关联

```php
// 查询用户昵称是think的用户
// 注意第一个参数是关联方法名（不是关联模型名）
$users = User::hasWhere('profile', ['nickname'=>'think'])->select();

// 可以使用闭包查询
$users = User::hasWhere('profile', function($query) {
    $query->where('nickname', 'like', 'think%');
})->select();
```

**一对多关联**  
定义一对多关联

```php
<?php
namespace app\index\model;

use think\Model;

class Article extends Model 
{
    public function comments()
    {
        return $this->hasMany('Comment');
    }
}
```

> hasMany\('关联模型','外键','主键'\);

使用关联

```php
$article = Article::get(1);
// 获取文章的所有评论
dump($article->comments);


//也可以进行条件搜索
dump($article->comments()->where('status',1)->select());
```

根据其关联表的条件进行反向搜索

```php
// 查询评论状态正常的文章
$list = Article::hasWhere('comments',['status'=>1])->select();
```

**多对多关联**  
定义多对多关联

```php
<?php
namespace app\index\model;

use think\Model;

class User extends Model 
{
    public function roles()
    {
        return $this->belongsToMany('Role');
    }
}
```

用户和角色就是一种多对多的关系

> belongsToMany\('关联模型类名','中间表','外键','关联键'\);

使用关联

```php
$user = User::get(1);
// 获取用户的所有角色
$roles = $user->roles;
foreach ($roles as $role) {
    // 输出用户的角色名
    echo $role->name;
    // 获取中间表模型
    dump($role->pivot);
}
```

新增关联

```php
$user = User::get(1);
// 给用户增加管理员权限 会自动写入角色表和中间表数据
$user->roles()->save(['name'=>'管理员']);
// 批量授权
$user->roles()->saveAll([
    ['name'=>'管理员'],
    ['name'=>'操作员'],
]);
```

只新增中间表

```php
$user = User::get(1);
// 仅增加管理员权限（假设管理员的角色ID是1）
$user->roles()->save(1);
```

更新中间表

```
user = User::get(1);
// 增加关联的中间表数据
$user->roles()->attach(1);

// 传入中间表的额外属性
$user->roles()->attach(1,['remark'=>'test']);

// 删除中间表数据
$user->roles()->detach([1,2,3]);
```

attach方法的返回值是一个Pivot对象实例

### 视图

可以使用助手函数 view 或者 控制器方法 fetch,

**fetch**

```php
return $this->fetch('admin@member/edit',['think' => 'php']);
```

**view助手**

```php
return view('admin@member/edit',['think' => 'php'])
```



