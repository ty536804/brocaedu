package Services

import (
	"brocaedu/Models/Elearn"
	"brocaedu/Models/Mofashuxue"
	"brocaedu/Models/Visit"
	"brocaedu/Pkg/setting"
	"fmt"
	"github.com/gin-gonic/gin"
	"strings"
	"sync"
	"time"
)

var wg sync.WaitGroup

func AddBrocasVisit(data map[string]interface{}) {
	uid := data["uuid"].(string)
	visit := Visit.GetVisit(uid)

	defer wg.Done()

	if visit.ID <= 0 { //新增浏览记录
		Visit.AddVisit(data)
	} else {
		Visit.UpdateVisit(data["uuid"].(string), data["visit_history"].(string))
	}
}

// elearn100
func AddElearnVisit(data map[string]interface{}) {
	uid := data["uuid"].(string)
	visit := Elearn.GetVisit(uid)

	defer wg.Done()

	if visit.ID <= 0 {
		Elearn.AddVisit(data)
	} else {
		Elearn.UpdateVisit(data["uuid"].(string), data["visit_history"].(string))
	}
}

func AddMoFaShuue() {

}

// @Summer 浏览记录
func AddVisit(c *gin.Context, url string) {
	reqURI := c.Request.URL.RequestURI()
	FromUrl := c.Request.Host + reqURI //来源页
	uid := strings.Split(strings.Replace(c.Request.RemoteAddr, ".", "", -1), ":")[0]
	FirstUrl := ""
	if c.Request.Referer() == "" {
		FirstUrl = ReplaceSiteUrl(c.Request.Host, url, reqURI) //来源页
	} else {
		FirstUrl = c.Request.Referer()
	}
	fmt.Println(c.GetHeader("Referer"), "22222")
	var data = make(map[string]interface{})
	data["uuid"] = uid
	data["FirstUrl"] = FirstUrl
	data["Ip"] = strings.Split(c.Request.RemoteAddr, ":")[0]
	data["FromUrl"] = FromUrl
	data["visit_history"] = c.Request.Referer()

	if c.Request.RemoteAddr != "" {
		wg.Add(3)
		go AddBrocasVisit(data)
		go AddElearnVisit(data)
		go MoFaAddVisit(c, url)
	} else {
		fmt.Println("没有拿到ip:网页地址：", c.Request.Referer())
	}
	wg.Wait()
}

func ReplaceSiteUrl(url, first, reqURI string) string {
	if !strings.Contains("127.0.0.1", url) || url == "" || url == "/" {
		if reqURI == "" {
			return first
		} else {
			return first + reqURI
		}
	} else {
		return url
	}
}

func GetFirstUrl(Referer, host, url, reqURI string) string {
	if Referer == "" {
		return setting.MoReplaceSiteUrl(host, url, reqURI) //来源页
	}
	return Referer
}

func MoFaAddVisit(c *gin.Context, url string) {
	defer wg.Done()
	reqURI := c.Request.URL.RequestURI()
	uid := strings.Split(strings.Replace(c.Request.RemoteAddr, ".", "", -1), ":")[0]
	visitHistory := c.Request.Referer()

	var visit Mofashuxue.MoVisit
	visit.Uuid = uid
	visit.FirstUrl = GetFirstUrl(c.Request.Referer(), c.Request.Host, url, reqURI)
	visit.FromUrl = c.Request.Host + reqURI //来源页
	visit.CreateTime = time.Now().Format("2006-01-02 15:04:05")

	var history Mofashuxue.MoHistory
	history.Uuid = uid
	history.VisitHistory = visitHistory

	if c.Request.RemoteAddr != "" {
		visits := Mofashuxue.GetVisit(uid)
		visit.Ip = strings.Split(c.Request.RemoteAddr, ":")[0]
		if visits.ID <= 0 { //新增浏览记录
			visit.Ip = strings.Split(c.Request.RemoteAddr, ":")[0]
			Mofashuxue.AddVisit(visit, history)
		} else { //更新
			visitInfo := Mofashuxue.GetHistory(uid)
			if visitInfo.VisitHistory == "" {
				visitInfo.VisitHistory = visitHistory
			} else {
				visitInfo.VisitHistory = visitInfo.VisitHistory + "<br/>" + visitHistory
			}
			Mofashuxue.EditHistory(uid, visitInfo)
		}
	}
}
