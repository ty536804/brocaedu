package Wap

import (
	"brocaedu/Models/Banner"
	"brocaedu/Models/Single"
	"brocaedu/Pkg/e"
	"brocaedu/Services"
	"github.com/gin-gonic/gin"
	"time"
)

// @Summer 首页
func Index(c *gin.Context) {
	Services.AddVisit(c)
	c.HTML(e.SUCCESS, "wap/index.html", gin.H{
		"title": "首页",
	})
}

// @Summer 首页API接口
func IndexInfo(c *gin.Context) {
	var data = make(map[string]interface{})
	data["banner"] = Banner.GetBannerData(1, 2) //轮播图
	data["learn"] = Banner.GetBannerByTag("学习问题")
	data["moreBanner"] = Banner.GetBannerByTag("多维情景")
	data["learn"] = Banner.GetBannerByTag("学习问题")
	data["threeBanner"] = Banner.GetBannerByTag("3-6")
	data["sevenBanner"] = Banner.GetBannerByTag("7-15")
	data["reasonBanner"] = Banner.GetBannerByTag("理由")
	data["brandBanner"] = Banner.GetBannerByTag("品牌介绍")
	data["sys"] = Single.GetConByTag("BROCA智能学练系统")
	data["small"] = Single.GetConByTag("小程序")
	e.Success(c, "首页", data)
}

// @Summer课程体系
func Subject(c *gin.Context) {
	ver := time.Now().Unix()
	Services.AddVisit(c)
	c.HTML(e.SUCCESS, "wap/subject.html", gin.H{
		"title": "课程体系",
		"time":  ver,
	})
}

// @Summer课程体系 API接口
func SubjectInfo(c *gin.Context) {
	var data = make(map[string]interface{})
	data["banner"] = Banner.GetBannerData(3, 2) //轮播图
	data["vsBanner"] = Banner.GetBannerByTag("vs")
	data["learnBanner"] = Banner.GetBannerByTag("学习场景")
	e.Success(c, "课程体系", data)
}

// @Summer AI学练系统
func Learn(c *gin.Context) {
	ver := time.Now().Unix()
	Services.AddVisit(c)
	c.HTML(e.SUCCESS, "wap/learn.html", gin.H{
		"title": "AI学联系统",
		"time":  ver,
	})
}

// @Summer AI学练系统
func LearnInfo(c *gin.Context) {
	var data = make(map[string]interface{})
	//data["banner"] = Banner.GetBannerData(1,0)  //轮播图
	//data["vsBanner"] = Banner.GetBannerByTag("vs")
	//data["learnBanner"] = Banner.GetBannerByTag("学习场景")
	e.Success(c, "AI学练系统", data)
}

// @Summer omo新模式
func Omo(c *gin.Context) {
	ver := time.Now().Unix()
	Services.AddVisit(c)
	c.HTML(e.SUCCESS, "wap/omo.html", gin.H{
		"title": "omo新模式",
		"time":  ver,
	})
}

// @Summer 加盟授权
func Authorize(c *gin.Context) {
	ver := time.Now().Unix()
	Services.AddVisit(c)
	c.HTML(e.SUCCESS, "wap/join.html", gin.H{
		"title": "加盟授权",
		"time":  ver,
	})
}

// @Summer 加盟授权
func AuthorizeInfo(c *gin.Context) {
	var data = make(map[string]interface{})
	data["banner"] = Banner.GetBannerData(7, 2) //轮播图
	data["small"] = Single.GetConByTag("品牌起源")
	e.Success(c, "加盟授权", data)
}

// @Summer 关于我们
func About(c *gin.Context) {
	ver := time.Now().Unix()
	Services.AddVisit(c)
	c.HTML(e.SUCCESS, "wap/about.html", gin.H{
		"title": "关于我们",
		"time":  ver,
	})
}

// @Summer 关于我们
func AboutInfo(c *gin.Context) {
	var data = make(map[string]interface{})
	data["banner"] = Banner.GetBannerData(2, 2) //轮播图
	data["brand"] = Single.GetConByTag("品牌介绍")
	data["small"] = Single.GetConByTag("品牌起源")
	data["reasonBanner"] = Banner.GetBannerByTag("品牌荣誉")
	e.Success(c, "关于我们", data)
}

// @Summer 地图
func Map(c *gin.Context) {
	c.HTML(e.SUCCESS, "wap/map.html", gin.H{
		"title": "地图",
	})
}
