package main

import (
	"fmt"
	"net"

	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
	log "github.com/sirupsen/logrus"
	"google.golang.org/grpc"

	"github.com/alramdein/auth-service/config"
	grpcDelivery "github.com/alramdein/auth-service/delivery/grpc"
	httpDelivery "github.com/alramdein/auth-service/delivery/http"
	"github.com/alramdein/auth-service/usecase"

	"github.com/alramdein/auth-service/pb"
	userClient "github.com/alramdein/user-service/client"
)

func main() {
	config.GetConf()

	authUsecase := usecase.NewAuthUsecase()

	// run http server in goroutine
	go func() {
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
	}()

	// grpc server
	lis, err := net.Listen("tcp", fmt.Sprintf(":%s", config.GRPCPort()))
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
		return
	}

	grpcServer := grpc.NewServer()

	grpcsvc := grpcDelivery.NewUserService()
	grpcsvc.RegisterAuthUsecase(authUsecase)

	pb.RegisterAuthServiceServer(grpcServer, grpcsvc)

	log.Info("grpc listen from ", config.GRPCPort())
	if err := grpcServer.Serve(lis); err != nil {
		log.Fatalf("failed to serve: %s", err)
		return
	}

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
