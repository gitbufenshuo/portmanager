package handler

import (
	"net/http"

	"github.com/gitbufenshuo/portmanager/server/config"
	"github.com/gitbufenshuo/portmanager/server/ctxdata"
	"github.com/gitbufenshuo/portmanager/server/driver"
	"github.com/labstack/echo"
)

/*
	format:  /key_prefix/appname/hostname -> hostname:port
	format:  /key_prefix/appname/hostname/port -> 1
	keys:    []
	args:    [appname, hostname, port]
	# 1. /key_prefix/ARGV[1]/ARGV[2] 如果存在直接返回 -1
	# 2. 从 ARGV[3] 开始，自增，一直找到 /key_prefix/ARGV[1]/ARGV[2]/ARGV[3] 不存在的
	# 3. set /key_prefix/ARGV[1]/ARGV[2] ARGV[1]:ARGV[2] ex key_ttl
	# 4. set /key_prefix/ARGV[1]/ARGV[2]/ARGV[3] 1 ex key_ttl
*/
var available_lua = `
local app_prefix = '/'..key_prefix..'/'..ARGV[1]
local target_key = '/'..key_prefix..'/'..ARGV[1]..'/'..ARGV[2]
local l = redis.call('ttl', target_key)
if (l == -2)
	return -1
end
local port_begin = ARGV[3]
local port_chosen = -1;
for i = port_begin, port_begin+99 do
	local port_key = '/'..key_prefix..'/'..ARGV[1]..'/'..ARGV[2]..'/'..i
	l = redis.call('ttl', port_key)
	if (l == -2)
		port_chosen = i
		break
	end
end
if (port_chosen == -1)
	return -1
end
redis.call('set', target_key, ARGV[1]..':'..ARGV[2], ex, key_ttl)
local port_key = '/'..key_prefix..'/'..ARGV[1]..'/'..ARGV[2]..'/'..port_chosen
redis.call('set', port_key, 1, ex, key_ttl)
return port_chosen
`

// local l = redis.call('ttl', mykey)
// if (l == -2)
// then
// 	redis.call('set', mykey, ARGV[2]..':'..ARGV[3], 'ex', key_ttl)
// 	return ARGV[3]
// else
// 	local app_prefix = key_prefix..'/'..ARGV[1]

// end

func Available(c echo.Context) error {
	cc := c.Get("cc").(*ctxdata.Cusctx)
	redisClient := driver.RedisConf.GetRedisSession()
	hostname := c.Param("hostname")
	appname := c.Param("appname")
	//
	var portConfig config.PortConfig
	if v, found := config.Conf.APPList[appname]; found {
		portConfig = v
	} else {
		cc.Errf("no_such_app_%v", appname)
		return c.JSON(http.StatusBadRequest, "no_such_app")
	}
	if cmd := redisClient.Eval(available_lua, nil, appname, hostname, portConfig); cmd != nil {
		val, err := cmd.Result()
		if err != nil {
			cc.Errf("eval_err_%v", err)
		}
		cc.Logf("eval_val_%v", val)
		return c.JSON(http.StatusOK, val)
	}
	return nil
}
