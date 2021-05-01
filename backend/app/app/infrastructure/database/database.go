package database

import (
	"fmt"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"time"
	"todo-list/app/config"
)

type Connection struct {
	DB *gorm.DB
}

func NewDB(conf *config.Config) (*Connection, error) {
	protocol := fmt.Sprintf("tcp(%s:%s)", conf.DB.Host, conf.DB.Port)
	connect := fmt.Sprintf(
		"%s:%s@%s/%s?parseTime=true&charset=utf8mb4",
		conf.DB.User,
		conf.DB.Pass,
		protocol,
		conf.DB.DBName,
	)
	conn, err := gorm.Open(mysql.Open(connect), &gorm.Config{})
	if err != nil {
		return nil, err
	}
	db, err := conn.DB()
	if err != nil {
		return nil, err
	}
	db.SetConnMaxLifetime(time.Minute * 5)

	return &Connection{DB: conn}, nil
}
