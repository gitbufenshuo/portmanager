package main

import (
	"fmt"
	"os"

	"github.com/BurntSushi/toml"
	"github.com/gitbufenshuo/portmanager/server/config"
	"github.com/gitbufenshuo/portmanager/server/driver"
	"github.com/gitbufenshuo/portmanager/server/handler"
	"github.com/gitbufenshuo/portmanager/server/mid"
	"github.com/labstack/echo"
)

func main() {
	{
		if _, err := toml.DecodeFile(os.Getenv("filename"), &config.Conf); err != nil {
			fmt.Println(err)
			os.Exit(1)
		}
	}
	{
		var conf driver.RedisConfig
		conf.HostPort = config.Conf.RedisConfig.ServerPort
		conf.Password = config.Conf.RedisConfig.Password
		if err := driver.Init(&conf); err != nil {
			fmt.Println(err)
			os.Exit(1)
		}
	}
	e := echo.New()
	e.Use(mid.SuperCtx(mid.DefaultLoggerConfig))
	e.Use(mid.RecoverMid)
	e.GET("/"+config.Conf.API.HTTPPrefix+"/available/:hostname/:appname", handler.Available)
	e.GET("/"+config.Conf.API.HTTPPrefix+"/heartbeat/:hostname/:appname", handler.Heartbeat)

	e.Start(os.Getenv("bind"))
}
