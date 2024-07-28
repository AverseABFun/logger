package logger

import (
	"fmt"
	"log"
	"os"

	formatColor "github.com/fatih/color"
)

const (
	// LogInfo is used to log information messages
	LogInfo = iota
	// LogError is used to log error messages
	LogError
	// LogWarning is used to log warning messages
	LogWarning
	// LogDebug is used to log debug messages
	LogDebug
	// LogFatal is used log fatal messages and quit
	LogFatal
)

var logger = log.New(os.Stderr, "", log.Lmsgprefix|log.Ltime)

var prefix = ""

func Log(msg string, logType int) {
	msg = formatColor.MagentaString(msg)
	switch logType {
	case LogError:
		logger.SetPrefix(formatColor.RedString("[ERROR] "))
	case LogWarning:
		logger.SetPrefix(formatColor.YellowString("[WARNING] "))
	case LogDebug:
		logger.SetPrefix(formatColor.BlueString("[DEBUG] "))
	case LogFatal:
		logger.SetPrefix(formatColor.HiRedString("[FATAL] "))
	default:
		logger.SetPrefix(formatColor.WhiteString("[INFO] "))
	}
	logger.SetPrefix(prefix + logger.Prefix())
	if logType == LogFatal {
		logger.Fatalln(msg)
	} else {
		logger.Println(msg)
	}
}

func Logf(logType int, format string, a ...any) {
	Log(fmt.Sprintf(format, a...), logType)
}

func NewlineWithoutPrefix() {
	logger.SetPrefix("")
	logger.SetFlags(0)
	logger.Print("\n")
	logger.SetFlags(log.Lmsgprefix | log.Ltime)
}

func SetPrefix(newPrefix string) {
	prefix = newPrefix
}

func GetPrefix() string {
	return prefix
}
