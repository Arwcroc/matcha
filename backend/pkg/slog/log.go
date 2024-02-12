package slog

import (
	"log"
	"sync"
)

type config struct {
	lock  *sync.RWMutex
	level Level
}

var conf = config{
	lock:  &sync.RWMutex{},
	level: LevelInfo,
}

func GetLogLevel() Level {
	conf.lock.RLock()
	defer conf.lock.RUnlock()
	return conf.level
}

func SetLogLevel(level Level) {
	conf.lock.Lock()
	defer conf.lock.Unlock()
	conf.level = level
}

func Log(level Level, v ...any) {
	if level < GetLogLevel() {
		return
	}
	log.Println(append([]interface{}{level.String()}, v...)...)
}

func Debug(v ...any) {
	Log(LevelDebug, v...)
}

func Info(v ...any) {
	Log(LevelInfo, v...)
}

func Warn(v ...any) {
	Log(LevelWarn, v...)
}

func Error(v ...any) {
	Log(LevelError, v...)
}
