<?php
require 'vendor/autoload.php';
function dump(...$vars) {
    ob_start();
    foreach ($vars as $var) {
        var_dump($var);
        echo PHP_EOL;
    }
    $data = ob_get_contents();
    ob_end_clean();
    echo $data;
}
function dd(...$var) {
    dump($var);
    exit(0);
}
class PhpGrpcClient {
    private $hostname = '127.0.0.1:8282';
    private $client = null;
    public function __construct()
    {
        $this->client = new App\Grpc\Sms\SmsServiceClient($this->hostname, [
            'credentials' => \Grpc\ChannelCredentials::createInsecure(),
            'timeout' => 5000,
        ]);
    }

    public function Send() {
        //新建map对象
        $dataMap = new \Google\Protobuf\Internal\MapField(
            \Google\Protobuf\Internal\GPBType::STRING,
            \Google\Protobuf\Internal\GPBType::MESSAGE,
            \Google\Protobuf\Any::class
        );
        //构建any值
        $val = new \Google\Protobuf\Int64Value();
        $val->setValue(1000);
        $valAny = new \Google\Protobuf\Any();
        $valAny->pack($val);
        //将any值装载到map里
        $dataMap['age'] = (new \Google\Protobuf\Any())
            ->setValue($valAny->getValue())
            ->setTypeUrl(str_replace('type.googleapis.com/', '', $valAny->getTypeUrl()));
        //构建请求的数据结构
        $request = new \App\Grpc\Sms\SendRequest();
        $request->setId(100);
        $request->setData($dataMap);
        //向服务端发送请求
        list($reply, $status) = $this->client->Send($request)->wait();
        if(isset($status) && $status->code == 0) {
            dump('调取成功');
        } else {
            dump('调取失败', $status);
        }
    }
}
$grpc = new PhpGrpcClient();
$grpc->Send();