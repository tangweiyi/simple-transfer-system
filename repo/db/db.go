package db

import (
	"fmt"
	"log"
	"time"
	"transfer-service/config"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
	"gorm.io/gorm/schema"
)

func NewGormDB(cfg config.DatabaseConfig) *gorm.DB {
	//idle_in_transaction_session_timeout/statement_timeout/lock_timeout=10000
	dsn := fmt.Sprintf("host=%s user=%s password=%s dbname=%s port=%d sslmode=disable TimeZone=Asia/Shanghai",
		cfg.Host, cfg.Username, cfg.Password, cfg.Database,
		cfg.Port)

	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{
		Logger: logger.Default,
		NowFunc: func() time.Time {
			return time.Now().UTC()
		},
		NamingStrategy: schema.NamingStrategy{
			SingularTable: true,
		},
	})

	if err != nil {
		log.Fatal(err)
	}

	sqlDB, _ := db.DB()

	sqlDB.SetMaxOpenConns(cfg.MaxConnections)
	sqlDB.SetMaxIdleConns(cfg.MaxConnections)
	sqlDB.SetConnMaxIdleTime(cfg.MaxIdleTime)

	log.Printf("Gorm loaded...")
	return db
}
