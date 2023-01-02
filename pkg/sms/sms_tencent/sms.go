package sms_tencent

import (
	"fmt"
	"github.com/tencentcloud/tencentcloud-sdk-go/tencentcloud/common"
	"github.com/tencentcloud/tencentcloud-sdk-go/tencentcloud/common/profile"
	sms "github.com/tencentcloud/tencentcloud-sdk-go/tencentcloud/sms/v20210111"

	"github.com/tencentcloud/tencentcloud-sdk-go/tencentcloud/common/errors"
)

const (
	secretKeyConfig = ""
	secretIdConfig = ""
)

/*
 * @Content : sms_tencent
 * @Author  : LiJunDong
 * @Time    : 2022-11-04$
 */

// TODO <LiJunDong : 2022-11-06 11:59:13> --- 未接入

func Send() {
	// 实例化一个认证对象，入参需要传入腾讯云账户secretId，secretKey,此处还需注意密钥对的保密
	// 密钥可前往https://console.cloud.tencent.com/cam/capi网站进行获取
	credential := common.NewCredential(
		"SecretId",
		"SecretKey",
	)
	// 实例化一个client选项，可选的，没有特殊需求可以跳过
	cpf := profile.NewClientProfile()
	cpf.HttpProfile.Endpoint = "sms.tencentcloudapi.com"
	// 实例化要请求产品的client对象,clientProfile是可选的
	client, _ := sms.NewClient(credential, "", cpf)

	// 实例化一个请求对象,每个接口都会对应一个request对象
	request := sms.NewSendSmsRequest()

	// 返回的resp是一个SendSmsResponse的实例，与请求对象对应
	response, err := client.SendSms(request)
	if _, ok := err.(*errors.TencentCloudSDKError); ok {
		fmt.Printf("An API error has returned: %s", err)
		return
	}
	if err != nil {
		panic(err)
	}
	// 输出json格式的字符串回包
	fmt.Printf("%s", response.ToJsonString())
}