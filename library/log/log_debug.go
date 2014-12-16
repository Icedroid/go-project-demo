// +build log-debug
package log

func (logger *Logger) Debug(format string, a ...interface{}) {
	logger.doPrintf(debugLevel, printDebugLevel, format, a...)
}

func Debug(format string, a ...interface{}) {
	gLogger.Debug(format, a...)
}
