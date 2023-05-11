package logger


import (
	"log"

)

//constructor
func newLogger() *LoggerLocal  {
	return &LoggerLocal{}
}

type LoggerLocal struct{}

func (l * LoggerLocal) Log(msg string)  {
	log.Print(msg)
}