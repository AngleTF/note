### 主服务器

---

__修改主服务器配置__

```
#开启二进制日志
log-bin=mysql-bin    

#服务id(id不能冲突)                    
server-id=1

#屏蔽的数据库                    
binlog-ignore-db=information_schema 
binlog-ignore-db=mysql

#同步的数据库
binlog-do-db=dbname                    
sercice mysql restart
```

__主服务器添加Slave用户__

```
GRANT FILE,REPLICATION,SLAVE ON \*.\* TO 'slave\_host'@'%' IDENTIFIED BY '123456';

FLUSH PRIVILEGES
```

![](/assets/import.png)

### 从服务器

---

__修改从服务器配置__

```
log-bin=mysql-bin
server-id=2

binlog-ignore-db=information_schema    
binlog-ignore-db=mysql

#只忽略指定的表
replicate_ignore_table=tablename    
replicate-ignore-db=mysql

#只应用指定的库
replicate-do-db=Easy             
replicate-do-db=test

只应用指定的表
replicate_do_table            

log-slave-updates
slave-skip-errors=all
slave-net-timeout=60
```

__从服务器执行SQL__

```
change master to master_host='101.132.182.191' , 
master_user='test' , 
master_password='test' ,

//上面记录的file值
master_log_file='AliYun-bin.000003' ,

//上面记录的position值                
master_log_pos=1402,                                
master_port=8765;
```

__启动从服务器__
```
start slave
```
__停止从服务器__
```
stop slave
```

__验证从服务器__
```
show slave status\G;
```

| KEY | VALUE |
| :--- | :--- |
| Slave\_IO\_Running | Yes |
| Slave\_SQL\_Running | Yes |
| Read\_Master\_Log\_Pos | 600 |

`Slave\_IO\_Running`必须为YES  
`Slave\_SQL\_Running`必须为YES  
在主库上创建一个数据库，在从库查看是否存在，存在为ok

