package mysql

import (
	"fmt"
	"log"
	"time"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

// NewMySQLConn 建立資料庫連線物件
func NewMySQLConn(user string, password string, host string, port string, database string, idle, maxConn, lifeTime int, debug bool) (db *gorm.DB, err error) {

	// 連線資訊
	connectionString := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?charset=utf8mb4&parseTime=true&loc=UTC",
		user,
		password,
		host,
		port,
		database)

	// 建立連線
	db, err = gorm.Open(mysql.Open(connectionString), &gorm.Config{
		QueryFields: true,
	})
	if err != nil {
		log.Fatalf("database connection failed : %v \n", err)
	}

	if db.Error != nil {
		log.Fatalf("database error %v \n", db.Error)
	}

	// TODO 視情境調整 設定連線池
	sqlDB, err := db.DB()
	if err != nil {
		log.Fatalf("database pool setting failed : %v \n", err)
	}

	sqlDB.SetMaxIdleConns(idle)
	sqlDB.SetMaxOpenConns(maxConn)
	sqlDB.SetConnMaxLifetime(time.Duration(lifeTime) * time.Hour)

	err = sqlDB.Ping()
	if err != nil {
		log.Fatalf("database ping err:%v\n", err)
	}

	if debug {
		db = db.Debug()
	}

	return
}
