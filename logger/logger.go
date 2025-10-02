package logger

import (
	"fmt"
	"io"
	"log"
	"os"
	"sync"
	"time"
)

var (
	logFile    *os.File
	logger     *log.Logger
	currentDay int
	mu         sync.Mutex
	logDir     = "logs" // 日志目录
)

func init() {
	err := os.MkdirAll(logDir, 0755)
	if err != nil {
		log.Fatalf("Failed to create log directory: %v", err)
	}
	initLogger()
}

func initLogger() {
	now := time.Now()
	fileName := fmt.Sprintf("%s/app_%s.log", logDir, now.Format("2006-01-02"))

	file, err := os.OpenFile(fileName, os.O_CREATE|os.O_WRONLY|os.O_APPEND, 0666)
	if err != nil {
		log.Fatalf("Failed to open log file: %v", err)
	}

	// 多写器：同时输出到控制台和文件
	multiWriter := io.MultiWriter(os.Stdout, file)

	logger = log.New(multiWriter, "", log.LstdFlags|log.Lshortfile)
	currentDay = now.Day()
	logFile = file
}

func checkDate() {
	now := time.Now()
	if now.Day() != currentDay {
		mu.Lock()
		defer mu.Unlock()

		// 再次检查，防止多个goroutine同时进入
		if now.Day() != currentDay {
			logFile.Close()
			initLogger()
		}
	}
}

func Printf(format string, v ...interface{}) {
	checkDate()
	logger.Output(2, fmt.Sprintf(format, v...))
}

func Println(v ...interface{}) {
	checkDate()
	logger.Output(2, fmt.Sprintln(v...))
}

func Fatalf(format string, v ...interface{}) {
	checkDate()
	logger.Output(2, fmt.Sprintf(format, v...))
	os.Exit(1)
}

func Fatalln(v ...interface{}) {
	checkDate()
	logger.Output(2, fmt.Sprintln(v...))
	os.Exit(1)
}

func GetLogger() *log.Logger {
	return logger
}
