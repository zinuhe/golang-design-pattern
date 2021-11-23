package abstractFactory

import (
  "fmt"
)

const (
	NotificationEmail = "Email"
	NotificationSMS   = "SMS"
)

// To manage notifications by Email and SMS
type INotificationFactory interface {
  SendNotification()
  GetSender() ISender
}

type ISender interface {
  GetSenderMethod() string
  GetSenderChannel() string
}


// Factory
func GetNotificationFactory (notificationType string) (INotificationFactory, error) {
  switch notificationType {
  case NotificationEmail:
    return &EmailNotification{}, nil
  case NotificationSMS:
    return &SMSNotification{}, nil
  default:
    return nil, fmt.Errorf("No Notification Type")
  }
}

func SendNotification(f INotificationFactory) {
  f.SendNotification()
}

func GetMethod(f INotificationFactory) {
  fmt.Println(f.GetSender().GetSenderMethod())
}
