package bridge

type MsgSender interface {
	Send(msg string) error
}

//短信通知，可能还有邮件通知，微信通知。。。
type SmsSender struct {
	telephones []string
}

func NewSmsSender(telephones []string) *SmsSender {
	return &SmsSender{telephones: telephones}
}

func (s *SmsSender) Send(msg string) error {
	return nil
}

type Notification interface {
	Notify(msg string)
}

//通知的紧急程度SEVERE（严重）、URGENCY（紧急）、NORMAL（普通）、TRIVIAL（无关紧要）
type SevereNotification struct {
	sender MsgSender
}

func NewSevereNotification(sender MsgSender) *SevereNotification {
	return &SevereNotification{sender: sender}
}

func (s SevereNotification) Notify(msg string) error {
	return s.sender.Send(msg)
}
