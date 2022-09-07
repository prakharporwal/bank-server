package utils

import (
	"github.com/google/uuid"
	"github.com/prakharporwal/bank-server/services/klog"
	"net/mail"
	"time"
)

func IsValidEmail(email string) bool {
	// validating using internal go mail lib
	_, err := mail.ParseAddress(email)
	if err != nil {
		klog.Error("Couldn't validate email address!")
		return false
	}
	return true
}

func GenerateTimeStampMicro() int64 {
	return time.Now().UnixMicro()
}

func GenerateTimeStampMilli() int64 {
	return time.Now().UnixMilli()
}

func GenerateUUID() uuid.UUID {
	uid, _ := uuid.NewRandom()
	return uid
}
