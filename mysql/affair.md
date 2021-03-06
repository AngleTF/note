### 事务的作用
保证操作的一致性 要么全部成功 要么全部失败, 比如说银行转钱的例子, A向B转了100块, A的银行卡的钱应该-100, B的银行卡应该+100, 这就同时操作了2条数据, 如果A刚扣完100, 突然银行的mysql(假设)挂了, 在重启mysql后会自动回滚未完成的事务, 这样A和B的金额就是最开始的样子

### 事务的作用场景
1. 多条数据操作时
2. 用户的一个提交需要处理多个业务


### 支持事务的存储引擎
1. InnoDB
2. BDB

### 事务的ACID特性
|名称|描述|
|---|---|
|原子性(Atomic)|一个事务必须被视为一个不可分割的最小工作单元，整个事务中的所有操作要么全部提交成功，要么全部失败回滚|
|一致性(Consist)|如果事务处理中途失败 , 那么不会影响任何数据|
|隔离性(Isolated)|通常来说，一个事务所做的修改在最终提交以前，对其他事务是不可见的(具体得看事务隔离级别)|
|持久性(Durable)|一旦事务提交，则所做修改就会被永久保存到数据库中|

### 事务的隔离级别
|隔离名称|会出现的情况|
|---|---|
|读未提交(read uncommitted)|脏读 , 不可重复读 , 幻读|
|读以提交(read committed)|不可重复读 , 幻读|
|可重复读(repeatable read)|幻读|
|可串行化(serializable)| |
多个线程开启各自的事务操作数据库中的数据时, 数据库系统要负责隔离操作 , 以保证各个线程在获取数据时的准确性
<br>

|情况|描述|
|---|---|
|脏读|用户A在事务中还未提交数据,被事务B获取了|
|不可重复读|同一查询在同一个事务中多次进行, 由于其他事务的"修改或者删除"提交导致返回不同的结果集|
|幻读|同一查询在同一个事务中多次进行, 由于其他事务的"插入"提交导致返回不同的结果集|

在MySQL数据库中默认的隔离级别为Repeatable read (可重复读)
### 如何操作和使用事务
关闭自动提交
```sql
set autocommit = 0;
```

开启事务
```sql
START TRANSACTION;
```

设置回滚点
```sql
SAVEPOINT name;
```

回滚至保存点
```sql
ROLLBACK TO name;
```

回滚全部
```sql
ROLLBACK;
```

提交事务
```sql
COMMIT;
```

获取事务的隔离性
```sql
select @@tx_isolation
```





