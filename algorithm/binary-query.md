### 时间空间复杂度

![](/assets/binary-query.png)  
可以看到, 二分查询0毫秒就找到了具有顺序10w个数组中的88888这个值, 而使用\(O\)^n的for循环则使用更多的时间

### 普通遍历查询

```php
function search_in_array2($arr, $val){
    $len = count($arr);
    for ($i = 0; $i < $len; $i++){
        if($arr[$i] === $val){
            return $i;
        }
    }
    return false;
}
```

### 二分查询

```php
function search_in_array($arr, $val , $flag = 1){
    static $num = 1;
    static $sum = 0;

    $sum += ceil(count($arr) / pow(2,$num)) * $flag;

    $offset = count($arr) - $sum;

    if (!isset($arr[$offset])){
        return false;
    }

    echo "{$offset}\n";

    $num++;

    if ($arr[$offset] === $val) {
        return $offset;
    } elseif ($arr[$offset] < $val){
        return search_in_array($arr , $val , -1);
    }elseif ($arr[$offset] > $val){
        return search_in_array($arr , $val , 1);
    }

    return false;
}
```

### 生成10w条顺序数组结构

```php
$arr = [];
//创造有顺序的值
for ($i=0; $i < 100000; $i++){
    $arr[$i] = $i;
}
```

### 测试二分查询时间和空间复杂度

```php
$s1 = microtime(true);

echo search_in_array($arr, 88888) . "\n";

$s2 = microtime(true);

echo "二分查询耗时:" . round($s2 - $s1, 20);
```

### 原理

让你猜一个 0 ~ 100 的数字 , 你会怎么做 ?  
如果是从 0 猜到 100 那你的逻辑就是 上述的  普通遍历 了  
而二分算法的是怎么做的呢?  
二分算法是将总的数据对半分开 , 而后根据结果继续分半  
50\(小了\) -&gt; 75\(小了\) -&gt; 87\(大了\) -&gt; 81\(大了\) -&gt; 78\(大了\) -&gt; 77\(大了\) -&gt; 76\(正确\)  
那么无论怎么猜\(100范围\) , 次数都是在 7 次之内

