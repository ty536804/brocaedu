package WeChat

import (
	db "brocaedu/Database"
	"fmt"
	"time"
)

type LookHistory struct {
	Id        int    `json:"id" gorm:"primary_key"`
	Url       string `json:"url" gorm:"type:varchar(100);not null;default '';comment:'观看地址' "`
	UserId    int    `json:"user_id" gorm:"not null;default 0;comment:'微信用户ID'"`
	CreatedAt string `json:"created_at" time_format:"2006-01-02 15:04:05"`
	UpdatedAt string `json:"updated_at" time_format:"2006-01-02 15:04:05"`
}

// @Summer 添加观看记录
func AddLook(data map[string]interface{}) bool {
	CreatedAt := time.Now().Format("2006-01-02 15:04:05")
	res := db.Db.Create(&LookHistory{
		Url:       data["url"].(string),
		UserId:    data["user_id"].(int),
		CreatedAt: CreatedAt,
		UpdatedAt: CreatedAt,
	})
	if res.Error != nil {
		fmt.Println("观看记录添加失败", res.Error.Error())
		return false
	}
	return true
}
