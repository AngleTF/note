### 安装windows

1. 为 PHP 的二进制可执行文件建立一个目录，例如 C:\bin

2. 将 ;C:bin 附加到 PATH 环境变量中（相关帮助）

3. 下载 https://phar.phpunit.de/phpunit-7.0.phar 并将文件保存到 C:\bin\phpunit.phar

4. 打开命令行（例如，按 WindowsR » 输入 cmd » ENTER)

5. 建立外包覆批处理脚本（最后得到 C:\bin\phpunit.cmd）：
```
C:\Users\username>  cd C:\bin
C:\bin>  echo @php "%~dp0phpunit.phar" %* > phpunit.cmd
C:\bin>  exit
新开一个命令行窗口，确认一下可以在任意路径下执行 PHPUnit：
```

### 安装查看
```
C:\Users\username> phpunit --version
PHPUnit x.y.z by Sebastian Bergmann and contributors.
```

### 官方文档
```
https://phpunit.readthedocs.io/zh_CN/latest/writing-tests-for-phpunit.html
```

