package main

import (
	"fmt"

	"github.com/spf13/viper"
)

func main() {
	viper.SetConfigName("toml")
	viper.SetConfigType("toml")
	viper.AddConfigPath(".")
	viper.AddConfigPath("./config")

	viper.SetDefault("redis.port", 6381)

	err := viper.ReadInConfig()

	if err != nil {
		fmt.Println("viper read in config error!")
		return
	}

	fmt.Println(viper.Get("app_name"))
	fmt.Println(viper.Get("log_level"))

	fmt.Println("server protocols: ", viper.GetStringSlice("server.protocols"))
	fmt.Println("server port : ", viper.GetIntSlice("server.ports"))

	if viper.IsSet("redis.port") {
		fmt.Println("redis.port is set")
	} else {
		fmt.Println("redis.port is not set")
	}

	fmt.Println("mysql setting: ", viper.GetStringMap("mysql"))
	fmt.Println("all setting: ", viper.AllSettings())
}
