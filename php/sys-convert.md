### 十进制转换N进制
```php
$ary_map = [
    '0', '1', '2', '3', '4', '5', '6', '7', '8', '9', 'A', 'B', 'C', 'D', 'E', 'F'
];

function decimal_conver($number, $ary)
{
    global $ary_map;
    $res = (int)$number / $ary;
    if ($res > 0) return decimal_conver($res, $ary) . $ary_map[$number % $ary];
}

var_dump(decimal_conver(314156, 16));
```

### N进制转换十进制