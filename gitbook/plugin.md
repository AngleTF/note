### 插件地址
可以通过 Gitbook 的插件来扩展 Gitbook 的功能，现有的 Gitbook 插件能够实现数学公式，Google 统计，评论等等功能。

所有的插件都可以从 https://plugins.gitbook.com/ 获取。

### 插件安装
Gitbook 安装插件比较简单，需要在项目下添加 book.json 文件，然后在其中添加

```
{
    "plugins": ["plugins1", "plugins2"],
    "pluginsConfig": {
        "plugins1": {}
    }
}
```
通过下面的命令安装book.json中需要的插件
```
$ gitbook install
```