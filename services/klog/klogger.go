package klog

import (
	"log"
)

func init() {
	log.SetFlags(log.Ltime)
}

func Info(args ...interface{}) {
	log.SetPrefix("INFO: ")
	log.Println(args...)
}

func Warn(args ...interface{}) {
	log.SetPrefix("WARNING: ")
	log.Println(args...)
}

func Error(args ...interface{}) {
	log.SetPrefix("ERROR: ")
	log.Println(args...)
}

func Debug(args ...interface{}) {
	log.SetPrefix("DEBUG: ")
	log.Println(args...)
}
