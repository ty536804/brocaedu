package Services

import (
	"brocaedu/Models/Admin"
	"github.com/gin-gonic/gin"
	"github.com/unknwon/com"
)

// @Summer 获取权限列表
// @Param total int 总计
// @Param list  []  权限列表
func GetAdmins(c *gin.Context) (res map[string]interface{}) {
	var data = make(map[string]interface{})
	page := com.StrTo(c.PostForm("page")).MustInt()
	data["list"] = Admin.GetAdminUserList(page, data)
	data["total"] = Admin.GetTotalAdmin()
	return data
}

// @Summer 通过制定ID获取管理员信息
func GetAdmin(c *gin.Context) (data map[string]interface{}) {
	id := com.StrTo(c.PostForm("id")).MustInt()
	res := make(map[string]interface{})
	res["user"] = Admin.GetAdminUser(id)
	return res
}
