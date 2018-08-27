package main

import (
	"fmt"
	"github.com/enjoyass/AliDysmsapi"
)

func main() {
	var ssr *dysmsapi.SendSmsReply
	ssr,_ =dysmsapi.SendSms("13402651404","LTAIzYQhVdws2SX5","mju7qFeeZP4XhARdBvjbViYAX9Ka7o","{\"code\":\"123456\"}","SMS_136450036")
	fmt.Printf("%#v",ssr)
	fmt.Println("侧死")
}