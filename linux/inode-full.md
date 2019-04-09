### 服务报错
```
no space left on device
```

### INode满了
![](/assets/inode.png)

/resource磁盘被文件INode被占满, 无法创建新的文件, 但是硬盘空间还有57G可以使用

INode是文件的索引文件信息, 占了一部分磁盘空间, 文件过多导致Inode不够用可以通过 `df -ih` 来查看所有磁盘INode的使用情况, 使用 `df -h`查看所有磁盘使用量


### 解决方法
1. 删除多余的垃圾文件来降低Inode
2. 格式化磁盘, 重新分配更多的Inode

### 格式化方式
```
#卸载文件系统
umount /dev/sda6

#建立文件系统，指定inode节点数
mkfs.ext3 /dev/sda6 -N "inode节点数"

#挂载文件系统
mout  /dev/sda6  /data
```

[注意] 调整inode数会格式化磁盘，执行前应确定磁盘上没有重要数据或是先备份数据