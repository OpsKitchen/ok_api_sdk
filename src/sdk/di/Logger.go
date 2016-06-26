package di

import "fmt"

type Logger struct {
}

func (logger *Logger) Debug(args ...interface{}) {
	fmt.Print("[DEBUG]: ")
	fmt.Println(args ...)
}

func (logger *Logger) Info(args ...interface{}) {
	fmt.Println("[INFO]: ")
	fmt.Println(args ...)
}

func (logger *Logger) Warn(args ...interface{}) {
	fmt.Println("[WARN]: ")
	fmt.Println(args ...)
}

func (logger *Logger) Error(args ...interface{}) {
	fmt.Println("[ERROR]: ")
	fmt.Println(args ...)
}

func (logger *Logger) Fatal(args ...interface{}) {
	fmt.Println("[FATAL]: ")
	fmt.Println(args ...)
}

func (logger *Logger) Panic(args ...interface{}) {
	fmt.Println("[PANIC]: ")
	fmt.Println(args ...)
}