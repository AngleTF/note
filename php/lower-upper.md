```
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



