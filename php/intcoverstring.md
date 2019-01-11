### 强制转换原理
各取某一位值, 然后进行换算
```php
<?php

$v = 12345670;
$str = '';
while(true){
    //取每位值的各位数, 然后将其转化为SACII
    $lastNum = $v % 10;

    $str .= chr($lastNum + 0x30);

    if ($v >= 10){
        $v = $v / 10;
    }else{
        //最后反转一下
        $str = strrev($str);
        break;
    }
}

//string '12345670'
var_dump($str);

//---------------------string cover int------------------------------
$v1 = '12355';
$v1 = strrev($v1);
$v2 = 0;

for ($i = 0; $i < strlen($v1); $i++){
    $v2 += $v1[$i] * pow(10, $i);
}

//int(12355)
var_dump($v2)

?>
```



