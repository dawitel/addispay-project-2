package utils

import (
    "log"
    "os"
)

type Logger struct {
    *log.Logger
}

var logger *Logger

// Initialize the logger
func init() {
    logLevel := os.Getenv("LOG_LEVEL")
    if logLevel == "" {
        logLevel = "INFO" // Default log level
    }

    logger = NewLogger(logLevel)
}

// NewLogger creates a new logger instance with the specified log level
func NewLogger(level string) *Logger {
    flags := log.Ldate | log.Ltime | log.Lshortfile
    logger := log.New(os.Stdout, "", flags)

    return &Logger{logger}
}

// GetLogger returns a logger instance
func GetLogger() *Logger {
    return logger
}

// Info logs informational messages
func (l *Logger) Info(v ...interface{}) {
    l.SetPrefix("[INFO]: ")
    l.Println(v...)
}
// Success logs success messages
func (l *Logger) Success(v ...interface{}) {
    l.SetPrefix("☑ [SUCCESS]: ")
    l.Println(v...)
}

// Warn logs warning messages
func (l *Logger) Warn(v ...interface{}) {
    l.SetPrefix("⚠[WARNING]: ")
    l.Println(v...)
}

// Error logs error messages
func (l *Logger) Error(v ...interface{}) {
    l.SetPrefix("🔴[ERROR]: ")
    l.Println(v...)
}
//  writes log messages to service.log file  
func (l *Logger) LogToFile(v ...interface{}) {
    l.SetPrefix("[LOG]: ")
    l.Println(v...)
}
