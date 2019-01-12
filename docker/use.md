### docker原理

镜像是类 , 容器是对象 , 通过镜像实例化容器 , 容器是虚拟的系统环境 , docker run 是实例化操作

### 修改国内镜像
```
/etc/docker/daemon.json
```
```json
{
"registry-mirrors": ["http://hub-mirror.c.163.com"]
}
```


### 启动Docker

```
docker run [OPTIONS] IMAGE [COMMAND] [ARG...]



OPTIONS


-a stdin	指定标准输入输出内容类型，可选 STDIN/STDOUT/STDERR 三项

-d	后台运行容器，并返回容器ID

-i	以交互模式运行容器，通常与 -t 同时使用

-t	为容器重新分配一个伪输入终端，通常与 -i 同时使用

--name="nginx-lb"	为容器指定一个名称

--dns 8.8.8.8	指定容器使用的DNS服务器，默认和宿主一致

--dns-search example.com	指定容器DNS搜索域名，默认和宿主一致

-h "mars"	指定容器的hostname

-e username="ritchie"	设置环境变量

--env-file=[]	从指定文件读入环境变量


--cpuset="0-2"

--cpuset="0,1,2"	绑定容器到指定CPU运行


-m	设置容器使用内存最大值

--net="bridge"	指定容器的网络连接类型，支持 bridge/host/none/container: 四种类型

--link=[]	添加链接到另一个容器

--expose=[]	开放一个端口或一组端口


-p	标识来绑定指定端口5000:5000

-P	将容器内部使用的网络端口映射到我们使用的主机上
```

### 初始化并且进入docker终端


Docker 允许你在容器内运行应用程序， 使用 docker run 命令来在容器内运行一个应用程序
```
docker run -d -p 5000:5000 training/webapp python app.py

docker run ubuntu:15.10 /bin/echo "Hello world"
```
ubuntu:15.10	15.10指定要运行的镜像，Docker首先从本地主机上查找镜像是否存在，如果不存在,Docker 就会从镜像仓库 Docker Hub 下载公共镜像, /bin/echo "Hello world"	在启动的容器里执行的命令


```
docker run -i -t ubuntu:15.10 /bin/bash
```
我们可以通过运行exit命令或者使用CTRL+D来退出容器。



### 容器ID
```
runoob@runoob:~$ docker run -d ubuntu:15.10 /bin/sh -c "while true; do echo hello world; sleep 1; done"
2b1b7a428627c51ab8810d541d759f072b4fc75487eed05812646b8534a2fe63
```
上面输出的是容器ID , 唯一



### 列出运行的容器


```
docker ps -a


-a	显示所有的容器，包括未运行的

-f	根据条件过滤显示的内容

--format	指定返回值的模板文件

-l	显示最近创建的容器

-n	列出最近创建的n个容器

--no-trunc	不截断输出

-q	静默模式，只显示容器编号

-s	显示总的文件大小
```

### 关闭容器
```
docker stop [OPTIONS] CONTAINER [CONTAINER...]
```


### 启动容器
```
docker start [OPTIONS] CONTAINER [CONTAINER...]
```


### 重启容器
```
docker restart [OPTIONS] CONTAINER [CONTAINER...]
```


### 查看容器端口
```
docker port [OPTIONS] CONTAINER [PRIVATE_PORT[/PROTO]]
```


### 查看容器进程
```
docker top [OPTIONS] CONTAINER [ps OPTIONS]
```


### 移除容器
```
docker rm [OPTIONS] CONTAINER [CONTAINER...]


-f	通过SIGKILL信号强制删除一个运行中的容器

-l	移除容器间的网络连接，而非容器本身

-v	删除与容器关联的卷
```
删除容器时，容器必须是停止状态，否则会下错误




### 杀掉一个运行中的容器
```
docker kill [OPTIONS] CONTAINER [CONTAINER...]


-s	向容器发送一个信号
```


### 进入已启动的容器

```
docker exec [OPTIONS] CONTAINER COMMAND [ARG...]


-d	分离模式: 在后台运行

-i	即使没有附加也保持STDIN 打开

-t	分配一个伪终端
```

```
[root@password ~]# docker ps CONTAINER ID    IMAGE        COMMAND       CREATED       STATUS       PORTS        NAMES
43d44ac8cf31    e934aafc2206    "/bin/bash"     17 hours ago    Up 17 hours               sharp_heyrovsky
[root@password ~]# docker exec -ti 43d44ac8cf31 /bin/bash
[root@43d44ac8cf31 /]# 
```



### 暂停/恢复容器的所有进程

```
docker pause [OPTIONS] CONTAINER [CONTAINER...]	//暂停

docker unpause [OPTIONS] CONTAINER [CONTAINER...]	//恢复
```



### 提交容器(容器建立新的镜像)

```
docker commit [OPTIONS] CONTAINER [REPOSITORY[:TAG]]

-a	提交的镜像作者

-c	使用Dockerfile指令来创建镜像

-m	提交时的说明文字

-p	在commit时，将容器暂停


43d44ac8cf31	容器ID

test/lmap:v1	指定要创建的目标镜像名




```

```
[root@password ~]# docker commit -a "taolifeng" -m "this is lmnp tests" 43d44ac8cf31 test/lmap:v1
sha256:d28c02d47d3e5e17b8809f3c64acbd25de20197c36b04b092f39060eda0c0b86
[root@password ~]# docker images
REPOSITORY       TAG         IMAGE ID      CREATED       SIZE
test/lmap        v1         d28c02d47d3e    10 seconds ago   199 MB
test/centos       6.7         fe0c250bff03    19 hours ago    191 MB
docker.io/hello-world  latest       e38bc07ac18e    13 days ago     1.85 kB
docker.io/centos    latest       e934aafc2206    2 weeks ago     199 MB
```







### 列出docker镜像

```
docker images [OPTIONS] [REPOSITORY[:TAG]]


-a                列出本地所有的镜像（含中间映像层，默认情况下，过滤掉中间映像层）

--digests         显示镜像的摘要信息

-f                显示满足条件的镜像

--format          指定返回值的模板文件

--no-trunc        显示完整的镜像信息

-q                只显示镜像ID
```

```
[root@password ~]# docker images
REPOSITORY       TAG         IMAGE ID      CREATED       SIZE
docker.io/hello-world  latest       e38bc07ac18e    12 days ago     1.85 kB
docker.io/ubuntu    15.10        9b9cb95443b5    21 months ago    137 MB
```
```
REPOSITORY      表示镜像的仓库源

TAG             镜像的标签

IMAGE ID        镜像ID

CREATED         镜像创建时间

SIZE            镜像大小
```

### 通过REPOSITORY:TAG来使用docker镜像
```
docker run -t -i ubuntu:15.10 /bin/bash
```



### 获取docker镜像

```
docker pull [OPTIONS] NAME[:TAG|@DIGEST]


-a	拉取所有 tagged 镜像

--disable-content-trust	忽略镜像的校验,默认开启
```

```
docker pull ubuntu:13.10
```

当我们在本地主机上使用一个不存在的镜像时 Docker 就会自动下载这个镜像。如果我们想预先下载这个镜像，我们可以使用 docker pull 命令来下载它。


### 镜像网站


https://hub.docker.com/

### 镜像搜索
```
docker search [OPTIONS] TERM


--automated	只列出 automated build类型的镜像

--no-trunc	显示完整的镜像描述

-s	列出收藏数不小于指定值的镜像
```

```
docker search httpd
```

```
NAME            镜像仓库源的名称

DESCRIPTION     镜像的描述

OFFICIAL        是否docker官方发布
```

使用docker pull httpd 拖取镜像



### 更新镜像
```
[root@password ~]# docker run -t -i ubuntu:15.10 /bin/bash
root@0ec8f128b13c:/# apt-get update
Ign http://archive.ubuntu.com wily InRelease     
Ign http://archive.ubuntu.com wily-updates InRelease 
Ign http://archive.ubuntu.com wily-security InRelease
Ign http://archive.ubuntu.com wily Release.gpg
Ign http://archive.ubuntu.com wily-updates Release.gpg
Ign http://archive.ubuntu.com wily-security Release.gpg
Ign http://archive.ubuntu.com wily Release
Ign http://archive.ubuntu.com wily-updates Release
Ign http://archive.ubuntu.com wily-security Release
```




### 构建镜像


我们使用命令 docker build ， 从零开始来创建一个新的镜像。为此，我们需要创建一个 Dockerfile 文件

```
runoob@runoob:~$ cat Dockerfile 
FROM  centos:6.7
MAINTAINER   Fisher "fisher@sudops.com"

RUN   /bin/echo 'root:123456' |chpasswd
RUN   useradd runoob
RUN   /bin/echo 'runoob:123456' |chpasswd
RUN   /bin/echo -e "LANG=\"en_US.UTF-8\"" >/etc/default/local
EXPOSE 22
EXPOSE 80
CMD   /usr/sbin/sshd -D
		
docker build -t runoob/centos:6.7 ./
```

-t	指定要创建的目标镜像名

./	Dockerfile地址


### 设置镜像标签( 标记本地镜像，将其归入某一仓库。)

```
docker tag [OPTIONS] IMAGE[:TAG] [REGISTRYHOST/][USERNAME/]NAME[:TAG]
		
root@runoob:~# docker tag ubuntu:15.10 runoob/ubuntu:v3
root@runoob:~# docker images  runoob/ubuntu:v3
REPOSITORY     TAG         IMAGE ID      CREATED       SIZE
runoob/ubuntu    v3         4e3b13c8a266    3 months ago    136.3 MB
```


### 帮助

```
docker stats --help
```


### 日志

```
docker logs [OPTIONS] CONTAINER 


-f	跟踪日志输出

--since	显示某个开始时间的所有日志

-t	显示时间戳

--tail	仅列出最新N条容器日志
```