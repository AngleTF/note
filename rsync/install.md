### 官方网站

从官方网站下载源码包, 进行安装.

```
https://rsync.samba.org/
```

### Git仓库

```
git clone git://git.samba.org/rsync.git
```

### 检验是否已安装

一般情况下, 操作系统已经自带了rsync, 所以我们在安装之前确认一下是否已经安装过了, 使用参数 --version 来确认

```
rsync --version
```

### 下载

我下载的是3.1.3版本的

```
wget https://download.samba.org/pub/rsync/src/rsync-3.1.3.tar.gz
```

### 配置, 编译, 安装

```
./configure
make && make install
```



