package mail

type EmailSender interface {
	SendTemplateEmail(templateName string, templateData string, to []string,
		cc []string, bcc []string, attachments []string) error
}
