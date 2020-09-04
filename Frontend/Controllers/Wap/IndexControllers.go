package Wap

import (
	"brocaedu/Models/Article"
	"brocaedu/Models/Banner"
	"brocaedu/Models/Single"
	"brocaedu/Pkg/e"
	"brocaedu/Services"
	"github.com/gin-gonic/gin"
)

var baseUrl = "http://www.brocaedu.com/"

// @Summer 首页
func Index(c *gin.Context) {
	var data = make(map[string]interface{})
	data["threeBanner"] = Banner.GetBannerByTag(1, 2, "3-6")
	data["sevenBanner"] = Banner.GetBannerByTag(1, 2, "7-15")
	data["banner"] = Banner.GetBannerData(1, 2) //轮播图
	Services.AddVisit(c, baseUrl+"wap")
	c.HTML(e.SUCCESS, "wap/index.html", gin.H{
		"title": "首页",
		"data":  data,
	})
}

// @Summer 首页API接口
func IndexInfo(c *gin.Context) {
	var data = make(map[string]interface{})
	data["ai"] = Banner.GetBannerByTag(1, 2, "解决方案")
	data["learn"] = Banner.GetBannerByTag(1, 2, "学习问题")
	data["moreBanner"] = Banner.GetBannerByTag(1, 2, "多维情景")
	data["learn"] = Banner.GetBannerByTag(1, 2, "学习问题")
	data["reasonBanner"] = Banner.GetBannerByTag(1, 2, "理由")
	data["brandBanner"] = Banner.GetBannerByTag(1, 2, "品牌介绍")
	data["sys"] = Banner.GetBannerByTag(1, 2, "BROCA智能学练系统")
	data["small"] = Banner.GetBannerByTag(1, 2, "小程序")
	var where = make(map[string]interface{})
	where["is_show"] = 1
	list := Article.GetArticles(1, where)
	if len(list) > 5 {
		list = list[0:4]
	}
	data["list"] = list
	e.Success(c, "首页", data)
}

// @Summer课程体系
func Subject(c *gin.Context) {
	Services.AddVisit(c, baseUrl+"sub")
	var data = make(map[string]interface{})
	data["banner"] = Banner.GetBannerData(3, 2) //轮播图
	c.HTML(e.SUCCESS, "wap/subject.html", gin.H{
		"title": "课程体系",
		"data":  data,
	})
}

// @Summer课程体系 API接口
func SubjectInfo(c *gin.Context) {
	var data = make(map[string]interface{})
	data["vsBanner"] = Banner.GetBannerByTag(3, 2, "vs")
	data["learnBanner"] = Banner.GetBannerByTag(3, 2, "学习场景")
	e.Success(c, "课程体系", data)
}

// @Summer 教学教研
func Learn(c *gin.Context) {
	var data = make(map[string]interface{})
	data["banner"] = Banner.GetBannerData(4, 2) //轮播图
	data["loop"] = Banner.GetBannerByTag(4, 1, "loop")
	data["leader"] = Banner.GetBannerByTag(4, 1, "leder")
	Services.AddVisit(c, baseUrl+"le")
	c.HTML(e.SUCCESS, "wap/learn.html", gin.H{
		"title": "教学教研",
		"data":  data,
	})
}

// @Summer 教学教研
func LearnInfo(c *gin.Context) {
	var data = make(map[string]interface{})
	data["selected"] = Banner.GetBannerByTag(4, 2, "选拔")
	data["checkAll"] = Single.GetConByTagAll(4, 2, "培训")
	data["appList"] = Banner.GetBannerByTag(4, 2, "app")
	e.Success(c, "教学教研", data)
}

// @Summer 加盟授权
func Authorize(c *gin.Context) {
	var data = make(map[string]interface{})
	data["banner"] = Banner.GetBannerData(7, 2) //轮播图
	Services.AddVisit(c, baseUrl+"authorize")
	c.HTML(e.SUCCESS, "wap/join.html", gin.H{
		"title": "加盟授权",
		"data":  data,
	})
}

// @Summer 加盟授权
func AuthorizeInfo(c *gin.Context) {
	var data = make(map[string]interface{})
	data["banner"] = Banner.GetBannerData(7, 2) //轮播图
	data["small"] = Single.GetConByTag(7, 2, "品牌起源")
	e.Success(c, "加盟授权", data)
}

// @Summer 关于我们
func About(c *gin.Context) {
	var data = make(map[string]interface{})
	data["banner"] = Banner.GetBannerData(2, 2) //轮播图
	Services.AddVisit(c, baseUrl+"mAbout")
	c.HTML(e.SUCCESS, "wap/about.html", gin.H{
		"title": "关于我们",
		"data":  data,
	})
}

// @Summer 关于我们
func AboutInfo(c *gin.Context) {
	var data = make(map[string]interface{})
	data["banner"] = Banner.GetBannerData(2, 2) //轮播图
	data["brand"] = Single.GetConByTag(2, 2, "品牌介绍")
	data["small"] = Single.GetConByTag(2, 2, "品牌起源")
	data["reasonBanner"] = Banner.GetBannerByTag(2, 2, "品牌荣誉")
	e.Success(c, "关于我们", data)
}

// @Summer 地图
func Map(c *gin.Context) {
	c.HTML(e.SUCCESS, "wap/map.html", gin.H{
		"title": "地图",
	})
}

// 全国中心
func Campus(c *gin.Context) {
	var data = make(map[string]interface{})
	data["banner"] = Banner.GetBannerData(6, 2)            //轮播图
	data["map"] = Banner.GetOneBanner(6, 2, "地图")          //
	data["hundred"] = Banner.GetOneBanner(6, 2, "200")     //
	data["thousand"] = Banner.GetOneBanner(6, 2, "100000") //
	data["offline"] = Banner.GetOneBanner(6, 2, "线下")      //
	c.HTML(e.SUCCESS, "wap/campus.html", gin.H{
		"title": "全国中心",
		"data":  data,
	})
}

// 全国中心
func AiLearn(c *gin.Context) {
	var data = make(map[string]interface{})
	data["banner"] = Banner.GetBannerData(5, 2) //轮播图
	c.HTML(e.SUCCESS, "wap/ai.html", gin.H{
		"title": "AI学习平台",
		"data":  data,
	})
}
