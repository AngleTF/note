## Git三种状态

---

### committed   \(已提交\)
![](/assets/gitstatus-3.png)

> 当我有新文件添加的时候通过git status发现文件未追踪, 这个时候需要使用git add 命令进行对文件的追踪, add 后的文件就是 已暂存状态staged


<br>
### modiffed \(已修改\)
![](/assets/gitstatus-4.png)

> 而后我们使用git commit 提交到自己的仓库, -m 是提交的注释 这是必须的, 提交后你的状态就变成 已提交状态committed


<br>

### staged \(已暂存\)
![](/assets/gitstatus-1.png)

> 当我对某个文件的数据进行修改的时候,  git通过对文件进行hash对比, 如果该文件与上次存储的文件hash值不相等, 那么git会判断改文件被修改会将状态改成modified, 我们需要对modified的文件重新git add



