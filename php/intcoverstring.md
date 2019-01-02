```
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
var_dump($str)

?>
```



