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
	# 1. /key_prefix/app/appname/hostname 如果存在直接返回 -1
	# 2. 从 port_begin 开始，自增，一直找到 /key_prefix/port/hostname/{port} 不存在的
	# 3. set /key_prefix/app/appname/hostname hostname:port ex key_ttl
	# 4. set /key_prefix/port/hostname/port appname ex key_ttl
*/
var available_lua = `
local target_key = '/'..'key_prefix/app'..'/'..ARGV[1]..'/'..ARGV[2]
local l = redis.call('ttl', target_key)
if (l ~= -2)
then
	return -1
end
local port_chosen = -1;
for i = port_begin, port_begin+99 do
	local port_key = '/'..'key_prefix/port'..'/'..ARGV[2]..'/'..i
	l = redis.call('ttl', port_key)
	if (l == -2)
	then
		port_chosen = i
		break
	end
end
if (port_chosen == -1)
then
	return -1
end
redis.call('set', target_key, ARGV[2]..':'..port_chosen, 'ex', key_ttl)
local port_key = '/'..'key_prefix/port'..'/'..ARGV[2]..'/'..port_chosen
redis.call('set', port_key, ARGV[1], 'ex', key_ttl)
return port_chosen
`

func Available(c echo.Context) error {
	cc := c.Get("cc").(*ctxdata.Cusctx)
	redisClient := driver.RedisConf.GetRedisSession()
	hostname := c.Param("hostname")
	appname := c.Param("appname")
	//
	if cmd := redisClient.Eval(available_lua, nil, appname, hostname); cmd != nil {
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
