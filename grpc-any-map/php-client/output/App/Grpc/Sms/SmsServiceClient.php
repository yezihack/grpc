<?php
// GENERATED CODE -- DO NOT EDIT!

namespace App\Grpc\Sms;

/**
 */
class SmsServiceClient extends \Grpc\BaseStub {

    /**
     * @param string $hostname hostname
     * @param array $opts channel options
     * @param \Grpc\Channel $channel (optional) re-use channel object
     */
    public function __construct($hostname, $opts, $channel = null) {
        parent::__construct($hostname, $opts, $channel);
    }

    /**
     * @param \App\Grpc\Sms\SendRequest $argument input argument
     * @param array $metadata metadata
     * @param array $options call options
     */
    public function Send(\App\Grpc\Sms\SendRequest $argument,
      $metadata = [], $options = []) {
        return $this->_simpleRequest('/sms.SmsService/Send',
        $argument,
        ['\App\Grpc\Sms\SendReply', 'decode'],
        $metadata, $options);
    }

    /**
     * @param \App\Grpc\Sms\HelloWorldRequest $argument input argument
     * @param array $metadata metadata
     * @param array $options call options
     */
    public function HelloWorld(\App\Grpc\Sms\HelloWorldRequest $argument,
      $metadata = [], $options = []) {
        return $this->_simpleRequest('/sms.SmsService/HelloWorld',
        $argument,
        ['\App\Grpc\Sms\HelloWorldReply', 'decode'],
        $metadata, $options);
    }
}