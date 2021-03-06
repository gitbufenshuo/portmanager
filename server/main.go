package main

import (
	"fmt"
	"os"
	"os/signal"
	"syscall"
	"time"

	"github.com/BurntSushi/toml"
	"github.com/gitbufenshuo/portmanager/server/config"
	"github.com/gitbufenshuo/portmanager/server/driver"
	"github.com/gitbufenshuo/portmanager/server/handler"
	"github.com/gitbufenshuo/portmanager/server/mid"
	"github.com/labstack/echo"
)

func main() {
	{ // global_config
		if _, err := toml.DecodeFile(os.Getenv("filename"), &config.Conf); err != nil {
			fmt.Println(err)
			os.Exit(1)
		} else {
			fmt.Printf("%v\n", config.Conf.String())
		}
	}
	{ // redis_config
		conf := &driver.RedisConf
		conf.HostPort = config.Conf.RedisConfig.ServerPort
		conf.Password = config.Conf.RedisConfig.Password
		conf.Database = config.Conf.RedisConfig.Database
		if err := driver.Init(conf); err != nil {
			fmt.Println(err)
			os.Exit(1)
		}
	}
	{ // handlers_config
		handler.Init()
	}
	e := echo.New()
	e.Use(mid.SuperCtx(mid.DefaultLoggerConfig))
	e.Use(mid.RecoverMid)
	e.GET("/"+config.Conf.API.HTTPPrefix+"/available/:hostname/:appname", handler.Available)
	e.GET("/"+config.Conf.API.HTTPPrefix+"/heartbeat/:hostname/:appname", handler.Heartbeat)
	go serve(e)
	{
		// 接收操作系统 signal
		c := make(chan os.Signal)
		signal.Notify(c, syscall.SIGTERM, syscall.SIGINT)
		<-c
		fmt.Println(">>>>>>>> stop server <<<<<<<<", time.Now().Unix())
	}
	e.Close()
	driver.RedisConf.Close()
	fmt.Println(">>>>>>>> stop server success <<<<<<<<", time.Now().Unix())
}
func serve(e *echo.Echo) {
	err := e.Start(os.Getenv("bind"))
	fmt.Println("serve:", err)
}
