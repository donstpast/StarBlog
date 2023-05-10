package model

import (
	"fmt"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"starblog/utils"
	"time"
)

var DB *gorm.DB

// InitDb 数据库入口文件
func InitDb() {
	var err error
	dsn := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?charset=utf8mb4&parseTime=True&loc=Local",
		utils.DbUser,
		utils.DbPassWord,
		utils.DbHost,
		utils.DbPort,
		utils.DbName,
	)
	DB, err = gorm.Open(mysql.New(mysql.Config{
		DSN:                       dsn,   // DSN data source name
		DefaultStringSize:         256,   // string 类型字段的默认长度
		DisableDatetimePrecision:  true,  // 禁用 datetime 精度，MySQL 5.6 之前的数据库不支持
		DontSupportRenameIndex:    true,  // 重命名索引时采用删除并新建的方式，MySQL 5.7 之前的数据库和 MariaDB 不支持重命名索引
		DontSupportRenameColumn:   true,  // 用 `change` 重命名列，MySQL 8 之前的数据库和 MariaDB 不支持重命名列
		SkipInitializeWithVersion: false, // 根据当前 MySQL 版本自动配置
	}), &gorm.Config{})
	if err != nil {
		fmt.Println("数据库连接失败，请检查参数：", err)
	}
	//AutoMigrate
	//AutoMigrate 用于自动迁移您的 schema，保持您的 schema 是最新的。
	//用模型创建的表中缺失的列，缺失的表都会根据数据库进行更新，但是用模型建立的表会自动加复数

	err = DB.AutoMigrate(&User{}, &Category{}, &Article{}, &Comment{}, &Friends{}, &Tag{}, &Visitor{})
	if err != nil {
		fmt.Println("数据库自动迁移失败，请检查：", err)
	}

	//连接池
	// 获取通用数据库对象 sql.DB ，然后使用其提供的功能
	sqlDB, _ := DB.DB()

	// SetMaxIdleConns 用于设置连接池中空闲连接的最大数量。
	sqlDB.SetMaxIdleConns(10)

	// SetMaxOpenConns 设置打开数据库连接的最大数量。
	sqlDB.SetMaxOpenConns(100)

	// SetConnMaxLifetime 设置了连接可复用的最大时间。
	sqlDB.SetConnMaxLifetime(10 * time.Second)

}
