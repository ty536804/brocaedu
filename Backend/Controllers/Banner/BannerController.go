package Banner

import (
	"brocaedu/Models/Banner"
	"brocaedu/Pkg/e"
	"brocaedu/Pkg/setting"
	"brocaedu/Services"
	"github.com/gin-gonic/gin"
	"github.com/unknwon/com"
)

// @Summer列表

// @Summer 获取所有图片
func GetBanners(c *gin.Context) {
	page := com.StrTo(c.Query("page")).MustInt()
	var data = make(map[string]interface{})
	data["count"] = Banner.GetBannerTotal()
	data["list"] = Banner.GetBanners(page)
	data["size"] = setting.PageSize
	e.Success(c, "获取banner列表", data)
}

func GetBanner(c *gin.Context) {
	id := com.StrTo(c.Query("id")).MustInt()
	var data = make(map[string]interface{})
	data["list"] = Services.GetNavs(data)
	data["detail"] = Banner.GetBanner(id)
	e.Success(c, "获取banner详情", data)
}

// @Summer banner保存
func AddBanner(c *gin.Context) {
	code, msg := Services.AddBanner(c)
	e.SendRes(c, code, msg, "")
}

// @Summer 删除banner
func DelBanner(c *gin.Context) {
	code, msg := Services.DelBanner(c)
	e.SendRes(c, code, msg, "")
}
