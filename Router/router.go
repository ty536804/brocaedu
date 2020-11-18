package Router

import (
	backend "brocaedu/Backend"
	v1 "brocaedu/Backend/Controllers/Admin"
	v3 "brocaedu/Backend/Controllers/Article"
	v2 "brocaedu/Backend/Controllers/Banner"
	m "brocaedu/Backend/Controllers/Message"
	nav "brocaedu/Backend/Controllers/Navs"
	v4 "brocaedu/Backend/Controllers/Single"
	frontend "brocaedu/Frontend/Controllers"
	"brocaedu/Frontend/Controllers/Wap"
	"brocaedu/Middleware/jwt"
	"brocaedu/Pkg/e"
	"github.com/gin-gonic/gin"
	"html/template"
	"io"
	"net/http"
	"os"
)

func unescaped(x string) interface{} { return template.HTML(x) }

func subYear(x string) interface{} {
	return x[0:4]
}

func subDate(x string) interface{} {
	return x[5:10]
}

func InitRouter() *gin.Engine {
	// 禁用控制台颜色，将日志写入文件时不需要控制台颜色。
	gin.DisableConsoleColor()

	// 记录到文件。
	f, _ := os.Create("gin.log")
	gin.DefaultWriter = io.MultiWriter(f, os.Stdin)
	r := gin.New()
	r.SetFuncMap(template.FuncMap{
		"unescaped": unescaped,
		"subYear":   subYear,
		"subDate":   subDate,
	})
	r.Use(gin.Logger(), gin.Recovery())

	dir := e.GetDir()
	//加载后端js、样式文件
	r.StaticFS("static", http.Dir(dir+"/Resources/Public"))
	//加载后端文件
	r.LoadHTMLGlob("Resources/View/**/*")
	//首页
	r.GET("/", frontend.Index)
	r.GET("/index", frontend.FrontEnd)
	r.GET("/about", frontend.About)
	r.GET("/aboutData", frontend.AboutData)
	r.GET("/subject", frontend.Subject)
	r.GET("/subjectData", frontend.SubjectData)
	r.GET("/research", frontend.Research)
	r.GET("/researchData", frontend.ResearchData)
	r.GET("/learn", frontend.Learn)
	r.GET("/learnData", frontend.LearnData)
	r.GET("/news", frontend.News)
	r.GET("/newList", frontend.NewList)
	r.GET("/detail", frontend.NewDetail)
	r.GET("/join", frontend.Authorize)
	r.GET("/joinData", frontend.JoinData)
	r.GET("/campus", frontend.Campus) //全国校区
	r.GET("/down", frontend.Down)
	r.GET("/weChat", frontend.GetWeChat)
	r.GET("/search", Wap.Search)
	//移动端
	r.GET("/wap", Wap.Index)
	r.GET("/wapInfo", Wap.IndexInfo)
	r.GET("/sub", Wap.Subject)
	r.GET("/subInfo", Wap.SubjectInfo)
	r.GET("/le", Wap.Learn)
	r.GET("/leInfo", Wap.LearnInfo)
	r.GET("/authorize", Wap.Authorize)
	r.GET("/authorizeInfo", Wap.AuthorizeInfo)
	r.GET("/mAbout", Wap.About)
	r.GET("/mAboutInfo", Wap.AboutInfo)
	r.GET("/map", Wap.Map)
	r.GET("/cam", Wap.Campus)
	r.GET("/ai", Wap.AiLearn)
	r.GET("/list", Wap.News)
	r.GET("/de", Wap.NewDetail)

	r.POST("/AddMessage", m.AddMessage)
	r.POST("/getNavList", nav.GetNavList) //添加导航API
	//Backend
	apiv1 := r.Group("/api/v1")
	apiv1.Use(jwt.JWT())
	{
		//上传图片
		apiv1.POST("/upload", backend.UploadFile)

		apiv1.POST("/menus", v1.GetPowers)
		apiv1.POST("/powerShow", v1.PowerShow)
		apiv1.POST("/subjectData", frontend.SubjectData)
		apiv1.POST("/researchData", frontend.ResearchData)
		apiv1.POST("/learnData", frontend.LearnData)
		apiv1.POST("/newList", frontend.NewList)
		apiv1.POST("/login", v1.Login)
		//用户列表
		apiv1.POST("/userData", v1.UserData)       //用户列表API
		apiv1.POST("/AddUser", v1.AddUser)         //用户列表API
		apiv1.POST("/GetUser", v1.GetUser)         //查看当用户API
		apiv1.GET("/logout", v1.LogOut)            //查看当用户API
		apiv1.POST("/editUser", v1.UpdateUser)     //获取站点信息API
		apiv1.POST("/detailsUser", v1.DetailsUser) //查看当用户API

		// 短信
		apiv1.POST("/editSms", v1.AddSms) //编辑短信
		apiv1.POST("/getSms", v1.GetSms)  //查看短信

		//站点信息
		apiv1.POST("/addSite", v1.AddSite) //添加站点信息API
		apiv1.POST("/getSite", v1.GetSite) //获取站点信息API

		//banner
		apiv1.POST("/getBanners", v2.GetBanners)
		apiv1.POST("/AddBanner", v2.AddBanner)
		apiv1.POST("/getBanner", v2.GetBanner)
		apiv1.POST("/delBanner", v2.DelBanner)
		//message
		apiv1.POST("/messageData", m.ListData)
		//导航
		apiv1.POST("/getNavs", nav.GetNavs)     //获取多条导航API
		apiv1.POST("/getNav", nav.GetNav)       //获取一条导航API
		apiv1.POST("/addNav", nav.AddNav)       //添加导航API
		apiv1.POST("/menuList", nav.GetNavList) //添加导航API
		//文章
		apiv1.POST("/articleList", v3.ShowList)  //文章列表API
		apiv1.POST("/getArticle", v3.GetArticle) //文章详情API
		apiv1.POST("/addArticle", v3.AddArticle) //文章详情API
		//单页
		apiv1.POST("/singleList", v4.ListData) //文章列表API
		apiv1.POST("/getSingle", v4.GetSingle) //文章详情API
		apiv1.POST("/addSingle", v4.AddSingle) //添加单页详情API
	}

	return r
}
