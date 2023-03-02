package common

import (
	"fmt"
	"io"
	"log"
	"os"
	"runtime"
	"time"
)

const totalStep = 5

const (
	ERROR = iota + 1
	WARNING
	INFO
	DEBUG
)

// https://zenn.dev/tharu/articles/8c2ec139615fc4
const (
	LOG_DIR        = "log/"
	LOG_EXT        = ".log"
	INFO_LOG_NAME  = "info_"
	ERROR_LOG_NAME = "error_"
	INFO_PREFIX    = "[INFO]"
	ERROR_PREFIX   = "[ERROR]"
	LOG_FORMAT     = "2006-01-02"
)

func SetLogLevel() int {
	switch GetServeConf().LogLevel {
	case "INFO", "info":
		return INFO
	case "DEBUG", "debug":
		return DEBUG
	case "ERROR", "error":
		return ERROR
	case "WARNING", "warning":
		return WARNING
	default:
		return INFO
	}
}

type Log struct {
	logger *log.Logger
	level  int
	step   int
	file   io.Writer
}

var infoLog *Log

func GetLog() *Log {
	return infoLog
}

func GetLogger() *log.Logger {
	return infoLog.logger
}

func NewLog() *Log {
	logfile, _ := os.OpenFile(getLogFileName(INFO_LOG_NAME), // ファイル名
		os.O_RDWR|os.O_CREATE|os.O_APPEND, // 開く際のモード(読み書きOK)
		0666)                              // パーミッション0666

	infoLog = &Log{
		logger: log.Default(),
		level:  SetLogLevel(),
		step:   1,
		file:   io.MultiWriter(os.Stdout, logfile),
	}
	return infoLog
}

func (l *Log) NextStep() {
	l.step = l.step + 1
}

func (l *Log) Debug(format string, args ...interface{}) {
	if l.level >= DEBUG {
		prefix := fmt.Sprintf("[%s] [Step %d/%d] ", "DEBG", l.step, totalStep)
		l.logger.SetOutput(l.file)
		l.logger.SetPrefix(prefix)
		l.logger.SetFlags(log.Ldate | log.Ltime)

		_, file, line, ok := runtime.Caller(1)
		if ok {
			caller := fmt.Sprintf("@%s:%d: ", file, line)
			l.logger.Printf(caller+format, args...)
		} else {
			l.logger.Printf(format, args...)
		}
	}
}

func (l *Log) Info(format string, args ...interface{}) {
	if l.level >= INFO {
		prefix := fmt.Sprintf("[%s] [Step %d/%d] ", "INFO", l.step, totalStep)
		l.logger.SetOutput(l.file)
		l.logger.SetPrefix(prefix)
		l.logger.SetFlags(log.Ldate | log.Ltime)
		l.logger.Printf(format, args...)
	}
}

func (l *Log) Error(format string, args ...interface{}) {
	if l.level >= ERROR {

		logfile, _ := os.OpenFile(getLogFileName(ERROR_LOG_NAME), // ファイル名
			os.O_RDWR|os.O_CREATE|os.O_APPEND, // 開く際のモード(読み書きOK)
			0666)                              // パーミッション0666
		// コンソールへのlog出力時にファイルにも出力する
		multiLogFile := io.MultiWriter(os.Stdout, logfile)

		prefix := fmt.Sprintf("[%s] [Step %d/%d] ", "EROR", l.step, totalStep)
		l.logger.SetOutput(multiLogFile)
		// l.logger.SetOutput(os.Stdout)
		l.logger.SetPrefix(prefix)
		l.logger.SetFlags(log.Ldate | log.Ltime)

		_, file, line, ok := runtime.Caller(1)
		if ok {
			caller := fmt.Sprintf("@%s:%d: ", file, line)
			l.logger.Printf(caller+format, args...)
		} else {
			l.logger.Printf(format, args...)
		}
	}
}

func (l *Log) Fatal(format string, args ...interface{}) {
	if l.level >= ERROR {
		prefix := fmt.Sprintf("[%s] [Step %d/%d] ", "EROR", l.step, totalStep)
		l.logger.SetOutput(l.file)
		l.logger.SetPrefix(prefix)
		l.logger.SetFlags(log.Ldate | log.Ltime)

		_, file, line, ok := runtime.Caller(1)
		if ok {
			caller := fmt.Sprintf("@%s:%d: ", file, line)
			l.logger.Fatalf(caller+format, args...)
		} else {
			l.logger.Fatalf(format, args...)
		}
	}
}

/*
名称：INFOログファイル名生成処理
概要：INFOログのファイル名を生成する
param : 無し
return : ファイル名
*/
func getLogFileName(fileName string) string {
	// ファイル名を生成(現在日付)
	return LOG_DIR + fileName + time.Now().Format(LOG_FORMAT) + LOG_EXT
}
