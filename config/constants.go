package config

import "time"

var CurrDbDateTime = time.Now().Format("2006-01-02 15:04:05")
var FormatDbDateTime = "02-01-2006 3:4:5 pm"
var FormatDbDate = "2006-01-02"
var FormatDbTime = "15:04:05"
var CheckInTime = " 12:00:00.000"
var CheckOutTime = " 11:00:00.000"
var ImageFolder = "file"
var ImageFolderUser = "file/user"

// log
const (
	LOG_DIRECTORY    = "./logs"
	SERVER_ERROR     = "something went wrong on the server side"
	INFO_LOG_FILE    = "infoLog"
	ERROR_LOG_FILE   = "errorLog"
	WARNING_LOG_FILE = "warningLog"
	LOGFILE_DATE     = "no"
)
