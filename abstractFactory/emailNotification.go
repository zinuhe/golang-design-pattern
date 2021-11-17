package abstractFactory

import (
  "fmt"
)

// Email notifications
type EmailNotification struct {
}

type EmailNotificationSender struct {
}


func (EmailNotification) SendNotification() {
  fmt.Println("Sending Email Notification")
}

func (EmailNotification) GetSender() ISender {
 return EmailNotificationSender{}
}


func (EmailNotificationSender) GetSenderMethod() string {
  return "Email"
}

func (EmailNotificationSender) GetSenderChannel() string {
  return "SES"
}
