package models

import "github.com/jinzhu/gorm"

func AutoData(Db *gorm.DB) bool {

	var count int

	errs := Db.Model(&LiteOauthUser{}).Count(&count).Error //判断是否存在数据

	if errs == nil && count == 0 {
		//新增一条数据
		Db.Create(&LiteOauthUser{
			Uid:      1,
			Nikename: "王大锤",
			Types:    1,
			Account:  "1234567@qq.com",
			Password: "93bcab4ab719fde430e5ad90656a240e", //1234@abcd
			Is_admin: 1,
			Status:   1,
		})
	}

	//同步表结构
	Db.AutoMigrate(&LiteOauthUser{}) //如果不存在就会生成一个新表

	Db.AutoMigrate(&LiteAdminMenu{}) //创建一张新的菜单表

	Db.AutoMigrate(&LiteTimeLine{}) //创建一张时间线表

	Db.AutoMigrate(&LiteArticleTag{}) //创建一张文章管理表

	Db.AutoMigrate(&LiteTag{}) //创建一张标签表

	Db.AutoMigrate(&LiteArticle{}) //创建一张文章

	Db.AutoMigrate(&LiteReview{}) //创建一张留言表

	Db.AutoMigrate(&LiteBase{}) //创建一张基础表

	Db.AutoMigrate(&LiteLog{}) //创建一张日志表

	Db.AutoMigrate(&LiteLink{}) //创建一张友情链接表

	Db.AutoMigrate(&LiteCrontab{}) //创建一张定时任务表

	Db.AutoMigrate(&LiteFiction{}) //小说列表上传

	Db.AutoMigrate(&LiteFictionLog{}) //小说日志

	Db.AutoMigrate(&LiteFictionOperation{}) //小说批次记录

	Db.AutoMigrate(&LiteBlackList{}) //黑名单

	return true
}
