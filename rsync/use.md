### 帮助手册
```
rsync --help
```

### 使用语法
本地备份
```
rsync [option...] src... [desc]
```

shell远程备份, 抓取
```
rsync [option...] [user@]host:src... [desc]
```

shell远程备份, 推送
```
rsync [option...] src... [user@]host:[desc]
```

守护进程远程备份, 抓取
```
rsync [option...] [user@]host::src... [desc] 
//或者
rsync [option...] rsync://[user@]host[:port]/src... [desc] 
```

守护进程远程备份, 推送
```
rsync [option...] src... [user@]host::[desc]
rsync [option...] src... rsync://[user@]host[:port]/dest
```
注意: 进程守护是 :: (两个冒号) 而shell备份是 : (一个冒号)

### 脚本免密码
将本机的id_rsa.pub的内容放入需要远程主机的用户中