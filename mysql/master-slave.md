# 主服务器

---

* ### 修改主服务器配置

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

* ### **主服务器添加Slave用户**

> GRANT FILE,REPLICATION,SLAVE ON \*.\* TO 'slave\_host'@'%' IDENTIFIED BY '123456';
>
> FLUSH PRIVILEGES

![](/assets/import.png)

# 从服务器

---

* ### 修改从服务器配置

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



