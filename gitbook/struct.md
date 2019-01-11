### 目录结构
```
    +------------
    | README.md
    | SUMMARY.md
    | book.json
```

### README.md

关于这本书的简介, 书写语法是Markdown
```
# Gitbook 使用入门


> GitBook 是一个基于 Node.js 的命令行工具，可使用 Github/Git 和 Markdown 来制作精美的电子书。

本书将简单介绍如何安装、编写、生成、发布一本在线图书。
```

### SUMMARY.md
关于这本书的目录结构, 用于显示gitbook左侧导航栏
```
* [介绍](README.md)
* [基本安装](installation/README.md)
   * [Node.js安装](installation/nodejs-install.md)
   * [Gitbook安装](installation/gitbook-install.md)
   * [Gitbook命令行速览](installation/gitbook-cli.md)
* [图书项目结构](book/README.md)
   * [README.md 与 SUMMARY编写](book/file.md)
   * [目录初始化](book/prjinit.md)
* [图书输出](output/README.md)
   * [输出为静态网站](output/static.md)
   * [输出PDF](output/pdfandebook.md)
* [发布](publish/README.md)
   * [发布到Github Pages](publish/gitpages.md)
   * [发布到公司文档服务器](publish/companyserver.md)
* [结束](end/README.md)
```