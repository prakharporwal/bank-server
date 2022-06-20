package services

import (
	"log"
)

type Logger struct {
	Flag string
}

func (l *Logger) Info(message string, args ...interface{}) {
	log.Println("INFO :", message, args)
}

func (l *Logger) Warn(message string, args ...interface{}) {
	log.Println("WARNING : ", message, args)
}

func (l *Logger) Error(message string, args ...interface{}) {
	log.Println("ERROR : ", message, args)
}

func (l *Logger) Debug(message string, args ...interface{}) {
	log.Println("DEBUG : ", message, args)
}
