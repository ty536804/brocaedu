package Commands

import (
	db "brocaedu/Database"
	"brocaedu/Models/Admin"
	"brocaedu/Models/Article"
	"brocaedu/Models/Banner"
	"brocaedu/Models/Campus"
	"brocaedu/Models/Message"
	"brocaedu/Models/Nav"
	"brocaedu/Models/Single"
	"brocaedu/Models/Site"
	"brocaedu/Models/Visit"
	"fmt"
)

func init() {
	fmt.Println("生成数据库文件")
	InitAdminDatabase()
}
func InitAdminDatabase() {
	//DropDatabase()
	db.Db.AutoMigrate(&Admin.SysAdminUser{})
	db.Db.AutoMigrate(&Admin.SysAdminDepartment{})
	db.Db.AutoMigrate(&Admin.SysAdminPosition{})
	db.Db.AutoMigrate(&Admin.SysAdminPower{})
	db.Db.AutoMigrate(&Banner.Banner{})
	db.Db.AutoMigrate(&Article.Article{})
	db.Db.AutoMigrate(&Banner.Banner{})
	db.Db.AutoMigrate(&Message.Message{})
	db.Db.AutoMigrate(&Nav.Nav{})
	db.Db.AutoMigrate(&Site.Site{})
	db.Db.AutoMigrate(&Single.Single{})
	db.Db.AutoMigrate(&Campus.Campus{})
	db.Db.AutoMigrate(&Visit.Visit{})
}

func DropDatabase() {
	if db.Db.HasTable(&Admin.SysAdminUser{}) {
		db.Db.DropTable(&Admin.SysAdminUser{})
	}
	if !db.Db.HasTable(&Admin.SysAdminDepartment{}) {
		db.Db.DropTable(&Admin.SysAdminDepartment{})
	}
	if !db.Db.HasTable(&Admin.SysAdminPosition{}) {
		db.Db.DropTable(&Admin.SysAdminPosition{})
	}
	if !db.Db.HasTable(&Admin.SysAdminPower{}) {
		db.Db.DropTable(&Admin.SysAdminPower{})
	}
	if !db.Db.HasTable(&Banner.Banner{}) {
		db.Db.DropTable(&Banner.Banner{})
	}
	//文章
	if !db.Db.HasTable(&Article.Article{}) {
		db.Db.DropTable(&Article.Article{})
	}
	//轮播图
	if !db.Db.HasTable(&Banner.Banner{}) {
		db.Db.DropTable(&Banner.Banner{})
	}
	//信息
	if !db.Db.HasTable(&Message.Message{}) {
		db.Db.DropTable(&Message.Message{})
	}
	//导航
	if !db.Db.HasTable(&Nav.Nav{}) {
		db.Db.DropTable(&Nav.Nav{})
	}
	//站点
	if !db.Db.HasTable(&Site.Site{}) {
		db.Db.DropTable(&Site.Site{})
	}
	//单页
	if !db.Db.HasTable(&Single.Single{}) {
		db.Db.DropTable(&Single.Single{})
	}
	//校园管理
	if !db.Db.HasTable(&Campus.Campus{}) {
		db.Db.DropTable(&Campus.Campus{})
	}
	//浏览记录
	if !db.Db.HasTable(&Visit.Visit{}) {
		db.Db.DropTable(&Visit.Visit{})
	}
}