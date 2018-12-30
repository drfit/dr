package logus

var (
	// InDebug Whether Dr. is in debug mode
	InDebug bool

	logger Logger
)

// Register replace a new logger to print log msg
func Register(l Logger) {
	if l != nil {
		logger = l
	}
}

// Info print Info level log msg
func Info(msg string) {
	logger.Info(msg)
}

// Debug print Debug level log msg
func Debug(msg string) {
	logger.Debug(msg)
}

// Warn print Warn level log msg
func Warn(msg string) {
	logger.Warn(msg)
}

// Error print Error level log msg
func Error(msg string) {
	logger.Error(msg)
}

// Panic print Err level log msg
func Panic(msg string) {
	logger.Panic(msg)
}
