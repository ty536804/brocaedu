package Banner

import (
	db "brocaedu/Database"
	"brocaedu/Models/Nav"
	"brocaedu/Pkg/setting"
	"fmt"
	"log"
	"time"
)

type Banner struct {
	db.Model

	Navs Nav.Nav `json:"nav" gorm:"FOREIGNKEY:Bposition;ASSOCIATION_FOREIGNKEY:ID"`

	Province   string    `json:"province" gorm:"type:varchar(190);not null;default '0';comment:'省'"`
	City       string    `json:"city" gorm:"type:varchar(190);not null;default '0';comment:'市'"`
	Area       string    `json:"area" gorm:"type:varchar(190);not null;default '0';comment:'区'"`
	Bname      string    `json:"bname" gorm:"type:varchar(190);not null;default '';comment:'名称'"`
	Bposition  int       `json:"bposition" gorm:"index;comment:'位置'"`
	Imgurl     string    `json:"imgurl" gorm:"type:varchar(190);not null;default '';comment:'图片地址'"`
	TargetLink string    `json:"target_link" gorm:"type:varchar(190);not null;default '';comment:'跳转链接'"`
	BeginTime  time.Time `json:"begin_time" time_format:"2006-01-02 15:04:05" gorm:"default '';comment:'显示开始时间'"`
	EndTime    time.Time `json:"end_time" time_format:"2006-01-02 15:04:05" gorm:"default '';comment:'显示结束时间'"`
	IsShow     int       `json:"is_show" gorm:"default '1';comment:'状态 1显示 2隐藏'"`
	ImageSize  string    `json:"image_size" gorm:"type:varchar(190);not null;default '';comment:'图片大小 长*高*宽'"`
	Info       string    `json:"info" gorm:"type:varchar(255);not null;default '';comment:'备注'"`
	Tag        string    `json:"tag" gorm:"type:varchar(190);not null;default '';comment:'标签'"`
	Type       int       `json:"type" gorm:"not null;default '1';comment:'1 PC 2 WAP'"`
}

// @Summer 添加banner
func AddBanner(data map[string]interface{}) bool {
	startTime := time.Now().Add(100 * time.Hour)
	err := db.Db.Create(&Banner{
		Province:   "10000",
		City:       "0",
		Area:       "0",
		Bname:      data["bname"].(string),
		Bposition:  data["bposition"].(int),
		Imgurl:     data["imgurl"].(string),
		Info:       data["info"].(string),
		TargetLink: data["target_link"].(string),
		IsShow:     data["is_show"].(int),
		Tag:        data["tag"].(string),
		Type:       data["type"].(int),
		BeginTime:  startTime,
		EndTime:    startTime,
	})

	if err.Error != nil {
		log.Printf("添加banner失败,%v", err)
		return false
	}
	return true
}

// @Summer 编辑banner
func EditBanner(id int, data interface{}) bool {
	edit := db.Db.Model(&Banner{}).Where("id = ?", id).Update(data)
	if edit.Error != nil {
		fmt.Print("编辑banner错误:", edit)
		return false
	}
	return true
}

func GetOneBanner(id, clientType int, tag string) (banner Banner) {
	db.Db.Where("bposition = ? and type = ? and tag =? ", id, clientType, tag).First(&banner)
	return
}

// @Summer获取所有banner
func GetBanners(page int) (banner []Banner) {
	offset := 0
	if page >= 1 {
		offset = (page - 1) * setting.PageSize
	}
	db.Db.Preload("Navs").Offset(offset).Limit(setting.PageSize).Order("id desc").Find(&banner)
	return
}

// @Summer 统计图片总数
func GetBannerTotal() (count int) {
	db.Db.Model(&Banner{}).Count(&count)
	return
}

// @Summer 获取图片列表
func GetBanner(id int) (banner Banner) {
	db.Db.Preload("Navs").Where("id = ?", id).First(&banner)
	return
}

// @Summer获取所有banner
func GetBannerData(bposition, clientType int) (banner []Banner) {
	db.Db.Where("bposition = ? and type = ? and tag = ? and is_show=1", bposition, clientType, "banner").Find(&banner)
	return
}

// @Summer获取所有banner
func GetData(bposition, posi int) (banner []Banner) {
	db.Db.Where("bposition = ? and posi= ?", bposition, posi).Order("sort desc").Find(&banner)
	return
}

// @Summer获取所有banner
func GetBannerByTag(poi, clientType int, tag string) (banner []Banner) {
	db.Db.Where("bposition = ? and type = ? and tag = ?", poi, clientType, tag).Find(&banner)
	return
}

// @Summer 删除banner
func DelBanner(id int) bool {
	if id < 1 {
		return false
	}
	err := db.Db.Delete(&Banner{}, "id =? ", id)
	if err.Error != nil {
		log.Printf("删除banner失败,%v", err)
		return false
	}
	return true
}

// @Summer通过描述获取图片
func GetBannerList(info string) (banner []Banner) {
	db.Db.Where("info = ?", info).Find(&banner)
	return
}
