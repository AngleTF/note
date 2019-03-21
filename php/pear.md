### 简介
PEAR是PHP扩展与应用库(the PHP Extension and Application Repository)的缩写。它是一个PHP扩展及应用的一个代码仓库，简单地说，PEAR之于PHP就像是CPAN(Comprehensive Perl Archive Network)之于Perl。
PEAR的基本目标是发展成为PHP扩展和库代码的知识库，而这个项目最有雄心的目标则是试图定义一种标准，这种标准将帮助开发者编写可移植、可重用的代码。

### 下载pear安装脚本
```
wget https://pear.php.net/go-pear.phar
```
也可以手动下载保存到本地目录, 下载后将脚本移动到php的目录(可以不移动)

### 安装pear
```
php go-pear.phar
```
安装过程中需要指定system和local模式和php.exe的目录路径, 其他的一般都有默认安装地址

### 验证安装
```
D:\WMAP\php-5.6.32>pear version
PEAR Version: 1.10.9
PHP Version: 5.6.32
Zend Engine Version: 2.6.0
Running on: Windows NT TF 6.1 build 7601 (Windows 7 Ultimate Edition Service Pack 1)
```

### ini的改动
```ini
;***** Added by go-pear
include_path=".;D:\WMAP\php-5.6.32\pear"
;*****
```


### 另一种安装方式

直接`双击php目录下的go-pear.bat`

