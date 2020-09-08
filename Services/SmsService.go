package Services

import (
	"brocaedu/Models/Admin"
	"brocaedu/Models/Site"
	"fmt"
	"io/ioutil"
	"net/http"
	"net/url"
	"strings"
)

var (
	_account  = "" //用户名是登录用户中心->国际短信->产品总览->APIID
	_password = "" //查看密码请登录用户中心->国际短信->产品总览->APIKEY
	_url      = ""
)

func init() {
	smsConfig := Admin.GetSmsConfig()
	_account = smsConfig.Account
	_password = smsConfig.Password
	_url = smsConfig.Url
}

// @Summer 发送验证码
func SendSms(mobile, msg string) {
	v := url.Values{}

	v.Set("account", _account)
	v.Set("password", _password)
	v.Set("mobile", mobile)
	v.Set("content", msg)

	body := ioutil.NopCloser(strings.NewReader(v.Encode())) //把form数据编下码
	client := &http.Client{}
	req, _ := http.NewRequest("POST", _url, body)

	req.Header.Set("Content-Type", "application/x-www-form-urlencoded; param=value")
	//fmt.Printf("看下发送的结构 %+v\n", req) //看下发送的结构

	resp, err := client.Do(req) //发送
	defer resp.Body.Close()     //一定要关闭resp.Body
	data, _ := ioutil.ReadAll(resp.Body)
	fmt.Println(string(data), err)
}

func SendSmsToClient(area, name, tel string) {
	site := Site.GetSite()
	var telList = strings.Split(strings.TrimSpace(site.AdminTel), ",")
	telList = append(telList, tel)
	if len(telList) > 0 {
		for k, tel := range telList {
			msg := ""
			if (k + 1) == len(telList) {
				msg = "我们已收到您的留言。我们的招商经理会在24小时内联系您，请您注意接听来自北京的电话，谢谢。"
			} else {
				msg = area + "的" + name + "留言了。联系" + tel + "留言来源布罗卡斯"
			}
			SendSms(tel, msg)
		}
	}
}
