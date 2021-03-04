package Mofashuxue

import (
	"fmt"
	"strings"
)

type Info struct {
	MName   string
	Area    string
	Tel     string
	Client  string
	Ip      string
	Uid     string
	Com     string
	MsgType int
}

// @Desc 表单提交到队列
func SendMessageForMq(MName, area, tel, webType, ip, webCom string) {
	if ipIndex := strings.LastIndex(ip, ":"); ipIndex != -1 {
		ip = ip[0:ipIndex]
	}
	word := new(Info)
	word.MName = MName
	word.Area = area
	word.Tel = tel
	word.Client = webType
	word.Ip = ip
	word.Com = webCom
	word.Uid = strings.Split(strings.Replace(ip, ".", "", -1), ":")[0]
	word.MsgType = 3
	result := magicDb.Create(&word)
	if result.Error != nil {
		fmt.Print("魔法添加留言失败", result)
	}
}
