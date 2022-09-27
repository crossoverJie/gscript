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

func (l *Log) String() string {
	line := l.ctx.GetStart().GetLine() - NativeLine
	return fmt.Sprintf("%d:%d: %s", line, l.ctx.GetStart().GetColumn(), l.msg)
}
