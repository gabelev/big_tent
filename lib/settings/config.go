package settings

import (
	"fmt"
	"os"

	"github.com/spf13/viper"
)

var Config struct {
	DbHost     string
	DbName     string
	DbUser     string
	DbPassword string
	AppPort    string
}

func SetConfig() {
	viper.SetConfigName("config")
	viper.SetConfigType("json")
	viper.AddConfigPath(
		"$GOPATH/src/github.com/gabelev/get_heard/",
	)
	if err := viper.ReadInConfig(); err == nil {
		fmt.Println("Processing Settings")
	} else {
		fmt.Println("Config Not Found")
		fmt.Printf(err.Error())
		os.Exit(-1)
	}
	setBaseConfig()
	fmt.Println("Let the Roucus Caucus Begin!")
}

func setBaseConfig() {
	Config.DbHost = viper.GetString("db_host")
	Config.DbName = viper.GetString("db_name")
	Config.DbUser = viper.GetString("db_user")
	Config.DbPassword = viper.GetString("db_password")
	Config.AppPort = viper.GetString("app_port")
}
