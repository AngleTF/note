### Composer是什么
Composer是 PHP 用来管理依赖（dependency）关系的工具。你可以在自己的项目中声明所依赖的外部工具库（libraries），Composer 会帮你安装这些依赖的库文件。

### Composer解决了什么?
1. 没有统一的资源仓库,到处乱找
2. 没有统一的安装方式.rar,zip,tar各种包都有.下载后自己得整理.
3. 遇到库的依赖关系,得自己再次下载解决.


### windows安装 Composer
PHP > 5.3
下载并且运行 [Composer-Setup.exe](https://getcomposer.org/Composer-Setup.exe)，它将安装最新版本的 Composer ，并设置好系统的环境变量，因此你可以在任何目录下直接使用 `composer` 命令。

### Composer架构思路
![](/assets/composer.png)

### 安装 ThinkPHP 为例
```linux
cd C:\Users\Administrator\Desktop\
mkdir test
﻿cd test
﻿```

使用Think核心包的 [composer.json](https://github.com/top-think/think/blob/5.1/composer.json) 放入test目录中

### 生成ThinkPHP核心文件
```
C:\Users\Administrator\Desktop\test>composer install
Loading composer repositories with package information
Updating dependencies (including require-dev)
Package operations: 2 installs, 0 updates, 0 removals
 - Installing topthink/think-installer (v1.0.12): Downloading (100%)
 - Installing topthink/framework (v5.1.8): Downloading (100%)
Writing lock file
Generating autoload files
```

### 目录结构

```
thinkphp        //TP核心文件

composer.json   //composer 下载列表

composer.lock   //composer 锁文件

vendor          //autoload 文件
```

### 更多
关于composer.json的写法请参考 http://docs.phpcomposer.com/01-basic-usage.html

### 官方网站
英文
https://packagist.org
中文
http://www.phpcomposer.com


### 自动加载
除了库的下载，Composer 还准备了一个自动加载文件，它可以加载 Composer 下载的库中所有的类文件。使用它，你只需要将下面这行代码添加到你项目的引导文件中
```php
require 'vendor/autoload.php';
```