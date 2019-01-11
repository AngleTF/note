### GitBook初始化
```
$ gitbook init
info: create SUMMARY.md
info: initialization is finished
```
在使用 gitbook init 之后本地会生成两个文件 README.md 和 SUMMARY.md ，这两个文件都是必须的，一个为介绍，一个为目录结构。


### 编辑左侧导航栏
GitBook使用SUMMARY.md文件组织整个内容的目录，比如可以新建 Faq.md 文件，来记录常见问题，并在 SUMMARY.md 文件中添加目录。
```
# Summary

* [简介](README.md)
* [常见问题](Faq.md)
```


### 开启本地服务器
gitbook serve 命令实际会先调用 gitbook build 编译书籍，完成后打开 web 服务器，默认监听本地 4000 端口，在浏览器打开 http://localhost:4000 即可浏览电子书。
```
$ gitbook serve
Live reload server started on port: 35729
Press CTRL+C to quit ...

info: 11 plugins are installed
info: 8 explicitly listed
info: loading plugin "expandable-chapters-small"... OK
info: loading plugin "livereload"... OK
info: loading plugin "highlight"... OK
info: loading plugin "search"... OK
info: loading plugin "lunr"... OK
info: loading plugin "sharing"... OK
info: loading plugin "fontsettings"... OK
info: loading plugin "theme-default"... OK
info: found 11 pages
info: found 21 asset files
info: >> generation finished with success in 1.4s !

Starting server ...
Serving book on http://localhost:4000
```

### 构建静态网站
该命令会在当前目录生成 _book 文件夹, 这个文件是静态网页的电子书
```
gitbook build
```

### 调试

```
$ gitbook build ./ --log=debug --debug

```


### 帮助

```
$ gitbook -h

  Usage: gitbook [options] [command]


  Options:

    -v, --gitbook [version]  specify GitBook version to use
    -d, --debug              enable verbose error
    -V, --version            Display running versions of gitbook and gitbook-cli
    -h, --help               output usage information


  Commands:

    ls                        List versions installed locally
    current                   Display currently activated version
    ls-remote                 List remote versions available for install
    fetch [version]           Download and install a <version>
    alias [folder] [version]  Set an alias named <version> pointing to <folder>
    uninstall [version]       Uninstall a version
    update [tag]              Update to the latest version of GitBook
    help                      List commands for GitBook
    *                         run a command with a specific gitbook version



```













