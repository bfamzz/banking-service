package mail

import (
	"context"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/service/sesv2"
	"github.com/aws/aws-sdk-go-v2/service/sesv2/types"
)

const (
	WelcomeTemplateName      = "WelcomeTemplate"
	VerificationTemplateName = "VerificationTemplate"
)

type SesSender struct {
	fromEmailAddress        string
	fromEmailNameAndAddress string
	client                  *sesv2.Client
}

func NewSesSender(sdkConfig aws.Config, fromEmailAddress string, fromEmailNameAndAddress string) EmailSender {
	return &SesSender{
		fromEmailAddress:        fromEmailAddress,
		fromEmailNameAndAddress: fromEmailNameAndAddress,
		client:                  sesv2.NewFromConfig(sdkConfig),
	}
}

func (emailSender *SesSender) SendTemplateEmail(templateName string, templateData string,
	to []string, cc []string, bcc []string, attachments []string) error {
	_, err := emailSender.client.SendEmail(context.TODO(), &sesv2.SendEmailInput{
		Content: &types.EmailContent{
			Template: &types.Template{
				TemplateName: &templateName,
				TemplateData: &templateData,
			},
		},
		FromEmailAddress: &emailSender.fromEmailNameAndAddress,
		Destination: &types.Destination{
			ToAddresses:  to,
			CcAddresses:  cc,
			BccAddresses: bcc,
		},
		ReplyToAddresses: []string{emailSender.fromEmailAddress},
	})
	return err
}
