### 什么是LAMNP
Linux简称为L, Apache简称为A, Mysql简称为M, Php简称为P, Nginx简称为N

### 操作系统选择
选择CentOS7 64位作为我们本次安装的操作系统, CentOs是redHat的免费版本, 可在云服务器厂商进行购买

### 远程工具的选择
安装软件之前需要一款合适的ssh连接工具, 我一直在使用Xshell感觉还不错

### 关闭SELinux防火墙
永久关闭防火墙, 编辑config将SELINUX=enforcing 改成 SELINUX=disabled
```shell
vi /etc/selinux/config
```
临时关闭防火墙
```shell
setenforce 0
```

### 可选择yum安装和源码安装
yum安装很快, 如果你是新手并不推荐你使用yum安装, 适时使用源码安装, 了解编译安装过程还是非常必要的.

### yum安装
