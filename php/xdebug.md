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

### PHPSTORM + Xdebug (全局配置, 所有目录均可使用)

**PHPStorm设置端口, 最好不要是9000端口, 因为和cgi默认端口冲突**
![](/assets/wx0511.png)

![](/assets/wx_20190511192512.png)

**设置PHP的PE文件**
![](/assets/wx_20190511192634.png)
![](/assets/wx_20190511192710.png)

**设置一个全局配置**
![](/assets/wx_20190511192928.png)

**测试Xdebug**

注意: 右上角的爬虫监听电话一定要开启, 如果开启失败看看是不是端口冲突, 修改php.ini的xdebug配置和phpStorm的xdebug端口

![](/assets/wx_20190511193033.png)

**输入浏览器地址**
![](/assets/wx_20190511193056.png)

![](/assets/wx_20190511193154.png)
浏览器请求后会自动打开调试器