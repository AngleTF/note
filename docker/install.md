### 安装Docker CE

docker 有两个可用的版本, Community Edition（CE）和Enterprise Edition（EE）。

下载docker依赖
```
yum install -y yum-utils device-mapper-persistent-data lvm2
```

设置稳定的存储库
```
yum-config-manager --add-repo https://download.docker.com/linux/centos/docker-ce.repo
```

可选：启用边缘和测试存储库
```
yum-config-manager --enable docker-ce-edge
yum-config-manager --enable docker-ce-test
```

禁用边缘或测试存储库
```
yum-config-manager --disable docker-ce-edge
yum-config-manager --disable docker-ce-test
```

下载docker ce
如果您启用了多个Docker存储库，则安装或更新而不指定版本yum install或 yum update命令将始终安装尽可能高的版本，
```
yum install docker-ce
```
或者下载旧版
```
yum -y install docker-io
```

选择版本安装
```
yum install <FULLY-QUALIFIED-PACKAGE-NAME>
```

查看全部的docker包版本
```
yum list docker-ce --showduplicates | sort -r
```

启动docker
```
systemctl start docker
```

通过运行hello-world 映像验证安装是否正确。
```
docker run hello-world
```

rpm包地址
```
https://download.docker.com/linux/centos/7/x86_64/stable/Packages/
```
 

卸载docker
```
yum remove docker-ce
rm -rf /var/lib/docker
```

创建组
```
groupadd docker
```

将您的用户添加到docker组中。
```
usermod -aG docker $USER
```

开机启动
```
chkconfig docker on
```


