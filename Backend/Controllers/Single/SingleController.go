package Single

import (
	"brocaedu/Models/Single"
	"brocaedu/Pkg/e"
	"brocaedu/Pkg/setting"
	"brocaedu/Services"
	"github.com/gin-gonic/gin"
	"github.com/unknwon/com"
)

// @Summer 单页列表
func ListData(c *gin.Context) {
	page := com.StrTo(c.Query("page")).MustInt()
	data := make(map[string]interface{})
	data["list"] = Single.GetSingles(page, data)
	data["count"] = Single.GetSingleTotal()
	data["size"] = setting.PageSize
	e.Success(c, "单页列表", data)
}

// @Summer 添加单页
func AddSingle(c *gin.Context) {
	code, msg := Services.AddSingle(c)
	e.Success(c, msg, code)
}

// @Summer文章详情Api
func GetSingle(c *gin.Context) {
	c.Request.Body = e.GetBody(c)
	id := com.StrTo(c.PostForm("id")).MustInt()
	var data = make(map[string]interface{})
	data["list"] = Services.GetNavs(data)
	data["detail"] = Single.GetSingle(id)
	e.Success(c, "单页文章详情", data)
}
