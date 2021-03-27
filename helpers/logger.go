package helpers

// LogInfo logs info messages
func LogInfo(message interface{}) {
	Logger.Info(message)
}

// LogError logs error messages
func LogError(message interface{}) {
	Logger.Error(message)
}

// LogFatal logs fatal messages
func LogFatal(message interface{}) {
	Logger.Fatal(message)
}
