package email_default

import (
	"github.com/jackylee92/rgo/core/rgconfig"
	"github.com/jackylee92/rgo/core/rgrequest"
	"github.com/jackylee92/rgo/util/rgemail"
)

/*
 * @Content : email_default
 * @Author  : LiJunDong
 * @Time    : 2022-11-06$
 */

const (
	defaultEmailHostConfig     = "default_email_host"
	defaultEmailPortConfig     = "default_email_port"
	defaultEmailUsernameConfig = "default_email_username"
	defaultEmailPasswordConfig = "default_email_password"
	defaultEmailFromConfig     = "default_email_from"
)

type Client struct {
	This     *rgrequest.Client
	ToEmails []string
	Content  string
}

// TODO <LiJunDong : 2022-11-06 14:34:51> --- 未成功

func (c *Client) SendCode() (err error) {
	emailClient := rgemail.EmailClient{
		This:     c.This,
		Title:    "验证码消息",
		To:       c.ToEmails,
		Content:  c.Content,
		From:     rgconfig.GetStr(defaultEmailFromConfig),
		Host:     rgconfig.GetStr(defaultEmailHostConfig),
		Port:     int(rgconfig.GetInt(defaultEmailPortConfig)),
		UserName: rgconfig.GetStr(defaultEmailUsernameConfig),
		Password: rgconfig.GetStr(defaultEmailPasswordConfig),
	}
	c.This.Log.Info("SendCode", emailClient)
	err = emailClient.Send()
	return err
}
