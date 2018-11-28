# **主服务器**

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

	



GRANT FILE,REPLICATION,SLAVE ON \*.\* TO 'slave\_host'@'%' IDENTIFIED BY '123456';

		

		FLUSH PRIVILEGES

		service mysql restart

		

		mysql&gt; show master status\G;

		\*\*\*\*\*\*\*\*\*\*\*\*\*\*\*\*\*\*\*\*\*\*\*\*\*\*\* 1. row \*\*\*\*\*\*\*\*\*\*\*\*\*\*\*\*\*\*\*\*\*\*\*\*\*\*\*

					 File: mysql-bin.000003			//记录

				 Position: 154						//记录

			 Binlog\_Do\_DB: Easy

		 Binlog\_Ignore\_DB: information\_schema,mysql

		Executed\_Gtid\_Set: 

		1 row in set \(0.00 sec\)



		ERROR: 

		No query specified

	

	从服务器

	

		\(从服务器配置文件里设置\)  

		log-bin=mysql-bin

		server-id=2

		

		binlog-ignore-db=information\_schema	

		binlog-ignore-db=mysql



		replicate\_ignore\_table=tablename	只忽略指定的表

		replicate-ignore-db=mysql

		

		replicate-do-db=Easy			 	只应用指定的库

		replicate-do-db=test

		replicate\_do\_table					只应用指定的表



		log-slave-updates

		slave-skip-errors=all

		slave-net-timeout=60

	

	执行

	

		change master to master\_host='101.132.182.191' , 

		master\_user='test' , 

		master\_password='test' ,

		master\_log\_file='AliYun-bin.000003' ,				//上面记录的file值

		master\_log\_pos=1402,								//上面记录的position值

		master\_port=8765;

		

		start slave		//启动

		

		stop slave 		//停止

		

		show slave status\G;

	

		Slave\_IO\_Running: Yes //此状态必须YES

		Slave\_SQL\_Running: Yes //此状态必须YES

		Read\_Master\_Log\_Pos: 600 //\#同步读取二进制日志的位置，大于等于&gt;=Exec\_Master\_Log\_Pos

		7、在主库上创建一个数据库，在从库查看是否存在，存在为ok。

	

	

	

	\#报错

	Expression \#1 of ORDER BY clause is not in GROUP BY clause and contains nonaggregated column 'information\_schema.PROFILING.SEQ' which is not functionally dependent on columns in GROUP BY clause; this is incompatible with sql\_mode=only\_full\_group\_by

	\#处理方式

	show variables like '%sql\_mode%'\G;

	set sql\_mode=\(select REPLACE\(@@sql\_mode,'ONLY\_FULL\_GROUP\_BY',''\)\);

