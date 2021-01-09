package Material

import (
	"brocaedu/Pkg/e"
	"brocaedu/Pkg/setting"
	"brocaedu/Services"
	"github.com/gin-gonic/gin"
	"github.com/unknwon/com"
)

func Index(c *gin.Context) {
	page := com.StrTo(c.Query("page")).MustInt()
	data := make(map[string]interface{})
	data["list"] = Services.GetMaterials(page, data)
	data["count"] = Services.GetTotalMaterials()
	data["size"] = setting.PageSize
	e.Success(c, "素材列表", data)
}

func AddMaterial(c *gin.Context) {
	code, msg := Services.AddMaterial(c)
	e.Success(c, msg, code)
}

// @Summer 单页详情Api
func GetMaterial(c *gin.Context) {
	c.Request.Body = e.GetBody(c)
	id := com.StrTo(c.PostForm("id")).MustInt()
	var data = make(map[string]interface{})
	data["detail"] = Services.GetMaterial(id)
	e.Success(c, "视频", data)
}
