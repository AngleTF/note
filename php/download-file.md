### php实现如何将一个文件以下载的显示呈现

```php
<?php
$fileName = 'some.png';

header("content-type:application/octet-stream");
header('Content-Disposition: attachment; filename="test.png"');
header('Content-Length:'.filesize($fileName));

echo file_get_contents($fileName);

```

其实就是设置header头, 并且将数据输出.