package defaultimpl

import "fmt"

type Logger struct {
}

func (logger *Logger) Debug(msg string) {
	fmt.Println("[DEBUG]: " + msg)
}

func (logger *Logger) Info(msg string) {
	fmt.Println("[INFO]: " + msg)
}

func (logger *Logger) Warn(msg string) {
	fmt.Println("[WARN]: " + msg)
}

func (logger *Logger) Error(msg string) {
	fmt.Println("[ERROR]: " + msg)
}

func (logger *Logger) Fatal(msg string) {
	fmt.Println("[FATAL]: " + msg)
}

func (logger *Logger) Panic(msg string) {
	fmt.Println("[PANIC]: " + msg)
}