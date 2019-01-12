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

### yum安装

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
DirectoryIndex index.php index.html index.htm	#(默认打开的FILES)
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


