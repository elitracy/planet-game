package logging

import (
	"fmt"
	"log"
	"path/filepath"
	"runtime"
	"strings"
	"time"

	"github.com/elitracy/planets/models"
	"gopkg.in/natefinch/lumberjack.v2"
)

const (
	colorReset  = "\033[0m"
	colorRed    = "\033[31m"
	colorGreen  = "\033[32m"
	colorYellow = "\033[33m"
	colorBlue   = "\033[34m"
	colorGrey   = "\033[90m"
)

type LogMessage struct {
	Time     time.Time
	Tick     int
	Color    string
	Level    string
	Filename string
	Message  string
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
	log.SetFlags(0)

	for msg := range l.queue {
		timeTick := fmt.Sprintf("%s%s|%05d%s", colorGrey, msg.Time.Format("15:04:05.000"), msg.Tick, colorReset)
		log.Printf("%s %s[%s] %s %s%s\n", timeTick, msg.Color, msg.Level, msg.Filename, msg.Message, colorReset)
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

func (l *Logger) log(level, color, format string, args ...any) {
	_, file, _, ok := runtime.Caller(2)
	fileName := "UNKNOWN"
	if ok {
		fileName = strings.ToUpper(filepath.Base(file))
	}

	msg := fmt.Sprintf(format, args...)
	logger.queue <- LogMessage{
		Time:     time.Now(),
		Tick:     models.GameStateGlobal.CurrentTick,
		Level:    level,
		Filename: fileName,
		Message:  msg,
		Color:    color,
	}
}

func Info(format string, args ...any)  { logger.log("TELEMETRY", colorReset, format, args...) }
func Error(format string, args ...any) { logger.log("FAULT", colorRed, format, args...) }
func Warn(format string, args ...any)  { logger.log("WARN", colorYellow, format, args...) }
func Ok(format string, args ...any)    { logger.log("STABLE", colorGreen, format, args...) }

func init() {
	logger = NewLogger("logs/debug.log")
}
