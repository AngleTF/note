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

### win10和php7安装过程中碰到的问题
```
Warning: require_once(phar://go-pear.phar/PEAR/REST/13.php):
failed to open stream: phar error: "PEAR/REST/13.php" is not a file in phar "go-pear.phar"
in phar://E:/WMAP/php-7.2.16-Win32-VC15-x64/pear/go-pear.phar/PEAR/Config.php on line 2067
```

我在stackoverflow中找到了正确的解决方式
https://stackoverflow.com/questions/55295935/install-pear-on-macos-mojave-10-14-3/55302957#55302957?newreg=0fd6c0d00347439cba7f3a1de726ced7

download the phar
https://objects-us-east-1.dream.io/kbfiles/pear/go-pear.phar


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

