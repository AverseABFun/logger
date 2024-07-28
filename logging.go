package logger

import (
	"fmt"
	"io"
	"log"
	"os"

	formatColor "github.com/fatih/color"
)

// Log type constants
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
	// LogNoPrefix sets the type prefix to nothing.
	LogNoPrefix
)

var logger = log.New(os.Stderr, "", log.Lmsgprefix|log.Ltime)

var flags = log.Lmsgprefix | log.Ltime

// Sets the logger's flags. Use the flags from the log package. log.Lmsgprefix is automatically set.
func SetFlags(newFlags int) {
	flags = log.Lmsgprefix | newFlags
	logger = log.New(os.Stderr, "", flags)
}

// Gets the logger's flags.
func GetFlags() int {
	return flags
}

// Sets the logger's stream.
func SetStream(stream io.Writer) {
	logger = log.New(stream, "", flags)
}

// Gets the logger's stream.
func GetStream() io.Writer {
	return logger.Writer()
}

var prefix = ""

// Log a message with a specified log type. Use the log type constants(LogInfo, LogError, LogWarning, LogDebug, and LogFatal) for the log type.
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
	case LogInfo:
		logger.SetPrefix(formatColor.WhiteString("[INFO] "))
	default:
		logger.SetPrefix("")
	}
	logger.SetPrefix(prefix + logger.Prefix())
	if logType == LogFatal {
		logger.Fatalln(msg)
	} else {
		logger.Println(msg)
	}
}

// Same as Log, except accepts a format string and format arguments.
func Logf(logType int, format string, a ...any) {
	Log(fmt.Sprintf(format, a...), logType)
}

// Logs a newline without any prefix(including the time, the custom prefix, and anything else from custom flags).
func NewlineWithoutPrefix() {
	logger.SetPrefix("")
	logger.SetFlags(0)
	logger.Print("\n")
	logger.SetFlags(flags)
}

// Sets the custom prefix. This is appended before the type prefix([ERROR], [WARNING], etc).
func SetPrefix(newPrefix string) {
	prefix = newPrefix
}

// Returns the custom prefix.
func GetPrefix() string {
	return prefix
}
