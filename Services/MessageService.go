package Services

import (
	"brocaedu/Models/Elearn"
	"brocaedu/Models/Message"
	"brocaedu/Models/Mofashuxue"
	"brocaedu/Pkg/e"
	"fmt"
	"github.com/astaxie/beego/validation"
	"github.com/gin-gonic/gin"
	"github.com/unknwon/com"
	"regexp"
	"strings"
	"time"
)

// @Summer提交留言
func AddMessage(c *gin.Context) (code int, msg string) {
	var data = make(map[string]interface{})
	if err := c.Bind(&c.Request.Body); err != nil {
		fmt.Println(err)
		return e.ERROR, "操作失败"
	}
	tel := com.StrTo(c.PostForm("tel")).String()
	re := regexp.MustCompile(`^1\d{10}$`)
	if !re.MatchString(tel) || len(tel) < 11 {
		return e.SUCCESS, "请填写有效的手机号码"
	}

	ip := strings.Split(c.Request.RemoteAddr, ":")[0]
	initTime := time.Now().Format("2006-01-02")
	total := Message.GetTotalMessage(ip, initTime+" 00:00:00", initTime+" 23:59:59")
	if total >= 5 {
		return e.SUCCESS, "提交成功"
	}

	mname := TrimHtml(com.StrTo(c.PostForm("mname")).String())
	area := TrimHtml(com.StrTo(c.PostForm("area")).String())
	webCom := com.StrTo(c.PostForm("com")).String()
	webClient := com.StrTo(c.PostForm("client")).String()

	valid := validation.Validation{}
	valid.Required(mname, "mname").Message("姓名不能为空")
	valid.Required(area, "area").Message("地区不能为空")
	valid.Required(tel, "tel").Message("选择是否展示")

	if !valid.HasErrors() {
		data["mname"] = mname
		data["area"] = area
		data["tel"] = tel
		data["content"] = ""
		data["com"] = webCom
		data["client"] = webClient
		data["ip"] = ip
		data["channel"] = 1
		SendSmsToClient(area, mname, tel) //发送短信
		Mofashuxue.SendMessageForMq(mname, area, tel, webClient, ip, webCom, "", 3)
		Elearn.AddMessage(c, mname, area, tel) //elearn100
		if Message.AddMessage(data) {
			return e.SUCCESS, "提交成功"
		}
		return e.ERROR, "提交失败"
	}
	return ViewErr(valid)
}

func AddMsg(c *gin.Context) (code int, msg string) {
	if err := c.Bind(&c.Request.Body); err != nil {
		fmt.Println(err)
		return e.ERROR, "操作失败"
	}
	tel := com.StrTo(c.PostForm("tel")).String()
	re := regexp.MustCompile(`^1\d{10}$`)

	if !re.MatchString(tel) || len(tel) < 11 {
		return e.SUCCESS, "请填写有效的手机号码"
	}

	ip := strings.Split(c.Request.RemoteAddr, ":")[0]
	initTime := time.Now().Format("2006-01-02")
	total := Mofashuxue.GetTotalMessage(ip, initTime+" 00:00:00", initTime+" 23:59:59")
	if total >= 5 {
		return e.SUCCESS, "提交成功"
	}

	mname := TrimHtml(com.StrTo(c.PostForm("mname")).String())
	area := TrimHtml(com.StrTo(c.PostForm("area")).String())
	webCom := com.StrTo(c.PostForm("com")).String()
	webClient := com.StrTo(c.PostForm("client")).String()
	msgType := com.StrTo(c.PostForm("msg_type")).MustInt()
	orgName := com.StrTo(c.PostForm("content")).String()

	valid := validation.Validation{}
	valid.Required(mname, "mname").Message("姓名不能为空")
	valid.Required(area, "area").Message("地区不能为空")
	valid.Required(tel, "tel").Message("选择是否展示")

	if !valid.HasErrors() {
		SendSmsToClient(area, mname, tel) //发送短信
		Mofashuxue.SendMessageForMq(mname, area, tel, webClient, ip, webCom, orgName, msgType)
		if Mofashuxue.SendMessage(mname, area, tel, webClient, ip, webCom) {
			return e.SUCCESS, "提交成功"
		}
		return e.ERROR, "提交失败"
	}
	return ViewErr(valid)
}

func TrimHtml(src string) string {
	//将HTML标签全转换成小写
	re, _ := regexp.Compile("\\<[\\S\\s]+?\\>")
	src = re.ReplaceAllStringFunc(src, strings.ToLower)
	//去除STYLE
	re, _ = regexp.Compile("\\<style[\\S\\s]+?\\</style\\>")
	src = re.ReplaceAllString(src, "")
	//去除SCRIPT
	re, _ = regexp.Compile("\\<script[\\S\\s]+?\\</script\\>")
	src = re.ReplaceAllString(src, "")
	//去除所有尖括号内的HTML代码，并换成换行符
	re, _ = regexp.Compile("\\<[\\S\\s]+?\\>")
	src = re.ReplaceAllString(src, "\n")
	//去除连续的换行符
	re, _ = regexp.Compile("\\s{2,}")
	src = re.ReplaceAllString(src, "\n")
	return strings.TrimSpace(src)
}
