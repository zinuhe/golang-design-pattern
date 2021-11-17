package abstractFactory

import (
  "fmt"
)

// SMS Notification
type SMSNotification struct {
}

type SMSNotificationSender struct {
}


func (SMSNotification) SendNotification() {
  fmt.Println("Sending SMS Notification")
}

func (SMSNotification) GetSender() ISender {
  return SMSNotificationSender{}
}


func (SMSNotificationSender) GetSenderMethod() string {
  return "SMS"
}

func (SMSNotificationSender) GetSenderChannel() string {
  return "Twilio"
}
