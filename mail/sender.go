package mail

type EmailSender interface {
	SendTemplateEmail(templateData string, to []string, cc []string, bcc []string, attachments []string) error
}
