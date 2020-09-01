package Services

import (
	"brocaedu/Models/Banner"
	"brocaedu/Pkg/e"
	"fmt"
	"github.com/astaxie/beego/validation"
	"github.com/gin-gonic/gin"
	"github.com/unknwon/com"
	"strings"
)

// @Summer 添加/编辑图片
func AddBanner(c *gin.Context) (code int, err string) {
	if err := c.Bind(&c.Request.Body); err != nil {
		fmt.Println(err)
		return e.ERROR, "操作失败"
	}

	id := com.StrTo(c.PostForm("id")).MustInt()
	bName := com.StrTo(c.PostForm("bname")).String()
	bPosition := com.StrTo(c.PostForm("bposition")).MustInt()
	imgUrl := com.StrTo(c.PostForm("imgurl")).String()
	targetLink := com.StrTo(c.PostForm("target_link")).String()
	isShow := com.StrTo(c.PostForm("is_show")).MustInt()
	info := com.StrTo(c.PostForm("info")).String()
	tag := com.StrTo(c.PostForm("tag")).String()
	clientType := com.StrTo(c.PostForm("type")).MustInt()

	if strings.HasPrefix(imgUrl, "/static/upload/") {
		imgUrl = strings.Replace(imgUrl, "/static/upload/", "", -1)
	}
	valid := validation.Validation{}
	valid.Required(bName, "bname").Message("名称不能为空")
	valid.Required(bPosition, "bposition").Message("展示位置必须选择")
	valid.Required(imgUrl, "imgurl").Message("上传图片")
	valid.Required(isShow, "is_show").Message("状态必须选择")
	valid.Required(tag, "tag").Message("标签不能为空")
	valid.Required(clientType, "type").Message("客户端不能为空")

	var data = make(map[string]interface{})
	if !valid.HasErrors() {
		data["target_link"] = targetLink
		data["info"] = info
		data["bname"] = bName
		data["bposition"] = bPosition
		data["imgurl"] = imgUrl
		data["is_show"] = isShow
		data["tag"] = tag
		data["type"] = clientType
		isOK := false
		if id < 1 {
			isOK = Banner.AddBanner(data)
		} else {
			isOK = Banner.EditBanner(id, data)
		}
		if isOK {
			return e.SUCCESS, "操作成功"
		}
		return e.ERROR, "操作失败"
	}
	return ViewErr(valid)
}

// @Summer 删除banner
func DelBanner(c *gin.Context) (code int, err string) {
	c.Request.Body = e.GetBody(c)
	id := com.StrTo(c.PostForm("id")).MustInt()
	isOK := Banner.DelBanner(id)
	if isOK {
		return e.SUCCESS, "操作成功"
	}
	return e.ERROR, "操作失败"
}

func GetBanner(tit string) (banner []Banner.Banner) {
	return Banner.GetBannerList(tit)
}
