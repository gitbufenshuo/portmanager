package handler

import (
	"fmt"
	"strings"

	"github.com/gitbufenshuo/portmanager/server/config"
)

// lua 脚本初始化什么的
func Init() {
	{ // available_lua
		keyTTLStr := fmt.Sprintf("%v", config.Conf.RedisConfig.KeyTTL)
		available_lua = strings.Replace(available_lua, "key_prefix", config.Conf.RedisConfig.KeyPrefix, -1)
		available_lua = strings.Replace(available_lua, "key_prefix", keyTTLStr, -1)
	}
}
