package settings

import (
	"fmt"
	"github.com/spf13/viper"
	"os"
)

//type application *Application

//var app application

func init() {
	fmt.Println("Determining config file location")
	switch os.Getenv("ENV") {
	case "dev":
		loadConfig("dev")
	case "stage":
		loadConfig("stage")
	case "prod":
		loadConfig("prod")
	default:
		loadConfig("dev")
	}
}

func loadConfig(env string) {
	viper.SetConfigName(".env." + env)
	viper.AddConfigPath(".")
	viper.SetConfigType("env")
	err := viper.ReadInConfig()
	if err != nil {
		panic(err)
	}
	viper.AutomaticEnv()

	fmt.Printf("Config path: %s", viper.ConfigFileUsed())

	//app = &Application{
	//	Server: Server{
	//		Port:     viper.GetInt("APP_PORT"),
	//		Hostname: viper.GetString("APP_HOSTNAME"),
	//		LogLevel: viper.GetString("LOG_LEVEL"),
	//		GinMode:  viper.GetString("GIN_MODE"),
	//	},
	//	Database: Database{
	//		Hostname: viper.GetString("DB_HOSTNAME"),
	//		Port:     viper.GetString("DB_PORT"),
	//		Username: viper.GetString("DB_USERNAME"),
	//		Password: viper.GetString("DB_PASSWORD"),
	//		Name:     viper.GetString("DB_NAME"),
	//	},
	//	JWT: JWT{
	//		Secret: viper.GetString("JWT_SECRET"),
	//	},
	//}

	// Ensure no secrets are printed in logs except in dev mode
	if viper.GetString("LOG_LEVEL") == "debug" && env == "dev" {
		fmt.Printf("Config: %+v", viper.AllSettings())
	}
}

//func GetConfig() *Application {
//	return app
//}
