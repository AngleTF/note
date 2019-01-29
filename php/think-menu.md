### tp5.0手册地址 
https://www.kancloud.cn/manual/thinkphp5/118003

### tp5.1手册地址
https://www.kancloud.cn/manual/thinkphp5_1/353946

### 先决条件
1. 了解命名空间
2. 了解composer自动加载机制
3. 了解psr4规范
4. 熟悉PDO,MYSQLI等
5. 了解smarty模板引擎机制

### 5.0和5.1差异
1. 目录结构调整
2. 命名空间调整(think\facade\....)
3. 配置文件获取调整为二级获取方式`config('app.app_debug')`
4. 取消内置常量, 通过APP类获取`|THINK_VERSION| App::version()|`
5. 取消配置项`app_namespace`, 改成环境变量


### 规范
1. 目录使用小写 + 下划线命名
2. 类库、函数文件统一以.php为后缀
3. 类的文件名均以命名空间定义，并且命名空间的路径和类库文件所在路径一致
4. 类文件采用驼峰法命名(首字母大写), 其它文件采用小写+下划线命名
5. 类名和类文件名保持一致，统一采用驼峰法命名
6. 类的命名采用大驼峰法 (User , UserType), 不需要添加后缀
7. 函数的命名使用小写字母和下划线(小写字母开头)的方式, 例如 get_client_ip
8. 方法的命名使用驼峰法(首字母小写) , 例如 getUserName
9. 属性的命名使用驼峰法(首字母小写), 例如 tableName
10. 常量以大写字母和下划线命名, 例如 APP_PATH
11. 配置参数以小写字母和下划线命名, 例如 url_route_on
12. 数据表和字段采用小写加下划线方式命名, 并注意字段名不要以下划线开头, 例如 think_user 表和 user_name字段



### 入口文件
```
项目根目录/public/index.php/modelName/controllerName/actionName
```



### 控制器
