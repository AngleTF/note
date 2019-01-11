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

关于这本书的目录结构, 用于显示gitbook左侧导航栏, 书写语法是Markdown

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

### 配置文件
这里主要讲 book.json 的配置及参数，gitbook 使用该文件来配置整本书的基本信息，结构，使用的插件等等信息，这是一个非常重要的配置文件，这是一个json格式的文件。

| Variable    | Description |
| ----------- | ------------- |
| root        | 文件根目录， 但不包括book.json  |
| structure   | README, SUMMARY... 的路径  |
| title       | header中的title标签值  |
| description | header中的meta标签name为description的值  |
| author      | header中的meta标签name为author的值  |
| isbn        | 国际标准图书编号 |
| language    | 默认值是en  |
| direction   | 文本的方向 rtl 或 ltr  |
| gitbook     | 应该使用的gitbook版本号, 例如 ">= 3.0.0"  |
| plugins     | 需要加载的插件列表 (具体详见插件使用)  |
| pluginsConfig| 插件的配置 (具体详见插件使用) |




