package config

import (
	"log"
	"os"
)

type Logger struct {
	infoLogger  *log.Logger
	errorLogger *log.Logger
}

func NewLogger() *Logger {
	file, err := os.OpenFile(
		"app.log",
		os.O_CREATE|os.O_APPEND|os.O_WRONLY,
		0666,
	)

	if err != nil {
		log.Fatal(err)
	}

	info := log.New(
		file,
		"[INFO] ",
		log.Ldate|log.Ltime,
	)

	errors := log.New(
		os.Stdout,
		"\033[41m[ERROR]\033[0m",
		log.Ldate|log.Ltime,
	)

	return &Logger{
		infoLogger:  info,
		errorLogger: errors,
	}
}

func (l *Logger) Info(message string) {
	l.infoLogger.Println(message)
}

func (l *Logger) Error(message string) {
	l.errorLogger.Println(message)
}
