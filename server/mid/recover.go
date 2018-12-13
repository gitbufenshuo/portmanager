package mid

import (
	"runtime"

	"github.com/gitbufenshuo/portmanager/server/ctxdata"
	"github.com/labstack/echo"
)

func RecoverMid(next echo.HandlerFunc) echo.HandlerFunc {
	return func(c echo.Context) error {
		defer func() {
			if r := recover(); r != nil {
				stack := make([]byte, 1024)
				runtime.Stack(stack, false)
				cc := c.Get("cc").(*ctxdata.Cusctx)
				cc.Errf("+++panicstack_\n%v\n%v\n---overstack", r, string(stack))
			}
		}()
		return next(c)
	}
}
