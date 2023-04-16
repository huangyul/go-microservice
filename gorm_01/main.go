package main

import (
	"fmt"
	"gorm.io/driver/mysql"
	"log"
	"os"
	"time"

	"gorm.io/gorm"
	"gorm.io/gorm/logger"
)

// Product 定义表结构
type Product struct {
	gorm.Model
	Code  string
	Price uint
}

func main() {
	// 参考 https://github.com/go-sql-driver/mysql#dsn-data-source-name 获取详情
	dsn := "root:root@tcp(47.106.214.127:3306)/gorm_test?charset=utf8mb4&parseTime=True&loc=Local"

	newLogger := logger.New(
		log.New(os.Stdout, "\r\n", log.LstdFlags), // io writer
		logger.Config{
			SlowThreshold:             time.Second, // Slow SQL threshold
			LogLevel:                  logger.Info, // Log level
			IgnoreRecordNotFoundError: true,        // Ignore ErrRecordNotFound error for logger
			ParameterizedQueries:      true,        // Don't include params in the SQL log
			Colorful:                  true,        // Disable color
		},
	)

	// Globally mode
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{
		Logger: newLogger,
	})

	if err != nil {
		panic(err)
	}
	fmt.Println("数据库连接成功")

	//_ = db.AutoMigrate(&Product{})

	db.Create(&Product{Code: "333", Price: 100})

	var product Product
	db.First(&product, 1)
	db.First(&product, "code = ?", "333")

	db.Model(&product).Update("Price", 200)

	db.Delete(&product, 1)
}
