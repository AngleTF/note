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

```
gitbook build
```

### 调试

```
$ gitbook build ./ --log=debug --debug

```














