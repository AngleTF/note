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

### fdisk命令

![](/assets/mount-1.png)

上图可以看到有三块硬盘, 其中有一块已经挂载\(vda\), 其余两块还未挂载\(vdb,vdc\) , 分别大小是60G和500G

```
/dev/sda1
```

hd表示IDE接口, sd表示SCSI接口, a,b,c 表示第几块硬盘, 这里表示第一块, 后面的数字表示第几个分区, 1,2,3,4主分区, 5表示第一个逻辑分区

文件系统类型Ext3 / swap / vfat（fat32）/ Ext4 / Ext2

### 创建挂载点

```
mkdir dir /resource /spider_log
```

### 分区

进入分区模式命令
```
fdisk /dev/vdb
```

分区模式
```
m  //分区模式帮助
n  //添加一个新分区
d  //删除分区
p  //查看分区状态
q  //退出
w  //保存结束分区
t  //创建swap分区

   n:
     p  //添加主分区
     e  //添加扩展分区
     l  //添加逻辑分区
``` 
![](/assets/mount-3.png)
创建主分区后, 保存退出 

### 格式化分区
![](/assets/mount-2.png)

### 挂载目录
![](/assets/mount-4.png)

### 扩展知识

1. 一个盘符4个主分区(主分区和扩展分区不能超过4个 扩展分区只能有1个 扩展分区不能直接存数据)
2. 硬盘,USB,光盘等多在/dev/ 目录下	
3. 主分区 可以变为扩展分区,扩展分区可以容纳更多的逻辑分区	
4. 目录称为挂载点
5. 目录和分区的链接称为挂载
	
> sda1	第一个硬盘第一个主分区
>> sdb5	第二个硬盘第一个逻辑分区 逻辑分区必须从5开始


