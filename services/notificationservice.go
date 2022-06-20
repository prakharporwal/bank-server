package services

type NotificatonService interface {
	CreateNotification()
	ListNotifications()
	SendEmail()
	SendMessage()
}

func ListNotifications() {
	// List all Notification
}

func CreateNotification() {
	// create a notification using firebase or AWS SNS
}

func SendEmail() {
	// send a email using AWS SES
}

func SendMessage() {
	// send a message to mobile using SNS or other service
}
