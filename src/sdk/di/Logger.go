package di

type Logger interface {
	Debug()
	Info()
	Warn()
	Error()
	Fatal()
	Panic()
}