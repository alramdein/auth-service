package config

import (
	"fmt"
	"time"

	"github.com/spf13/viper"
)

func GetConf() {
	viper.SetConfigName("config")
	viper.SetConfigType("yml")
	viper.AddConfigPath(".")
	err := viper.ReadInConfig()
	if err != nil {
		panic(fmt.Errorf("fatal error config file: %w ", err))
	}
}

func HTTPHost() string {
	return viper.GetString("http.host")
}

func HTTPPort() string {
	return viper.GetString("http.port")
}

func GRPCHost() string {
	return viper.GetString("grpc.host")
}

func GRPCPort() string {
	return viper.GetString("grpc.port")
}

func GRPCTimeout() time.Duration {
	return viper.GetDuration("grpc.timeout")
}

func GRPCIdleConnPoll() int {
	return viper.GetInt("grpc.idle_conn_pool")
}

func GRPCMaxConnPoll() int {
	return viper.GetInt("grpc.max_conn_pool")
}

func GRPCUserServiceHost() string {
	return viper.GetString("grpc.user_service_host")
}

func GRPCUserServicePort() string {
	return viper.GetString("grpc.user_service_port")
}

func DBHost() string {
	return viper.GetString("database.host")
}

func DBPort() string {
	return viper.GetString("database.port")
}

func DBUsername() string {
	return viper.GetString("database.username")
}

func DBPassword() string {
	return viper.GetString("database.password")
}

func DBName() string {
	return viper.GetString("database.dbname")
}

func JWTSecret() string {
	return viper.GetString("jwt_secret")
}
