package main

import (
  "fmt"
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

// Factory
func getNotificationFactory (notificationType string) (INotificationFactory, error) {
  switch notificationType {
  case "SMS":
    return &SMSNotification{}, nil
  case "Email":
    return &EmailNotification{}, nil
  default:
    return nil, fmt.Errorf("No Notification Type")
  }
}

func sendNotification(f INotificationFactory) {
  f.SendNotification()
}

func getMethod(f INotificationFactory) {
  fmt.Println(f.GetSender().GetSenderMethod())
}

func main() {
  smsFactory, _   := getNotificationFactory("SMS")
  emailFactory, _ := getNotificationFactory("Email")

  sendNotification(smsFactory)
  sendNotification(emailFactory)

  getMethod(smsFactory)
  getMethod(emailFactory)

}
