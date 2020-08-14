package Services

import (
	"brocaedu/Models/Visit"
	"fmt"
	"github.com/gin-gonic/gin"
	"strings"
)

func AddBrocasVisit(c *gin.Context) {
	uid := strings.Split(strings.Replace(c.Request.RemoteAddr, ".", "", -1), ":")[0]
	visit := Visit.GetVisit(uid)
	if visit.ID <= 0 { //新增浏览记录
		Visit.AddVisit(c)
	} else {
		Visit.UpdateVisit(c)
	}
}

// @Summer 浏览记录
func AddVisit(c *gin.Context) {
	if c.Request.RemoteAddr != "" {
		//AddElearnVisit(c)
		AddBrocasVisit(c)
	} else {
		fmt.Println("没有拿到ip:网页地址：", c.Request.Referer())
	}
}