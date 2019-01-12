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

1. 从扩展库下载自己需要的 tar 包
2. wget https://pecl.php.net/get/apcu-5.1.11.tgz
3. 解压tar包
4. tar -zxvf apcu-5.1.11.tgz
5. cd apcu-5.1.11
6. 通过 phpize 生成 configure 文件
7. phpize
8. 运行configure文件
9. ./configure
10. 编译 , 安装
11. make && make install
12. 重启 webservice
13. nginx -s reload