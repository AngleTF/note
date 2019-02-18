### 查看当前挂载硬盘

```
[root@iZbp1dai2p6qhwsx5u3aw2Z ~]# df -h
Filesystem      Size  Used Avail Use% Mounted on
/dev/vda1        40G   12G   26G  32% /
devtmpfs        1.9G     0  1.9G   0% /dev
tmpfs           1.9G     0  1.9G   0% /dev/shm
tmpfs           1.9G  492K  1.9G   1% /run
tmpfs           1.9G     0  1.9G   0% /sys/fs/cgroup
tmpfs           379M     0  379M   0% /run/user/0

```

### fdisk

![](/assets/mount-1.png)

上图可以看到有三块硬盘, 其中有一块已经挂载(vda), 其余两块还未挂载(vdb,vdc) , 分别大小是60G和500G

### 创建挂载点
```
mkdir dir /resource /spider_log
```

