package mysql

import (
	"database/sql"
	"demo-test/internal/models"
	"fmt"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

var MysqlDB *sql.DB

const (
	MySQLDSN = "root:123456@tcp(127.0.0.1:3306)/test"
)

func NewMysqlDB(dsn string) (*sql.DB, error) {
	db, err := sql.Open("mysql", dsn)
	if err != nil {
		return nil, err
	}
	return db, nil
}

func InitMysqlDB() {
	dsn := "root:@tcp(127.0.0.1:3306)/myapp?charset=utf8mb4&parseTime=True&loc=Local"
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		panic("failed to connect database")
	}

	// 自动迁移（可选，如果表结构已存在可省略）
	db.AutoMigrate(&models.User{})

	// 插入测试数据
	db.Create(&models.User{Name: "Alice", Email: "alice@example.com", Age: 25})

	// 查询并打印
	var users []models.User
	db.Find(&users)
	for _, u := range users {
		fmt.Printf("User: %+v\n", u)
	}
}
