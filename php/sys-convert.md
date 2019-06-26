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

**公式**
```
x = q * n + r

//反复除进制   x为10进制, q为x的商, n为进制, r为x的余数
314156 = 19634 * 16 + 12 (C)
19634  = 1227  * 16 + 2  (2)
1227   = 76    * 16 + 11 (B)
76     = 4     * 16 + 12 (C)
4      = 0     * 16 + 4  (4)

//4CB2C
```

### N进制转换十进制
```php
$ary_map = [
    '0', '1', '2', '3', '4', '5', '6', '7', '8', '9', 'A', 'B', 'C', 'D', 'E', 'F'
];
function conver_decimal($str, $ary)
{
    global $ary_map;
    $len = strlen((string)$str);
    if ($len > 0) return (int)(array_search($str[0], $ary_map) * pow($ary, $len - 1)) + conver_decimal(substr($str, 1), $ary);
}

var_dump(conver_decimal('4CB2C', 16));
```

**公式**
```
(4 * 16^4) + (12 * 16^3) + (11 * 16^2) + (2 * 16) + 12
```
