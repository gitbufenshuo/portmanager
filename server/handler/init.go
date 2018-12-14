package handler

import (
	"fmt"
	"strings"

	"github.com/gitbufenshuo/portmanager/server/config"
)

// lua 脚本初始化什么的
func Init() {
	{
		keyTTLStr := fmt.Sprintf("%v", config.Conf.RedisConfig.KeyTTL)
		portBeginStr := fmt.Sprintf("%v", config.Conf.APP.PortBegin)
		// available_lua
		available_lua = strings.Replace(available_lua, "key_prefix", config.Conf.RedisConfig.KeyPrefix, -1)
		available_lua = strings.Replace(available_lua, "key_ttl", keyTTLStr, -1)
		available_lua = strings.Replace(available_lua, "port_begin", portBeginStr, -1)
		// heartbeat_lua
		heartbeat_lua = strings.Replace(heartbeat_lua, "key_prefix", config.Conf.RedisConfig.KeyPrefix, -1)
		heartbeat_lua = strings.Replace(heartbeat_lua, "key_ttl", keyTTLStr, -1)
	}
}
