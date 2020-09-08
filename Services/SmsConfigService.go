package Services

import (
	"brocaedu/Models/Admin"
	"brocaedu/Pkg/e"
	"fmt"
	"github.com/astaxie/beego/validation"
	"github.com/gin-gonic/gin"
	"github.com/unknwon/com"
)

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
