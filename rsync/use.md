### 帮助手册
```
rsync --help
```

### 语法
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