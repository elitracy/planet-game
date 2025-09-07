package logging

import (
	"gopkg.in/natefinch/lumberjack.v2"
	"log"
	"time"
)

const DEFAULT_LOG_LEVEL = "DEBUG"

type LogMessage struct {
	Time    time.Time
	Level   string
	Package string
	Message string
}

type Logger struct {
	queue    chan LogMessage
	filepath string
}

var logger *Logger

func (l *Logger) run() {
	log.SetOutput(
		&lumberjack.Logger{
			Filename:   l.filepath,
			MaxSize:    30,
			MaxBackups: 3,
			MaxAge:     14,
			Compress:   true,
		},
	)

	for msg := range l.queue {
		log.Printf("[%s] [%s] %s\n", msg.Level, msg.Package, msg.Message)
	}
}

func NewLogger(filepath string) *Logger {
	l := &Logger{
		queue:    make(chan LogMessage, 10),
		filepath: filepath,
	}

	go l.run()
	return l
}

func Log(message, pkg string, level ...string) {
	msgLevel := DEFAULT_LOG_LEVEL
	if len(level) > 0 && level[0] != "" {
		msgLevel = level[0]
	}

	logger.queue <- LogMessage{
		Time:    time.Now(),
		Level:   msgLevel,
		Package: pkg,
		Message: message,
	}
}

func init() {
	logger = NewLogger("logs/debug.log")
}
