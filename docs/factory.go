type Notifier interface {
	Send(msg string)
}

type EmailNotifier struct{}
func (e *EmailNotifier) Send(msg string) {
 fmt.Println("Email:", msg)
}
type SMSNotifier struct{}
func (s *SMSNotifier) Send(msg string) {
 fmt.Println("SMS:", msg)
}
func GetNotifier(channel string) Notifier {
 switch channel {
 case "email":
  return &EmailNotifier{}
 case "sms":
  return &SMSNotifier{}
 default:
  return nil
 }
}
func main() {
 notifier := GetNotifier("email")
 notifier.Send("Hello from factory pattern")
}