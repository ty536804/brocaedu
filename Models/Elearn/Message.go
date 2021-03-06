package Elearn

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"strings"
	"time"
)

// @Summer 留言
type JfsdMessageEnglish struct {
	ID         int    `gorm:"primary_key" json:"id"`
	Name       string `json:"name" gorm:"type:varchar(100); not null; default ''; comment:'姓名' "`
	Tel        string `json:"tel" gorm:"type:varchar(20); not null; default ''; comment:'电话' "`
	Content    string `json:"content" gorm:"type:varchar(9999); not null; default ''; comment:'地区' "`
	CreateTime int64  `json:"create_time" gorm:"comment:'创建时间'; default '0'" `
	Status     int    `json:"status" gorm:"not null; default '0' "`
	Ip         string `json:"ip" gorm:"type:varchar(100);not null; default ''; comment:'ip' "`
	VisitUuid  string `json:"visit_uuid" gorm:"type:varchar(32);not null; default ''; comment:'访问uuid' "`
}

// @Summer elearn100 添加留言
func AddMessage(c *gin.Context, mname, area, tel string) {
	uid := strings.Split(strings.Replace(c.Request.RemoteAddr, ".", "", -1), ":")[0]
	result := elearnDb.Create(&JfsdMessageEnglish{
		Name:       mname,
		Tel:        tel,
		Content:    area,
		Ip:         strings.Split(c.Request.RemoteAddr, ":")[0],
		CreateTime: time.Now().Unix(),
		VisitUuid:  uid,
	})
	if result.Error != nil {
		fmt.Printf("elelarn100 留言失败：%s", result.Error)
	} else {
		fmt.Print("elelarn100 留言成功")
	}
}
