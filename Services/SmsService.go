package Services

import (
	"brocaedu/Models/Admin"
	"brocaedu/Models/Site"
	"brocaedu/Pkg/e"
	"fmt"
	"github.com/astaxie/beego/validation"
	"github.com/gin-gonic/gin"
	"github.com/unknwon/com"
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

// @Summer 配置短信
func AddSms(c *gin.Context) (code int, msg string) {
	var data = make(map[string]interface{})
	if err := c.Bind(&c.Request.Body); err != nil {
		fmt.Println()
		return e.ERROR, "操作失败"
	}
	id := com.StrTo(c.PostForm("id")).MustInt()
	APIID := com.StrTo(c.PostForm("account")).String()
	APIKEY := com.StrTo(c.PostForm("password")).String()
	url := com.StrTo(c.PostForm("url")).String()

	valid := validation.Validation{}
	valid.Required(APIID, "account").Message("APIID不能为空")
	valid.Required(APIKEY, "password").Message("APIKEY不能为空")
	valid.Required(url, "url").Message("url不能为空")

	if !valid.HasErrors() {
		data["account"] = APIID
		data["password"] = APIKEY
		data["url"] = url

		isOk := false
		if id < 1 {
			isOk = Admin.AddSms(data)
		} else {
			isOk = Admin.EditSms(id, data)
		}
		if isOk {
			return e.SUCCESS, "操作成功"
		}
		return e.ERROR, "操作失败"
	}
	return ViewErr(valid)
}
