package logger

func newLoggerDummy() *LoggerDummy {
	return &LoggerDummy{}
}

type LoggerDummy struct{}

func (l *LoggerDummy) Log(msg string)  {
	
}