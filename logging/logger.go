package logging

import (
	"fmt"
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
	queue chan LogMessage
}

var logger *Logger

func (l *Logger) run() {
	for msg := range l.queue {
		fmt.Printf("[%s] [%s] [%s] %s\n", msg.Time.Format(time.RFC3339), msg.Level, msg.Package, msg.Message)
	}
}

func NewLogger() *Logger {
	l := &Logger{
		queue: make(chan LogMessage, 10),
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
	logger = NewLogger()
}
