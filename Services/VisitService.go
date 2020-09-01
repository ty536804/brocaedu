package Services

import (
	"brocaedu/Models/Elearn"
	"brocaedu/Models/Visit"
	"fmt"
	"github.com/gin-gonic/gin"
	"strings"
	"sync"
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
	var data = make(map[string]interface{})
	data["uuid"] = uid
	data["FirstUrl"] = FirstUrl
	data["Ip"] = strings.Split(c.Request.RemoteAddr, ":")[0]
	data["FromUrl"] = FromUrl
	data["visit_history"] = c.Request.Referer()

	if c.Request.RemoteAddr != "" {
		wg.Add(2)
		go AddBrocasVisit(data)
		go AddElearnVisit(data)
	} else {
		fmt.Println("没有拿到ip:网页地址：", c.Request.Referer())
	}
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
