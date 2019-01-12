### 什么是LAMNP

Linux简称为L, Apache简称为A, Mysql简称为M, Php简称为P, Nginx简称为N

### 操作系统选择

选择CentOS7 64位作为我们本次安装的操作系统, CentOs是redHat的免费版本, 可在云服务器厂商进行购买

### 远程工具的选择

安装软件之前需要一款合适的ssh连接工具, 我一直在使用Xshell感觉还不错

### 关闭SELinux防火墙

永久关闭防火墙, 编辑config将SELINUX=enforcing 改成 SELINUX=disabled

```
vi /etc/selinux/config
```

临时关闭防火墙

```
setenforce 0
```

### 可选择yum安装或源码安装

yum安装很快, 如果你是新手并不推荐你使用yum安装, 适时使用源码安装, 了解编译安装过程还是非常必要的.

### yum安装LMAP

安装PHP, -y自动安装依赖, 不需要手动yes

```
yum install php php-devel -y
```

安装并启动Apache

```
yum install httpd -y
```

启动Apache

```
service httpd start
```

设置Apache为开机启动

```
systemctl enable httpd
```

编辑Apache配置文件

```
vi /etc/httpd/conf/httpd.conf
```

将 DocumentRoot  "/var/www/html" 改为 DocumentRoot  "/var/www" 这是网站根目录 当然你也可以不更改

```
<Directory "/var/www">
#将none改为All(允许.htaccess文件进行apache规则重写)
AllowOverride none
#Allow open access:
Require all granted
</Directory>
<IfModule dir_module>
DirectoryIndex index.php index.html index.htm    #(默认打开的FILES)
</IfModule>
```

安装MariaDB, CentOS 7.0中，已经使用MariaDB替代了MySQL数据库

```
yum install mariadb mariadb-server
```

启动MariaDB

```
service mariadb start
```

拷贝配置文件

```
cp /usr/share/mysql/my-huge.cnf /etc/my.cnf
```

初始化数据库

```
mysql_secure_installation
```

测试文件

```
vim /var/www/index.php

#写入
<?php
    phpinfo();
?>
```

修改php配置文件, 开启报错

```
display_errors = Off
error_reporting = E_ALL | E_STRICT
```

修改过配置文件后, 需要重启Apache

```
service httpd restart
```

### LMNP源码安装

首先我们先设置好目录结构

目录结构

```
    /
    +--/data/
    |    |
    |    +--www            //站点根目录
    |    |
    |    +--mysql        //mysql安装/database地址
    |    |
    |    +--nginx        //nginx安装地址
    |    |
    |    +--php            //php安装地址
    |    |
    |    +--package        //源码包
    |    |
    |    |--boost        //mysql boost c++ lib
    |    
    |
    +--/etc/
    |    |
    |    +--my.cnf        //mysql配置文件
    |    |
    |    +--nginx        //nginx配置目录
    |    |
    |    +--php.ini        //php配置文件
    |    |
    |    |
    |    +--init.d
    |        |
    |        +--php-fpm    //fpm启动
    |        |
    |        +--mysqld    //mysql启动
    |
    |--/sbin/
        |
        +--nginx        //nginx启动
```

添加需要的组合用户

```
groupadd -r mysql 
useradd -rg mysql mysql 

groupadd -r nginx
useradd -r -g nginx  nginx

groupadd -r php
useradd -rg php php

groupadd -r www 
useradd -rg www www
```

创建安装目录

```
mkdir -p /data/php /data/nginx /data/temp /data/www /data/package /data/boost
```

创建mysql的log文件, 并设置目录权限最高级

```
touch /data/temp/mysql.log
chmod 777 -R /data/temp
```

安装所有依赖

```
yum -y install cmake make gcc-c++ ncurses-devel perl-Data-Dumper vim wget
yum -y install gcc gcc-c++ automake pcre pcre-devel zlib zlib-devel openssl openssl-devel
yum -y install libmcrypt libmcrypt-devel mhash mhash-devel libxml2 libxml2-devel bzip2 bzip2-devel  autoconf curl-devel libjpeg-devel libpng-devel freetype-devel
```

下载并解压源码包

```
tar zxvf ./mysql-boost-5.7.22.tar.gz -C /data/package 
tar zxvf ./nginx-1.14.0.tar.gz -C /data/package
tar zxvf ./php-7.2.5.tar.gz -C /data/package
```

编译安装mysql, mysql 5.7.5开始Boost库是必需的, 最新版的 mysql 都使用cmake命令进行安装

```
cd /data/package/mysql-5.7.22 
cmake -DCMAKE_INSTALL_PREFIX=/data/mysql   -DSYSCONFDIR=/etc -DMYSQL_USER=mysql  -DWITH_MYISAM_STORAGE_ENGINE=1  -DWITH_INNOBASE_STORAGE_ENGINE=1 -DWITH_ARCHIVE_STORAGE_ENGINE=1  -DWITH_MEMORY_STORAGE_ENGINE=1  -DWITH_READLINE=1   -DMYSQL_UNIX_ADDR=/tmp/mysql.sock  -DMYSQL_TCP_PORT=3306  -DENABLED_LOCAL_INFILE=1  -DENABLE_DOWNLOADS=1  -DWITH_PARTITION_STORAGE_ENGINE=1  -DEXTRA_CHARSETS=all  -DDEFAULT_CHARSET=utf8  -DDEFAULT_COLLATION=utf8_general_ci  -DWITH_DEBUG=0  -DMYSQL_MAINTAINER_MOODE=0 -DWITH_BOOST=/data/package/mysql-5.7.22/boost/boost_1_59_0
make && make install
```

部分参数说明

| cmake参数 | cmake值 |
| --- | --- |
| DCMAKE\_INSTALL\_PREFIX | mysql安装路径 |
| DMYSQL\_DATADIR | mysql数据库位置 |
| DSYSCONFDIR | 默认my.cnf选项文件目录 |

初始化mysql

```
chown -R mysql:mysql /data/mysql
/data/mysql/bin/mysqld --initialize-insecure --user=mysql --basedir=/data/mysql --datadir=/data/mysql/data
cp /data/mysql/support-files/mysql.server /etc/init.d/mysqld
chmod 777 /etc/init.d/mysqld
```

将mysql/bin加入全局变量

```
echo "export PATH=\$PATH:/data/mysql/bin" >> /etc/profile
source /etc/profile
```

设置mysql为开机自启

```
chkconfig mysqld  on 
chkconfig --add mysqld
```

启动mysql

```
service mysqld start
```

编译安装nginx

```
cd /data/package/nginx-1.14.0
./configure --prefix=/data/nginx --sbin-path=/usr/sbin/nginx --conf-path=/etc/nginx/nginx.conf --error-log-path=/data/temp/nginx_error.log --http-log-path=/data/temp/nginx_access.log --user=nginx --group=nginx --with-http_ssl_module --with-http_realip_module --with-http_addition_module --with-http_sub_module --with-http_dav_module --with-http_flv_module --http-client-body-temp-path=/data/temp/client --http-proxy-temp-path=/data/temp/proxy --http-fastcgi-temp-path=/data/temp/fastcgi --with-mail --with-mail_ssl_module --with-pcre
make && make install
```

赋予nginx目录权限

```
chown -R nginx:nginx /data/nginx
```

启动nginx

```
/sbin/nginx
```

编译安装PHP

```
cd /data/package/php-7.2.5
./configure --prefix=/data/php --with-config-file-path=/etc --with-gd --with-iconv --with-zlib --enable-bcmath --enable-shmop --enable-sysvsem --enable-inline-optimization --enable-mbregex --enable-fpm --enable-mbstring --enable-ftp --with-openssl --enable-pcntl --enable-sockets --with-xmlrpc --enable-zip --enable-soap --with-png-dir --with-gettext --with-curl --with-jpeg-dir --with-freetype-dir --with-mysqli --enable-embedded-mysqli --with-pdo-mysql -enable-xml --with-bz2
make && make install
```

赋予php目录权限

```
chown -R php:php /data/php
```

将php/bin加入全局变量

```
echo "export PATH=\$PATH:/data/php/bin" >> /etc/profile
source /etc/profile
```

php-fpm 是php-cgi的管理器, php-cgi是nginx与php沟通的桥梁, 我们需要配置一下

```
cp /data/package/php-7.2.5/sapi/fpm/init.d.php-fpm /etc/init.d/php-fpm
chmod +x /etc/init.d/php-fpm

cp /data/php/etc/php-fpm.conf.default /data/php/etc/php-fpm.conf

cp /data/php/etc/php-fpm.d/www.conf.default /data/php/etc/php-fpm.d/www.conf
cp /data/package/php-7.2.5/php.ini-production /etc/php.ini
```

启动php-fpm

```
service php-fpm start
```



