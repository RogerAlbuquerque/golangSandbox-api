package config

import (
	"fmt"
	"log"
	"os"
	"path/filepath"
	"runtime"
	"time"
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
		log.Lshortfile,
	)

	errors := log.New(
		os.Stdout,
		"\033[41m[ERROR]\033[0m >> ",
		log.Llongfile,
	)

	return &Logger{
		infoLogger:  info,
		errorLogger: errors,
	}
}

func (l *Logger) Info(message string) {
	_, file, line, ok := runtime.Caller(1)

	if !ok {
		file = "unknown"
		line = 0
	}

	filename := filepath.Base(file)

	timestamp := time.Now().Format(
		"2006-01-02 15:04:05",
	)

	fmt.Printf(
		"|\033[90m%s\033[0m| \033[100m[INFO]\033[0m >> %s >> %s:%d\n\n",
		timestamp,
		message,
		filename,
		line,
	)
}

func (l *Logger) Error(message string) {
	_, file, line, ok := runtime.Caller(1)

	if !ok {
		file = "unknown"
		line = 0
	}

	filename := filepath.Base(file)

	timestamp := time.Now().Format(
		"2006-01-02 15:04:05",
	)

	fmt.Printf(
		"|\033[90m%s\033[0m| \033[41m[ERROR]\033[0m >> %s >> %s:%d\n\n",
		timestamp,
		message,
		filename,
		line,
	)
	log.Fatal()

	// l.errorLogger.Println(message)
}
