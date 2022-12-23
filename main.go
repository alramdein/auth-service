package main

import (
	"fmt"

	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
	log "github.com/sirupsen/logrus"

	"github.com/alramdein/auth-service/config"
	httpDelivery "github.com/alramdein/auth-service/delivery/http"

	userClient "github.com/alramdein/user-service/client"
)

func main() {
	config.GetConf()
	// db, err := gorm.Open(postgres.Open(getDatabaseDSN()), &gorm.Config{})
	// if err != nil {
	// 	log.Error(err.Error())
	// 	panic("failed to connect database")
	// }

	e := echo.New()
	uc, err := userClient.NewClient(
		composeGRPCTarget(),
		config.GRPCTimeout(),
		config.GRPCIdleConnPoll(),
		config.GRPCMaxConnPoll(),
	)
	if err != nil {
		log.Error(err.Error())
		panic("failed to connect to grcp client")
	}

	httpDelivery.NewAuthHandler(e, uc)

	e.Use(middleware.Logger())
	e.Start(composeHTTPServerAddress())
}

func composeDatabaseDSN() string {
	return fmt.Sprintf("host=%s user=%s password=%s dbname=%s port=%s sslmode=disable TimeZone=Asia/Jakarta",
		config.DBHost(), config.DBUsername(), config.DBPassword(), config.DBName(), config.DBPort(),
	)
}

func composeHTTPServerAddress() string {
	return fmt.Sprintf("%v:%v", config.HTTPHost(), config.HTTPPort())
}

func composeGRPCTarget() string {
	return fmt.Sprintf("%s:%s", config.GRPCUserServiceHost(), config.GRPCUserServicePort())
}
