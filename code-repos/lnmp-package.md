### 依赖安装
```
﻿yum -y install lrzsz unzip wget
```

### 下载一键部署安装包
```
wget http://tlf.freenovel.vip/assets/lnmp.zip
```

### 使用
```
unzip lnmp.zip
cd lnmp
vim run.sh
:set ff=unix
:wq
chmod +x run.sh
./run.sh
```