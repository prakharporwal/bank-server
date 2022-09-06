package klog

import (
	"log"
)

type KLogger interface {
	Info()
	Warn()
	Error()
	Debug()
}

func Info(message string, args interface{}) {
	log.Println("INFO :", message, args)
}

func Warn(message string, args ...interface{}) {
	log.Println("WARNING : ", message, args)
}

func Error(message string, args ...interface{}) {
	log.Println("ERROR : ", message, args)
}

func Debug(message string, args ...interface{}) {
	log.Println("DEBUG : ", message, args)
}
