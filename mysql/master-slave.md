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







