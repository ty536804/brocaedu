package Controllers

import (
	"brocaedu/Models/Article"
	"brocaedu/Models/Banner"
	"brocaedu/Models/Single"
	"brocaedu/Pkg/e"
	"brocaedu/Pkg/setting"
	"brocaedu/Services"
	"github.com/gin-gonic/gin"
	"github.com/unknwon/com"
)

var baseUrl = "http://www.brocaedu.com/"

func Index(c *gin.Context) {
	Services.AddVisit(c, baseUrl+"index")
	c.HTML(e.SUCCESS, "index/index.html", gin.H{
		"title": "首页",
	})
}

// @Summer 首页
func FrontEnd(c *gin.Context) {
	var where = make(map[string]interface{})
	where["is_show"] = 1
	list := Article.GetArticles(1, setting.PageSize, where)
	if len(list) > 5 {
		list = list[0:4]
	}
	var data = make(map[string]interface{})
	data["list"] = list                         //新闻列表
	data["banner"] = Banner.GetBannerData(1, 1) //轮播图
	data["learn"] = Banner.GetOneBanner(1, 1, "你正在面临的学习问题")
	data["plan"] = Banner.GetBannerByTag(1, 1, "解决方案")
	data["ai"] = Banner.GetOneBanner(1, 1, "ai")
	data["small"] = Single.GetConByTag(1, 1, "3-6")
	data["seven"] = Single.GetConByTag(1, 1, "7-12")
	data["brand"] = Single.GetConByTag(1, 1, "品牌介绍")
	data["reason"] = Banner.GetBannerByTag(1, 1, "理由")
	data["online"] = Banner.GetBannerByTag(1, 1, "线上系统")
	e.Success(c, "首页", data)
}

func About(c *gin.Context) {
	Services.AddVisit(c, baseUrl+"about")
	c.HTML(e.SUCCESS, "index/about.html", gin.H{
		"title": "关于我们",
	})
}

func AboutData(c *gin.Context) {
	var data = make(map[string]interface{})
	data["banner"] = Banner.GetBannerData(2, 1)
	data["brand"] = Single.GetConByTag(2, 1, "19世纪")
	e.Success(c, "关于我们", data)
}

// @Summer课程体系
func Subject(c *gin.Context) {
	Services.AddVisit(c, baseUrl+"subject")
	c.HTML(e.SUCCESS, "index/subject.html", gin.H{
		"title": "课程体系",
	})
}

// @Summer课程体系
func SubjectData(c *gin.Context) {
	var data = make(map[string]interface{})
	data["banner"] = Banner.GetBannerData(3, 1)
	e.Success(c, "课程体系", data)
}

// @Summer教研教学
func Research(c *gin.Context) {
	Services.AddVisit(c, baseUrl+"research")
	c.HTML(e.SUCCESS, "index/research.html", gin.H{
		"title": "教研教学",
		"loop":  Banner.GetBannerByTag(4, 1, "loop"),
	})
}

// @Summer教研教学
func ResearchData(c *gin.Context) {
	var data = make(map[string]interface{})
	data["banner"] = Banner.GetBannerData(4, 1)
	data["reason"] = Banner.GetBannerByTag(4, 1, "leder")
	data["teacher"] = Banner.GetOneBanner(4, 1, "teacher")
	data["app"] = Banner.GetBannerByTag(4, 1, "辅学APP")
	e.Success(c, "教研教学", data)
}

// @Summer AI学习平台
func Learn(c *gin.Context) {
	Services.AddVisit(c, baseUrl+"learn")
	ai := Banner.GetBannerData(5, 1) //轮播图
	c.HTML(e.SUCCESS, "index/ai.html", gin.H{
		"title": "ai学习平台",
		"ai":    ai,
	})
}

// @Summer AI学习平台
func LearnData(c *gin.Context) {
	var data = make(map[string]interface{})
	data["banner"] = Banner.GetBannerData(5, 1)
	e.Success(c, "教研教学", data)
}

// @Summer OMO模式
func Omo(c *gin.Context) {
	Services.AddVisit(c, baseUrl+"omo")
	c.HTML(e.SUCCESS, "index/omo.html", gin.H{
		"title": "OMO模式",
	})
}

// @Summer全国校区
func Campus(c *gin.Context) {
	Services.AddVisit(c, baseUrl+"campus")
	c.HTML(e.SUCCESS, "index/campus.html", gin.H{
		"title": "全国校区",
	})
}

// @Summer 新闻动态
func News(c *gin.Context) {
	Services.AddVisit(c, baseUrl+"news")
	c.HTML(e.SUCCESS, "index/new.html", gin.H{
		"title": "新闻动态",
	})
}

// @Summer 新闻动态列表
func NewList(c *gin.Context) {
	page := com.StrTo(c.Query("page")).MustInt()
	var data = make(map[string]interface{})
	data["is_show"] = 1
	data["list"] = Article.GetArticles(page, setting.PageSize, data)
	data["banner"] = Banner.GetBannerData(8, 1) //轮播图
	data["count"] = e.GetPageNum(Article.GetArticleTotal())
	e.Success(c, "新闻列表", data)
}

// @Summer 新闻详情
func NewDetail(c *gin.Context) {
	id := com.StrTo(c.DefaultQuery("id", "0")).MustInt()
	_url := baseUrl + "detail?id=" + string(id)
	Services.AddVisit(c, _url)
	c.HTML(e.SUCCESS, "index/detail.html", gin.H{
		"title":  "新闻详情",
		"detail": id,
	})
}

// @Summer 新闻详情
func NewDetailData(c *gin.Context) {
	c.Request.Body = e.GetBody(c)
	id := com.StrTo(c.DefaultQuery("id", "0")).MustInt()
	var data = make(map[string]interface{})
	data["detail"] = Article.GetArticle(id)
	e.Success(c, "文章详情", data)
}

// @Summer 加盟授权
func Authorize(c *gin.Context) {
	Services.AddVisit(c, baseUrl+"join")
	c.HTML(e.SUCCESS, "index/join.html", gin.H{
		"title": "加盟授权",
	})
}

// @Summer 加盟授权数据接口
func JoinData(c *gin.Context) {
	c.Request.Body = e.GetBody(c)
	var data = make(map[string]interface{})
	e.Success(c, "success", data)
}

// @Summer 加盟授权
func Down(c *gin.Context) {
	Services.AddVisit(c, baseUrl+"down")
	c.HTML(e.SUCCESS, "index/down.html", gin.H{
		"title": "APP下载",
	})
}
