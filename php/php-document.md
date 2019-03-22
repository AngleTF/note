### phpdocumentor的作用和介绍
PHPDocumentor是一个用PHP写的工具，对于有规范注释的php程序，它能够快速生成具有相互参照，索引等功能的API文档。老的版本是phpdoc，从1.3.0开始，更名为phpDocumentor，新的版本加上了对php5语法的支持，同时，可以通过在客户端浏览器上操作生成文档，文档可以转换为PDF,HTML,CHM几种形式，非常的方便。

### 下载和安装

```
D:\WMAP\php-5.6.32>pear channel-discover pear.phpdoc.org
Adding Channel "pear.phpdoc.org" succeeded
Discovery of channel "pear.phpdoc.org" succeeded

D:\WMAP\php-5.6.32>pear install phpdoc/phpDocumentor
downloading phpDocumentor-2.9.0.tgz ...
Starting to download phpDocumentor-2.9.0.tgz (16,175,113 bytes)
........done: 16,175,113 bytes
install ok: channel://pear.phpdoc.org/phpDocumentor-2.9.0
```
### 验证
```
D:\WMAP\php-5.6.32>phpdoc -V
phpDocumentor version 2.9.0
```

### 添加PHP_PEAR_BIN_DIR环境变量
PHP_PEAR_BIN_DIR = you phpdoc path

不进行这一步退出后会出现
Could not open input file: \phpdoc.php

### 生成文档
```
phpdoc -f hello.php -t doc

-d 对于目录输入
-f 对于文件输入

-t 文档输出

--template 选择模板 (zend/checkstyle/clan/responsive-twig)
```


### 生成文档过程中的错误
```
PHPDocumentor - Graphviz 'dot' command not found
```
解决此错误的步骤如下：

1. 从 https://graphviz.gitlab.io/_pages/Download/Download_windows.html 下载zip解压到 c:\some\where\graphviz
2. 添加c:\some\where\graphviz\bin到您的环境变量PATH
3. 运行phpdoc