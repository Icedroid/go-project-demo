// +build !log-debug
package log

func (logger *Logger) Debug(format string, a ...interface{}) {
}

func Debug(format string, a ...interface{}) {
}
