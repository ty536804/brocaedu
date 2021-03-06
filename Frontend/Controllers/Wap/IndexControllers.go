package Wap

import (
	"brocaedu/Models/Article"
	"brocaedu/Models/Banner"
	"brocaedu/Models/Single"
	"brocaedu/Models/WeChat"
	"brocaedu/Pkg/e"
	"brocaedu/Pkg/setting"
	"brocaedu/Services"
	"github.com/gin-gonic/gin"
	"github.com/unknwon/com"
	"strconv"
	"strings"
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
	list := Article.GetArticles(1, setting.PageSize, where)
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
	data["banner"] = Banner.GetBannerData(6, 2)       //轮播图
	data["map"] = Banner.GetOneBanner(6, 2, "地图")     //
	data["offline"] = Banner.GetOneBanner(6, 2, "线下") //
	Services.AddVisit(c, baseUrl+"cam")
	c.HTML(e.SUCCESS, "wap/campus.html", gin.H{
		"title": "全国中心",
		"data":  data,
	})
}

// AI学习平台
func AiLearn(c *gin.Context) {
	var data = make(map[string]interface{})
	data["banner"] = Banner.GetBannerData(5, 2)            //轮播图
	data["ai"] = Banner.GetBannerByTag(5, 2, "imgList")    //
	data["school"] = Banner.GetOneBanner(5, 2, "school")   //
	data["aiAuto"] = Banner.GetBannerByTag(5, 2, "aiAuto") //
	Services.AddVisit(c, baseUrl+"ai")
	c.HTML(e.SUCCESS, "wap/ai.html", gin.H{
		"title": "AI学习平台",
		"data":  data,
	})
}

// @Summer 新闻资讯
func News(c *gin.Context) {
	var data = make(map[string]interface{})
	data["is_show"] = 1
	data["list"] = Article.GetArticles(1, setting.PageSize, data)
	data["banner"] = Banner.GetBannerData(8, 2) //轮播图
	data["count"] = e.GetPageNum(Article.GetArticleTotal())
	Services.AddVisit(c, baseUrl+"list")
	c.HTML(e.SUCCESS, "wap/new.html", gin.H{
		"title": "新闻资讯",
		"data":  data,
	})
}

// @Summer 新闻详情
func NewDetail(c *gin.Context) {
	id := com.StrTo(c.DefaultQuery("id", "0")).MustInt()
	_url := baseUrl + "de?id=" + string(id)
	Services.AddVisit(c, _url)
	detail := Article.GetArticle(id)
	c.HTML(e.SUCCESS, "wap/detail.html", gin.H{
		"title":  "新闻详情",
		"detail": detail,
	})
}

func Search(c *gin.Context) {
	c.HTML(e.SUCCESS, "wap/search.html", gin.H{})
}

// 视频列表
func VideoList(c *gin.Context) {
	page := com.StrTo(c.Query("page")).MustInt()
	data := make(map[string]interface{})
	data["is_show"] = 1
	data["list"] = Services.GetMaterials(page, data)
	data["count"] = e.GetPageNum(Services.GetTotalMaterials())
	data["size"] = setting.PageSize
	c.HTML(e.SUCCESS, "wap/videoList.html", gin.H{
		"title": "视频列表",
		"data":  data,
	})
}

//视频播放
func Video(c *gin.Context) {
	id := com.StrTo(c.Query("id")).MustInt()
	video := Services.GetMaterial(id)
	c.HTML(e.SUCCESS, "wap/video.html", gin.H{
		"title": "视频",
		"video": video,
	})
}

func CheckVideoPwd(c *gin.Context) {
	id := com.StrTo(c.PostForm("id")).MustInt()
	videoPwd := com.StrTo(c.PostForm("video_pwd")).String()
	if id < 1 {
		e.Error(c, "ID不能为空", "")
		return
	}
	video := Services.GetMaterial(id)

	if video.Code != videoPwd {
		e.Error(c, "视频播放码不正确", "")
		return
	}
	data := make(map[string]interface{})
	data["url"] = baseUrl + "videoDetail/id?=" + strconv.Itoa(id)
	uuid := strings.Split(strings.Replace(c.Request.RemoteAddr, ".", "", -1), ":")[0]
	uid, _ := strconv.Atoi(uuid)
	data["user_id"] = uid
	WeChat.AddLook(data)
	e.Success(c, "视频", video)
}
