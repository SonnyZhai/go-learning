package main

import (
	"fmt"
	"log"
	"os"
	"time"
)

// Logger 接口定义了日志操作的基本方法
type Logger interface {
	Log(message string)
}

// ConsoleLogger 结构体用于在终端输出日志
type ConsoleLogger struct{}

// Log 实现 Logger 接口，向终端写日志
func (c ConsoleLogger) Log(message string) {
	fmt.Println(time.Now().Format("2006-01-02 15:04:05"), message)
}

// FileLogger 结构体用于向文件写日志
type FileLogger struct {
	File *os.File
}

// Log 实现 Logger 接口，向文件写日志
func (f FileLogger) Log(message string) {
	log.SetOutput(f.File)
	log.Println(time.Now().Format("2006-01-02 15:04:05"), message)
}

// NewFileLogger 用于创建一个 FileLogger，并打开指定的日志文件
func NewFileLogger(filePath string) (*FileLogger, error) {
	file, err := os.OpenFile(filePath, os.O_CREATE|os.O_WRONLY|os.O_APPEND, 0666) // 打开日志文件
	if err != nil {
		return nil, err
	}
	return &FileLogger{File: file}, nil
}

func SimpleLog() {
	// 创建一个终端日志记录器
	consoleLogger := ConsoleLogger{}

	// 使用终端日志记录器写日志
	consoleLogger.Log("This is a log message to the console")

	// 创建一个文件日志记录器
	fileLogger, err := NewFileLogger("logFile.txt")
	if err != nil {
		fmt.Println("Error creating file logger:", err)
		return
	}
	defer fileLogger.File.Close()

	// 使用文件日志记录器写日志
	fileLogger.Log("This is a log message to the file")

	// 可以灵活地使用不同的日志记录器
	var logger Logger

	logger = consoleLogger
	logger.Log("Logging to the console again")

	logger = fileLogger
	logger.Log("Logging to the file again")

}
