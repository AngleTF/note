### 前端层
1. 使用worker进行一些大数据的运算\(防止线程阻塞 , 导致用户体验差\)
2. 前端启用应用缓存CACHE MANIFEST增加用户的体验
3. 将CSS放在页面最上方, JS放在页面最下方, 因为浏览器下载完所有CSS后才会进行渲染工作, 如果JS优先于CSS进行下载, 那么就会阻塞或者减慢浏览器渲染的进度
 
### 请求层
1. 使用精灵图
2. 压缩js , html , css... 等静态文件
3. 提交数据使用ajax, 减少不必要的数据请求
4. 使用CDN, 对请求进行减流, 能快速响应用户的请求
5. 使用懒加载, 用户未使用的数据不请求
6. 使用http的Gzip压缩和分块传输
7. 合并请求(合并图片, css, js)
8. 对不常更新的css,js, logo图片设置Http的cache-control和expire的头信息
9. 减少cookie数据传输, 因为cookie存在于每次请求中, 不宜使用过大的cookie数据

### 服务层
1. webService使用Nginx\(异步非阻塞\) , 启用适当的worker分支 , 配置合适的事件驱动模块\(epoll库\) , 开启防止惊群现象 \(accept\_mutex on\)
2. 适当的对服务器配置进行升级
3. 使用 mecache / redis 对不常更改的数据进行缓存
4. 配置适当的mysql配置 , 开启慢查询日志, 对慢查询SQL语句进行 review
5. 对数据库的读写分离

 

