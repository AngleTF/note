### 引入自动加载, 实例化路由类

```php
namespace app;

include_once 'psr4.php';

use app\route;

$route = new route\Route();
```

### 定义rule方法, 如果规则匹配则调用app\Demo的dump方法

```php
$route->rule('/index/hello','app\Demo::dump',['method' => 'Get', 'port' => 80, 'scheme' => 'http']);


class Demo{
    public function dump(){
        echo "success";
    }
}
```

### 浏览器测试

![](/assets/2019-02-11_161818.png)  
使用 `/index/hello` 成功访问

![](/assets/2019-02-11_161931.png)  
使用 `/index/hello1` 访问失败

### Route的 PATH\_INFO 简单实现

```php
namespace app\route;

class Route
{
    public $allowMethod = ['GET', 'POST', 'PUT', 'DELETE'];
    public $allowScheme = ['HTTP', 'HTTPS'];

    public $defaultOption = [
        'method' => '*',
        'port' => '*',
        'scheme' => '*',
        'couper' => '/'
    ];

    public $requestOption = [];

    public function   __construct()
    {
        $this->setRequestOption();
    }


    public function rule($route, $rule, $option = [])
    {
        $option = array_merge($this->defaultOption, $option);

        if ($this->checkMethod(strtoupper($option['method'])) &&
        $this->checkScheme(strtoupper($option['scheme'])) &&
        $this->checkPort(strtoupper($option['port'])) &&
        $this->checkRule($route)){
            $this->disposer($rule);
        }
    }


    private function checkMethod($method)
    {

        if ($method == '*') {
            return true;
        }

        if (!in_array($method, $this->allowMethod)) {
            return false;
        }

        if ($this->requestOption['method'] == $method) {
            return true;
        }

        return false;
    }

    private function checkScheme($scheme){

        if ($scheme == '*') {
            return true;
        }

        if (!in_array($scheme, $this->allowScheme)) {
            return false;
        }

        if ($this->requestOption['scheme'] == $scheme) {
            return true;
        }

        return false;
    }

    private function checkPort($port){

        if ($port == '*') {
            return true;
        }

        if ($this->requestOption['port'] == $port) {
            return true;
        }

        return false;
    }

    private function checkRule($rule){
        $pass = true;
        $couper = $this->defaultOption['couper'];

        $rule_exp = explode($couper, trim($rule,' /'));

        $rule_req = explode($couper, trim($this->requestOption['path_info'],' /'));

        if(count($rule_exp) != count($rule_req)){
            return false;
        }

        foreach ($rule_req as $k => $v){
            if($v != $rule_exp[$k]){
                $pass = false;
            }
        }
        //var_dump($rule_exp, $rule_req, $pass);

        return $pass;
    }

    private function setRequestOption()
    {
        $this->requestOption = [
            'method' => $_SERVER['REQUEST_METHOD'],
            'scheme' => strtoupper($_SERVER['REQUEST_SCHEME']),
            'port' => $_SERVER['SERVER_PORT'],
            'path_info' => $_SERVER['PATH_INFO']
        ];
    }

    private function disposer($rule)
    {
        switch (true){
            case is_string($rule):
                call_user_func($rule);
                break;
            case $rule instanceof \Closure:
                $res =  $rule();
                if(is_string($res)){
                    echo $res;
                }
                break;
        }
    }
}
```



