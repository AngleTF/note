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
echo "zend_extension = \"C:/path/to/php/ext/xdebug.dll\"" >> C:/path/to/php/php.ini
```

