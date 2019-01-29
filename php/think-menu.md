### tp5.0手册地址 
https://www.kancloud.cn/manual/thinkphp5/118003

### tp5.1手册地址
https://www.kancloud.cn/manual/thinkphp5_1/353946

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
4. 类文件采用驼峰法命名（首字母大写），其它文件采用小写+下划线命名
5. 类名和类文件名保持一致，统一采用驼峰法命名