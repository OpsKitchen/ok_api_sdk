package logger

import (
	"fmt"
	"os"
)

type Logger struct {
	Level Level
}

func (logger *Logger) SetLevel(level Level) {
	logger.Level = level
}

func (logger *Logger) Debug(args ...interface{}) {
	if logger.Level >= DebugLevel {
		fmt.Print("[DEBUG]: ")
		fmt.Println(args...)
	}
}

func (logger *Logger) Info(args ...interface{}) {
	if logger.Level >= InfoLevel {
		fmt.Println("[INFO]: ")
		fmt.Println(args...)
	}
}

func (logger *Logger) Warn(args ...interface{}) {
	if logger.Level >= WarnLevel {
		fmt.Println("[WARN]: ")
		fmt.Println(args...)
	}
}

func (logger *Logger) Error(args ...interface{}) {
	if logger.Level >= ErrorLevel {
		fmt.Println("[ERROR]: ")
		fmt.Println(args...)
	}
}

func (logger *Logger) Fatal(args ...interface{}) {
	if logger.Level >= FatalLevel {
		fmt.Println("[FATAL]: ")
		fmt.Println(args...)
	}
	os.Exit(1)
}

func (logger *Logger) Panic(args ...interface{}) {
	if logger.Level >= PanicLevel {
		fmt.Println("[PANIC]: ")
		fmt.Println(args...)
	}
}
