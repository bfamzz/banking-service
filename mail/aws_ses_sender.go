package mail

import (
	"context"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/service/sesv2"
	"github.com/aws/aws-sdk-go-v2/service/sesv2/types"
	"github.com/bfamzz/banking-service/util"
)

const (
	welcomeTemplateName = "WelcomeTemplate"
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

func (emailSender *SesSender) SendTemplateEmail(templateData string, to []string, cc []string, bcc []string, attachments []string) error {
	_, err := emailSender.client.SendEmail(context.TODO(), &sesv2.SendEmailInput{
		Content: &types.EmailContent{
			// Simple: &types.Message{
			// 	Body: &types.Body{
			// 		Html: &types.Content{
			// 			Data:    util.StringPointer(charset),
			// 			Charset: util.StringPointer(charset),
			// 		},
			// 		Text: &types.Content{
			// 			Data:    util.StringPointer(charset),
			// 			Charset: util.StringPointer(charset),
			// 		},
			// 	},
			// 	Subject: &types.Content{
			// 		Data:    &subject,
			// 		Charset: util.StringPointer(charset),
			// 	},
			// },
			Template: &types.Template{
				TemplateName: util.StringPointer(welcomeTemplateName),
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
