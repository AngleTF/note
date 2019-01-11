### 大小写转化php code
我们都知道, 计算机并不认识 0 和 1 以外的东西, 我们将特定编码的值存入显存, 显卡在进行渲染工作, 比如 1 这个值, 需要将 1 转化成显卡认识的 ASCII 49 (10进制), 然后存入显存内再由显卡去工作
```php
<?php
//1100001 1000001
function strLower($s){
    if (!is_string($s)) return "";
    foreach (str_split($s) as $k => &$v){
        $as = ord($v);
        $s[$k] = $as >= 65 && $as <= 90 ? chr($as | 0x20) : $v;
    }
    return $s;
}

function strUpper($s){
    if (!is_string($s)) return "";
    foreach (str_split($s) as $k => &$v){
        $as = ord($v);
        $s[$k] = $as >= 97 && $as <= 122 ? chr($as &  0b1011111) : $v;
    }
    return $s;
}

var_dump(
    strUpper('HELLO worD qwertyuiopasdfghjklzxcvbnm, 你好 世界')
);
```
提示: 字母A与a的ASCII相差32(十进制)

### 汇编show_str示例
向界面输出 绿色的 "welcome to masm!"
![](/assets/2019-01-11_113347.png)
![](/assets/2019-01-11_113624.png)

