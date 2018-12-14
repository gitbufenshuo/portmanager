package handler

import (
	"net/http"

	"github.com/gitbufenshuo/portmanager/server/ctxdata"
	"github.com/gitbufenshuo/portmanager/server/driver"
	"github.com/labstack/echo"
)

/*
	format:  /key_prefix/app/appname/hostname -> hostname:port
	format:  /key_prefix/port/hostname/port -> appname
	keys:    []
	args:    [appname, hostname]
	# 1. get /key_prefix/app/appname/hostname 不存在返回 -1
	# 2. get /key_prefix/port/hostname/port appname 不存在返回 -1
	# 3.
	# 4.
*/
var heartbeat_lua = `
local appname = ARGV[1]
local hostname = ARGV[2]
local app_key = string.format("/key_prefix/app/%s/%s", appname, hostname)

local app_value = redis.call('get', app_key)
if (app_value == nil)
then
	return -1
end
local port_key = string.format("/key_prefix/port/%s", app_value)
port_key = string.gsub(port_key, ":", "/")
local port_value = redis.call('get', port_key)
if (port_value == nil)
then
	return -2
end
redis.call('set', app_key, app_value, 'ex', key_ttl)
redis.call('set', port_key, port_value, 'ex', key_ttl)
return 1
`

func Heartbeat(c echo.Context) error {
	cc := c.Get("cc").(*ctxdata.Cusctx)
	redisClient := driver.RedisConf.GetRedisSession()
	hostname := c.Param("hostname")
	appname := c.Param("appname")
	//
	if cmd := redisClient.Eval(heartbeat_lua, nil, appname, hostname); cmd != nil {
		val, err := cmd.Result()
		if err != nil {
			cc.Errf("eval_err_%v", err)
			return c.JSON(http.StatusBadRequest, err.Error())
		}
		cc.Logf("eval_val_%v", val)
		return c.JSON(http.StatusOK, val)
	}
	return nil
}
