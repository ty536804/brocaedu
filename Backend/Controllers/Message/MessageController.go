package Message

import (
	"brocaedu/Models/Message"
	"brocaedu/Pkg/e"
	"brocaedu/Pkg/setting"
	"brocaedu/Services"
	"github.com/gin-gonic/gin"
	"github.com/unknwon/com"
)

// @Summer 留言列表
func ListData(c *gin.Context) {
	page := com.StrTo(c.Query("page")).MustInt()
	data := make(map[string]interface{})
	data["list"] = Message.GetMessages(page)
	data["count"] = Message.GetMessageTotal()
	data["size"] = setting.PageSize
	e.Success(c, "留言列表", data)
}

// @Summer 添加留言
func AddMessage(c *gin.Context) {
	_, msg := Services.AddMessage(c)
	e.Success(c, msg, "")
}

func AddMsg(c *gin.Context) {
	c.Header("Access-Control-Allow-Origin", "*")
	c.Header("Access-Control-Allow-Headers", "Content-Type,AccessToken,X-CSRF-Token, Authorization, Token")
	c.Header("Access-Control-Allow-Methods", "POST, GET, OPTIONS")
	c.Header("Access-Control-Expose-Headers", "Content-Length, Access-Control-Allow-Origin, Access-Control-Allow-Headers, Content-Type")
	c.Header("Access-Control-Allow-Credentials", "true")
	_, msg := Services.AddMsg(c)
	e.Success(c, msg, "")
}
