package ctxdata

import (
	"bytes"

	"fmt"

	"github.com/labstack/echo"
)

type Cusctx struct {
	C echo.Context
	B *bytes.Buffer
}

func New(b *bytes.Buffer) *Cusctx {
	var cc Cusctx
	cc.B = b
	return &cc
}
func (cc *Cusctx) Logf(format string, data ...interface{}) {
	if cc.B == nil {
		fmt.Printf("L:"+format+":::", data...)
		return
	}
	cc.B.WriteString(fmt.Sprintf("L:"+format, data...))
	cc.B.WriteString(":::")
}
func (cc *Cusctx) Errf(format string, data ...interface{}) {
	if cc.B == nil {
		fmt.Printf("E:"+format+":::", data...)
		return
	}
	cc.B.WriteString(fmt.Sprintf("E:"+format, data...))
	cc.B.WriteString(":::")
}
