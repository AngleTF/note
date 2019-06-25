### 云通讯手机短信通知类

```php
<?php
class YunTongXun
{
    private $accountSId = '';
    private $authToken = '';
    private $appId = '';
    private $templateId = 1;

    /**
     * 替换模板的数据
     * @var array
     */
    private $templateReplaceData = [];

    /**
     * 请求地址
     * @var string
     */
    private $requestUrl = "https://app.cloopen.com:8883/2013-12-26/Accounts/%s/SMS/TemplateSMS?sig=%s";

    /**
     * 手机号码
     * @var string
     */
    private $phone;

    /**
     * 请求时间戳 格式"yyyyMMddHHmmss"
     * @var string
     */
    private $nowTime;

    /**
     * 错误信息, 如果成功则为null
     * @var string
     */
    private $error = null;

    /**
     * 错误代码, 如果成功则为0
     * @var int
     */
    private $errorCode = 0;

    /**
     * YunTongXun constructor.
     * @param array $config ['account_sid' => '', 'app_id' => '', 'auth_token' => '', 'template_id' => '', 'phone' => '']
     */
    public function __construct($config)
    {
        $this->nowTime = date('YmdHis');
        $this->setPhone($config['phone']);
        $this->setAccountSId($config['account_sid']);
        $this->setAppId($config['app_id']);
        $this->setAuthToken($config['auth_token']);
        $this->setTemplateId($config['template_id']);
        $this->initUri();
    }

    /**
     * 初始化资源定位符
     */
    public function initUri()
    {
        $sig = strtoupper(md5($this->accountSId . $this->authToken . $this->nowTime));
        $this->requestUrl = sprintf($this->requestUrl, $this->accountSId, $sig);
    }

    /**
     * 发送短信
     * @param array $data
     * @return bool
     * @throws \Exception
     */
    public function send($data = null)
    {
        if ($data !== null) {
            $this->setTemplateData($data);
        }

        $authorization = base64_encode($this->accountSId . ':' . $this->nowTime);

        $http_header = [
            'Accept' => 'application/json',
            'Content-Type' => 'application/json;charset=utf-8',
            'Authorization' => $authorization
        ];

        $http_body = [
            'to' => $this->phone,
            'appId' => $this->appId,
            'templateId' => $this->templateId,
            'datas' => $this->templateReplaceData
        ];

        $str_json = json_encode($http_body);

        //追加Content-Length
        $http_header['Content-Length'] = strlen($str_json);

        $this->mergeHttpHeader($http_header);

        $ch = curl_init($this->requestUrl);
        curl_setopt($ch, CURLOPT_HTTPHEADER, $http_header);
        curl_setopt($ch, CURLOPT_POSTFIELDS, $str_json);
        curl_setopt($ch, CURLOPT_RETURNTRANSFER, true);
        $result = curl_exec($ch);

        $err_code = curl_errno($ch);
        $http_code = curl_getinfo($ch, CURLINFO_HTTP_CODE);

        if ($err_code != 0 || $http_code < 200 || $http_code >= 300 || $result === false) {
            throw new \Exception("http request error, 
            url:{$this->requestUrl}, httpCode: {$http_code}, curlCode: {$err_code}, curlMsg:" . curl_error($ch));
        }

        curl_close($ch);

        return $this->parse($result);
    }

    /**
     * 根据规则解析返回值, 并且返回是否请求成功
     * @param $result
     * @return bool
     * @throws \Exception
     */
    public function parse($result)
    {
    
        $result_arr = json_decode($result, true);

        if (!isset($result_arr['statusCode'])) {
            throw new \Exception("解析响应失败, result: {$result}");
        }

        if ($result_arr['statusCode'] != "000000") {
            $this->errorCode = $result_arr['statusCode'];
            $this->error = isset($result_arr['statusMsg']) ? $result_arr['statusMsg'] : "错误";
            return false;
        }

        return true;
    }

    public function mergeHttpHeader(&$header)
    {
        $new_header = [];
        foreach ($header as $k => $v) {
            $new_header[] = "{$k}: $v";
        }
        $header = $new_header;
    }

    public function setPhone($phone)
    {
        $this->phone = $phone;
    }

    public function getPhone()
    {
        return $this->phone;
    }

    public function setTemplateData($data)
    {
        $this->templateReplaceData = $data;
    }

    public function getTemplateData($data)
    {
        return $this->templateReplaceData;
    }

    public function getErrorCode()
    {
        return $this->errorCode;
    }

    public function getError()
    {
        return $this->error;
    }

    public function setAccountSId($account_id)
    {
        $this->accountSId = $account_id;
    }

    public function getAccountSId()
    {
        return $this->accountSId;
    }

    public function setAuthToken($authToken)
    {
        $this->authToken = $authToken;
    }

    public function getAuthToken()
    {
        return $this->authToken;
    }

    public function setAppId($appId)
    {
        $this->appId = $appId;
    }

    public function getAppId()
    {
        return $this->appId;
    }

    public function setTemplateId($templateId)
    {
        $this->templateId = $templateId;
    }

    public function getTemplateId()
    {
        return $this->templateId;
    }
}
```

### 使用
```php
$yun = new YunTongXun([
        'phone' => '123456789',
        'account_sid' => '111',
        'app_id' => '222',
        'auth_token' => 'QWERTYUIOP',
        'template_id' => 1,
    ]);
    
    
try {

    //发送短信通知
    if ($yun->send(['模板变量1', '模板变量2'])) {
        echo 'success';
    } else {
        echo $yun->getError();
    }

} catch (\Exception $e) {
    echo "Exception:{$e->getMessage()}\n";
}
```
