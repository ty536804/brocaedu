package Mofashuxue

import (
	"fmt"
	"strings"
	"time"
)

type Message struct {
	ID        int       `gorm:"primary_key" json:"id"`
	CreatedAt time.Time `json:"created_at" time_format:"2006-01-02 15:04:05"`
	UpdatedAt time.Time `json:"updated_at" time_format:"2006-01-02 15:04:05"`
	Content   string    `json:"content" gorm:"type:varchar(255);not null; default ''; comment:'机构名称' "`
	Mname     string    `json:"mname" gorm:"type:varchar(100);not null; default ''; comment:'留言姓名' "`
	Area      string    `json:"area" gorm:"type:varchar(100);not null; default ''; comment:'区域' "`
	Tel       string    `json:"tel" gorm:"type:varchar(20);not null; default ''; comment:'留言电话' "`
	Client    string    `json:"client" gorm:"type:varchar(190);not null; default ''; comment:'客户端' "`
	Ip        string    `json:"ip" gorm:"type:varchar(50);not null; default ''; comment:'ip地址' "`
	VisitUuid string    `json:"visit_uuid" gorm:"type:varchar(32);not null; default ''; comment:'用户ID' "`
	MsgType   int       `json:"msg_type" gorm:"type:not null; default '0'; comment:'1 魔法数学 2布卢卡斯' "`
	Com       string    `json:"com" gorm:"type:varchar(190);not null; default ''; comment:'留言来源页' "`
}

// @Desc 表单提交到队列
func SendMessageForMq(MName, area, tel, webType, ip, webCom, orgName string, msgType int) {
	if ipIndex := strings.LastIndex(ip, ":"); ipIndex != -1 {
		ip = ip[0:ipIndex]
	}
	word := new(Message)
	word.Mname = MName
	word.Area = area
	word.Tel = tel
	word.Client = webType
	word.Ip = ip
	word.Com = webCom
	word.VisitUuid = strings.Split(strings.Replace(ip, ".", "", -1), ":")[0]
	word.MsgType = msgType
	word.Content = orgName //机构名称
	result := magicDb.Create(&word)
	if result.Error != nil {
		fmt.Print("魔法添加留言失败", result)
	}
}

// @Desc 表单提交到队列
func SendMessage(MName, area, tel, webType, ip, webCom string) bool {
	if ipIndex := strings.LastIndex(ip, ":"); ipIndex != -1 {
		ip = ip[0:ipIndex]
	}
	word := new(Message)
	word.Mname = MName
	word.Area = area
	word.Tel = tel
	word.Client = webType
	word.Ip = ip
	word.Com = webCom
	word.VisitUuid = strings.Split(strings.Replace(ip, ".", "", -1), ":")[0]
	word.MsgType = 3
	result := magicDb.Create(&word)
	if result.Error != nil {
		fmt.Print("大英语添加留言失败", result)
		return false
	}
	return true
}

func GetTotalMessage(uid string, ftime, ltime string) (count int) {
	magicDb.Model(&Message{}).Where("ip = ? AND created_at >= ? AND created_at <= ?", uid, ftime, ltime).Count(&count)
	return
}
