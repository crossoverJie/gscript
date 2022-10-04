package log

import (
	"fmt"
	"github.com/antlr/antlr4/runtime/Go/antlr"
)

var NativeLine int

type Log struct {
	ctx antlr.ParserRuleContext
	msg string
}

func NewLog(ctx antlr.ParserRuleContext, msg string) *Log {
	return &Log{ctx: ctx, msg: msg}
}
func RuntimePanic(ctx antlr.ParserRuleContext, msg string) {
	log := NewLog(ctx, msg)
	panic(log)
}

func (l *Log) String() string {
	line := l.ctx.GetStart().GetLine() - NativeLine + 1
	return fmt.Sprintf("%d:%d: %s", line, l.ctx.GetStart().GetColumn(), l.msg)
}
