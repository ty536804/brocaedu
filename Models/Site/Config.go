package Site

import (
	db "brocaedu/Database"
	"fmt"
	"time"
)

type WeChat struct {
	ID        int       `gorm:"primary_key" json:"id"`
	APPID     string    `json:"appid"`
	APPSECRET string    `json:"appsecret"`
	GRANTTYPE string    `json:"granttype"`
	CreatedAt time.Time `json:"created_at" time_format:"2006-01-02 15:04:05"`
	UpdatedAt time.Time `json:"updated_at" time_format:"2006-01-02 15:04:05"`
}

// @Summer 添加/编辑 微信配置基础信息
func AddWeChatConfig(id int, data map[string]interface{}) bool {
	if id > 0 {
		res := db.Db.Create(&WeChat{
			APPID:     data["appid"].(string),
			APPSECRET: data["appsecret"].(string),
			GRANTTYPE: data["granttype"].(string),
		})
		if res.Error != nil {
			fmt.Println("添加失败")
			return false
		}
	} else {
		res := db.Db.Where("id = ", id).Model(&WeChat{}).Update(data)
		if res.Error != nil {
			fmt.Println("编辑失败")
			return false
		}
	}
	return true
}

// @Summer 获取微信配置
func GetWeChatConfig() (weChat WeChat) {
	db.Db.First(&weChat)
	return
}
