package log

import (
	"fmt"
	"io"
	"os"
	"path"
	"runtime"
	"sync"
	"time"
)

// mutex mutex for output
var mutex sync.Mutex

// DefaultLogger is a default logger
var DefaultLogger = &Logger{
	Name:      "default",
	level:     DEBUG,
	WithColor: false,
	Disable:   false,
	Outputer:  os.Stderr,
}

// SetLevel 设置默认Logger的日志级别
func SetLevel(level uint8) {
	mutex.Lock()
	defer mutex.Unlock()

	DefaultLogger.level = level
}

// EnableColor 设置默认Logger的颜色
func EnableColor(withColor bool) {
	mutex.Lock()
	defer mutex.Unlock()

	DefaultLogger.WithColor = withColor
}

// level
const (
	DEBUG uint8 = iota
	INFO  uint8 = iota
	WARN  uint8 = iota
	ERROR uint8 = iota
	FATAL uint8 = iota
)

// Logger .
type Logger struct {
	Name      string
	level     uint8
	WithColor bool
	Disable   bool
	Outputer  io.Writer
}

// GetLogger 获取新Logger
func GetLogger(name string) *Logger {
	return &Logger{
		Name:      name,
		WithColor: false,
		Disable:   false,
		Outputer:  os.Stderr,
	}
}

func (logger *Logger) SetOutput(w io.Writer) {
	logger.Outputer = w
}

func (logger *Logger) Output() io.Writer {
	return logger.Outputer
}

func (logger *Logger) SetLevel(level uint8) {
	logger.level = level
}

func (logger *Logger) Level() uint8 {
	return logger.level
}

func (logger *Logger) Prefix() string {
	return ""
}

func (logger *Logger) SetPrefix(string) {

}

func (logger *Logger) EnableColor(colored bool) {
	logger.WithColor = colored
}

func (logger *Logger) Disabled(disable bool) {
	logger.Disable = disable
}

func (logger *Logger) Header(time string, level uint8, filepath string, line int) string {
	levelName := fmt.Sprintf("%s", levelNames[level])
	levelColor := levelColors[level]

	if logger.WithColor {
		levelName = Colored(levelColor, levelName)
	}

	return fmt.Sprintf("%s [%s][%s:%d]", time, levelName, filepath, line)
}

func (logger *Logger) Log(level uint8, msg string) error {
	if !logger.Disable && level >= logger.level {
		_, filename, line, _ := runtime.Caller(2)
		pkgName := path.Base(path.Dir(filename))
		filepath := path.Join(pkgName, path.Base(filename))
		now := time.Now().Format("2006/01/02 15:04:05")

		header := logger.Header(now, level, filepath, line)
		mutex.Lock()
		defer mutex.Unlock()
		_, err := fmt.Fprintf(logger.Outputer, "%s %s\n", header, msg)
		return err
	}
	return nil
}

// Print 普通
func (logger *Logger) Print(a ...interface{}) {
	logger.Log(INFO, fmt.Sprint(a...))
}

// Printf 普通
func (logger *Logger) Printf(format string, a ...interface{}) {
	logger.Log(INFO, fmt.Sprintf(format, a...))
}

// Debug 调试
func (logger *Logger) Debug(a ...interface{}) {
	logger.Log(DEBUG, fmt.Sprint(a...))
}

// Debugf 调试
func (logger *Logger) Debugf(format string, a ...interface{}) {
	logger.Log(DEBUG, fmt.Sprintf(format, a...))
}

// Info 普通
func (logger *Logger) Info(a ...interface{}) {
	logger.Log(INFO, fmt.Sprint(a...))
}

// Infof 普通
func (logger *Logger) Infof(format string, a ...interface{}) {
	logger.Log(INFO, fmt.Sprintf(format, a...))
}

// Warn 警告
func (logger *Logger) Warn(a ...interface{}) {
	logger.Log(WARN, fmt.Sprint(a...))
}

// Warnf 警告
func (logger *Logger) Warnf(format string, a ...interface{}) {
	logger.Log(WARN, fmt.Sprintf(format, a...))
}

// Error 错误
func (logger *Logger) Error(a ...interface{}) {
	logger.Log(ERROR, fmt.Sprint(a...))
}

// Errorf 错误
func (logger *Logger) Errorf(format string, a ...interface{}) {
	logger.Log(ERROR, fmt.Sprintf(format, a...))
}

// Fatal 错误并退出
func (logger *Logger) Fatal(a ...interface{}) {
	logger.Log(FATAL, fmt.Sprint(a...))
	os.Exit(1)
}

// Fatalf 错误并退出
func (logger *Logger) Fatalf(format string, a ...interface{}) {
	logger.Log(FATAL, fmt.Sprintf(format, a...))
	os.Exit(1)
}

func (logger *Logger) Panic(...interface{})          {}
func (logger *Logger) Panicf(string, ...interface{}) {}

var levelNames = [4]string{"DEBUG", "INFO", "WARN", "ERROR"}

var colors = map[string]int{
	"black":   0,
	"red":     1,
	"green":   2,
	"yellow":  3,
	"blue":    4,
	"magenta": 5,
	"cyan":    6,
	"white":   7,
}
var levelColors = map[uint8]string{
	DEBUG: "magenta",
	INFO:  "green",
	WARN:  "yellow",
	ERROR: "red",
	FATAL: "magenta",
}

// Colored 为字符串加上颜色
func Colored(color string, text string) string {
	return fmt.Sprintf("\033[3%dm%s\033[0m", colors[color], text)
}

// Debugf 调试
func Debugf(format string, a ...interface{}) error {
	return DefaultLogger.Log(DEBUG, fmt.Sprintf(format, a...))
}

// Infof 普通
func Infof(format string, a ...interface{}) error {
	return DefaultLogger.Log(INFO, fmt.Sprintf(format, a...))
}

// Warnf 警告
func Warnf(format string, a ...interface{}) error {
	return DefaultLogger.Log(WARN, fmt.Sprintf(format, a...))
}

// Errorf 错误
func Errorf(format string, a ...interface{}) error {
	return DefaultLogger.Log(ERROR, fmt.Sprintf(format, a...))
}

// Fatalf 错误并退出
func Fatalf(format string, a ...interface{}) {
	DefaultLogger.Log(ERROR, fmt.Sprintf(format, a...))
	os.Exit(0)
}

// Debug 调试
func Debug(a ...interface{}) error {
	return DefaultLogger.Log(DEBUG, fmt.Sprint(a...))
}

// Info 普通
func Info(a ...interface{}) error {
	return DefaultLogger.Log(INFO, fmt.Sprint(a...))
}

// Warn 普通
func Warn(a ...interface{}) error {
	return DefaultLogger.Log(WARN, fmt.Sprint(a...))
}

// Error 错误
func Error(a ...interface{}) error {
	return DefaultLogger.Log(ERROR, fmt.Sprint(a...))
}

// Fatal 错误并退出
func Fatal(a ...interface{}) {
	DefaultLogger.Log(ERROR, fmt.Sprint(a...))
	os.Exit(0)
}
