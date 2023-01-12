package sms_aliyun

import (
	"errors"
	openapi "github.com/alibabacloud-go/darabonba-openapi/v2/client"
	dysmsapi20170525 "github.com/alibabacloud-go/dysmsapi-20170525/v3/client"
	util "github.com/alibabacloud-go/tea-utils/v2/service"
	"github.com/alibabacloud-go/tea/tea"
	"github.com/jackylee92/rgo/core/rgconfig"
	"github.com/jackylee92/rgo/core/rgrequest"
)

/*
 * @Content : sms_tencent
 * @Author  : LiJunDong
 * @Time    : 2022-11-04$
 */

// 【阿里云短信测试】您正在使用阿里云短信测试服务，体验验证码是：1234，如非本人操作，请忽略本短信！
type Client struct {
	This         *rgrequest.Client
	Title        string // 【阿里云短信测试】
	Phone        string // 手机号
	TemplateCode string // 在阿里云申请的内容模版编号
	Code         string // 模版中的变量，验证码
}

func (c *Client) Send() error {
	return c.do()
}

/**
 * 使用AK&SK初始化账号Client
 * @param accessKeyId
 * @param accessKeySecret
 * @return Client
 * @throws Exception
 */
func createClient(accessKeyId *string, accessKeySecret *string) (_result *dysmsapi20170525.Client, _err error) {
	config := &openapi.Config{
		AccessKeyId:     accessKeyId,
		AccessKeySecret: accessKeySecret,
	}
	config.Endpoint = tea.String("dysmsapi.aliyuncs.com")
	_result = &dysmsapi20170525.Client{}
	_result, _err = dysmsapi20170525.NewClient(config)
	return _result, _err
}

func (m *Client) do() (_err error) {
	accessKeyId := rgconfig.GetStr("access_key_id")
	accessKeySecret := rgconfig.GetStr("access_key_secret")
	if accessKeyId == "" || accessKeySecret == "" {
		return errors.New("AccessKey获取失败")
	}
	client, _err := createClient(tea.String(accessKeyId), tea.String(accessKeySecret))
	if _err != nil {
		return _err
	}

	sendSmsRequest := &dysmsapi20170525.SendSmsRequest{
		SignName:      tea.String(m.Title),
		TemplateCode:  tea.String(m.TemplateCode),
		PhoneNumbers:  tea.String(m.Phone),
		TemplateParam: tea.String("{\"code\":\"" + m.Code + "\"}"),
	}
	runtime := &util.RuntimeOptions{}
	tryErr := func() (_e error) {
		defer func() {
			if r := tea.Recover(recover()); r != nil {
				_e = r
			}
		}()
		// 复制代码运行请自行打印 API 的返回值
		_, _err = client.SendSmsWithOptions(sendSmsRequest, runtime)
		if _err != nil {
			return _err
		}

		return nil
	}()

	if tryErr != nil {
		var error = &tea.SDKError{}
		if _t, ok := tryErr.(*tea.SDKError); ok {
			error = _t
		} else {
			error.Message = tea.String(tryErr.Error())
		}
		// 如有需要，请打印 error
		_, _err = util.AssertAsString(error.Message)
		if _err != nil {
			return _err
		}
	}
	return _err
}
