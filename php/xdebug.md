### Linux, Mac安装
```
#需要先安装pecl

pecl install xdebug
```

### windows 安装

**下载**
```
https://xdebug.org/download.php
```

**移动文件到php的ext文件下**
```
mv ./xdebug.dll /path/to/php/ext/
```

**修改php.ini**
```
[xdebug]
zend_extension = "E:/WMAP/php-7.2.16-Win32-VC15-x64/ext/php_xdebug.dll"
xdebug.auto_trace=on
xdebug.collect_params=on
xdebug.collect_return=on
xdebug.idekey = PHPSTORM
xdebug.remote_enable = on
xdebug.remote_handler = dbgp
xdebug.remote_host= localhost
xdebug.remote_port = 9001
```

### PHPSTORM + Xdebug (全局配置)

PHPStorm设置端口, 最好不要是9000端口, 因为和cgi默认端口冲突
![](/assets/wx0511.png)

![](/assets/wx_20190511192512.png)

设置PHP的PE文件
![](/assets/wx_20190511192634.png)

