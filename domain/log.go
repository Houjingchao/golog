package domain

import (
	"log"
	"github.com/Houjingchao/golog/consts"
    "github.com/ivpusic/grpool"
)

type Logger struct {
	level int
	err   *log.Logger
	warn  *log.Logger
	info  *log.Logger
	debug *log.Logger
	p     *grpool.Pool
}

func (ll *Logger) SetLevel(l int) {
	ll.level = l
}

func (ll *Logger) Error(format string, v ...interface{}) {
	if consts.LevelError > ll.level {
		return
	}
	ll.p.JobQueue <- func() {
		ll.err.Printf(format, v...)
	}
}

func NewGoLog(level int,) * Logger{
	return &Logger{

	}
}