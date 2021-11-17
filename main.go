package main

import (
  "fmt"
  af "github.com/zinuhe/golang-design-pattern/abstractFactory"
)

func main(){
  fmt.Println("Golang Design Patterns")

  fmt.Println("\nAbstract Factory Design Pattern")

  smsFactory, _   := af.GetNotificationFactory(af.NotificationSMS)
  emailFactory, _ := af.GetNotificationFactory(af.NotificationEmail)

  af.SendNotification(smsFactory)
  af.SendNotification(emailFactory)

  af.GetMethod(smsFactory)
  af.GetMethod(emailFactory)

}
