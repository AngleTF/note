### PHP扩展下载仓库
```
https://pecl.php.net/index.php
```
### windows下

![](/assets/下载.jpg)
1. 从扩展库下载自己需要的 DLL 扩展
2. 然后将 DLL 文件放入PHP的 ext 目录
3. 修改 php.ini 配置文件 添加一条 extension=xxx.dll
4. 重启 php-cgi 或者 nginx , apache 就能生效

### Linux下
一般源码包中都带有所有的扩展

1. 从扩展库下载自己需要的 tar 包, 比如apcu
2. wget https://pecl.php.net/get/apcu-5.1.11.tgz
3. tar -zxvf apcu-5.1.11.tgz
4. cd apcu-5.1.11
5. 通过 phpize 命令生成 configure 文件
6. ./configure 
7. make && make install
8. 修改 php.ini 文件, extension = '/.../apcu.so'
9. nginx -s reload

注:: 有一些在configure时是需要参数的