### 下载安装pptpd, ppp
```
yum -y install pptpd ppp
```

### 修改pptpd配置文件(/etc/pptpd.conf)
```
localip 192.168.0.1
remoteip 192.168.0.234-238
```
**localip**
> 如果是本机映射到路由器，该地址要改为路由器的入口地址

**remoteip**
> 外部拨号用户地址的地址池

### 修改pptp的dns配置(/etc/ppp/options.pptpd)
```
ms-dns 8.8.8.8
ms-dns 8.8.4.4
```

### 修改拨号用户账号密码(/etc/ppp/chap-secrets)

```
//格式 [account] pptpd [password] [allowip]; * 表示ip通配符

abc pptpd taolifeng *
```

### 修改系统文件(/etc/sysctl.conf)
```
//新增或者修改成

net.ipv4.ip_forward=1
```

**修改后退出保存**
```
sysctl -p
```

### iptables规则配置
```
iptables -t nat -A POSTROUTING -s 192.168.2.0/16 -o eth0 -j MASQUERADE  //是ifconfig得到的IP
```

### 启动pptpd
```
service pptpd start
```

### windows下使用pptp客户端
![](/assets/pptp.png)
![](/assets/pptp-1.png)
![](/assets/pptp-2.png)
![](/assets/pptp-3.png)

填入VPN的IP地址


![](/assets/pptp-4.png)

填入VPN的用户名和密码(例如上述例子使用的是 用户名:abc, 密码:taolifeng)

