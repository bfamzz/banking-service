package mail

type EmailSender interface {
	SendEmail(subject string, content string, to []string, cc []string,
		bcc []string, attachments []string,
	) error
}

type GmailSender struct {
	name string
	fromEmailAddress string
	fromEmailPassword string
}

func NewGmailSender(name string, fromEmailAddress string, fromEmailPassword string) EmailSender {
	return &GmailSender{
		name: name,
		fromEmailAddress: fromEmailAddress,
		fromEmailPassword: fromEmailPassword,
	}
}

func (emailSender *GmailSender) SendEmail(subject string, content string, to []string, cc []string,
	bcc []string, attachments []string) error {
	
	return nil
}