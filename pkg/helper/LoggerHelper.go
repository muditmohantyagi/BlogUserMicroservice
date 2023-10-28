package helper

import (
	"net/http"
	"os"
	"time"

	"blog.com/config"
	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
)

/*

   methods named after the log level for log.Print-style logging
   methods ending in "w" for loosely-typed structured logging
   methods ending in "f" for log.Printf-style logging
   methods ending in "ln" for log.Println-style logging

   For example, the methods for InfoLevel are:

Info(...any)           Print-style logging
Infow(...any)          Structured logging (read as "info with")
Infof(string, ...any)  Printf-style logging
Infoln(...any)         Println-style logging
ex:ILog.Info("Success msg")
ex:ELog.Error(err.Error())
ex:WLog.Warn("Error msg")

ex:ELog.Info("Error msg")
*/
//Info log
var ILog = CustomStructuredLogs(config.INFO_LOG_FILE)

// Error log
var ELog = CustomStructuredLogs(config.ERROR_LOG_FILE)

// Warning log
var WLog = CustomStructuredLogs(config.WARNING_LOG_FILE)

func CustomStructuredLogs(logFileName string) *zap.SugaredLogger {
	var logfileLocation string
	if config.GetEnvWithKey("LOGFILE_DATE", "no") == "yes" {
		dirExist, err := IsDirectory(config.LOG_DIRECTORY)
		if err != nil && !dirExist {
			os.Mkdir(config.LOG_DIRECTORY, os.ModePerm)
			logfileLocation = config.GetEnvWithKey("LOGFILE", config.LOG_DIRECTORY) + "/" + logFileName + "_" + time.Now().Format("02-Jan-2006")
		} else {
			logfileLocation = config.GetEnvWithKey("LOGFILE", config.LOG_DIRECTORY) + "/" + logFileName + "_" + time.Now().Format("02-Jan-2006")
		}
	} else {
		dirExist, err := IsDirectory(config.LOG_DIRECTORY)
		if err != nil && !dirExist {
			os.Mkdir(config.LOG_DIRECTORY, os.ModePerm)
			logfileLocation = config.GetEnvWithKey("LOGFILE", config.LOG_DIRECTORY) + "/" + logFileName
		} else {
			logfileLocation = config.GetEnvWithKey("LOGFILE", config.LOG_DIRECTORY) + "/" + logFileName
		}
	}
	var cfg zap.Config
	cfg.OutputPaths = []string{"stdout", logfileLocation, "stderr"}
	cfg.Encoding = config.GetEnvWithKey("logfile_encoding", "json")
	cfg.Level = zap.NewAtomicLevelAt(zapcore.DebugLevel)
	cfg.EncoderConfig = zapcore.EncoderConfig{MessageKey: config.GetEnvWithKey("logfile_messageKey", "message"),
		TimeKey:      config.GetEnvWithKey("logfile_time", "time"),
		EncodeTime:   zapcore.ISO8601TimeEncoder,
		LevelKey:     config.GetEnvWithKey("logfile_level", "level"),
		EncodeLevel:  zapcore.CapitalLevelEncoder,
		CallerKey:    config.GetEnvWithKey("logfile_callerKey", "callerKey"),
		EncodeCaller: zapcore.ShortCallerEncoder}

	logger, err := cfg.Build()
	if err != nil {
		panic(err)
	}
	defer logger.Sync()
	log := logger.Sugar()
	defer log.Sync()
	return log
}
func Recover(c *gin.Context) {
	if r := recover(); r != nil {
		ELog.Error("exit crawl: PANIC occured")
		c.JSON(int(http.StatusInternalServerError), gin.H{
			"status": 500,
			"error":  config.SERVER_ERROR,
		})
	}
}

func IsDirectory(path string) (bool, error) {
	fileInfo, err := os.Stat(path)
	if err != nil {
		return false, err
	}
	return fileInfo.IsDir(), err
}
