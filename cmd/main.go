package main

import (
	"fmt"
	"transfer-service/config"
	"transfer-service/repo"
	"transfer-service/repo/db"
	"transfer-service/router"
	"transfer-service/service"

	"github.com/gin-gonic/gin"
	"github.com/golang-migrate/migrate/v4"
	_ "github.com/golang-migrate/migrate/v4/database/postgres"
	_ "github.com/golang-migrate/migrate/v4/source/file"
)

func main() {
	config.LoadConfig()
	dbCfg := config.DatabaseConfig{}
	config.ReadConfig(config.DatabaseConfigPath, &dbCfg)

	migration, err := migrate.New(
		"file://./migration",
		fmt.Sprintf("postgres://%s:%s@%s:%d/%s?sslmode=disable", dbCfg.Username, dbCfg.Password, dbCfg.Host, dbCfg.Port, dbCfg.Database),
	)
	if err != nil {
		panic(err)
	}
	err = migration.Up()
	if err != nil && err != migrate.ErrNoChange {
		panic(err)
	}

	db := db.NewGormDB(dbCfg)
	repo := repo.NewGormRepo(db)
	service := service.NewAccountService(repo)

	engine := gin.Default()
	accountHandler := router.NewAccountHandler(service)
	accountHandler.Register(engine)

	_ = engine.Run("0.0.0.0:8080")
}
