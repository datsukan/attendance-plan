package repository

import (
	"context"
	"fmt"
	"net/mail"

	"github.com/aws/aws-sdk-go-v2/service/sesv2"
	"github.com/aws/aws-sdk-go-v2/service/sesv2/types"
	"github.com/aws/aws-sdk-go/aws"
	"github.com/datsukan/attendance-plan/backend/infrastructure"
)

// EmailRepository はメールの repository を表すインターフェースです。
type EmailRepository interface {
	Send(ctx context.Context, to, subject, body string) (string, error)
}

// EmailRepositoryImpl はメールの repository の実装を表す構造体です。
type EmailRepositoryImpl struct {
	Client      infrastructure.MailClient
	SenderEmail string
	SenderName  string
}

// NewEmailRepository は EmailRepository を生成します。
func NewEmailRepository(client infrastructure.MailClient, senderEmail, senderName string) EmailRepository {
	return &EmailRepositoryImpl{
		Client:      client,
		SenderEmail: senderEmail,
		SenderName:  senderName,
	}
}

// Send はメールを送信します。
func (r *EmailRepositoryImpl) Send(ctx context.Context, to, subject, body string) (string, error) {
	if r.Client == nil {
		return "", fmt.Errorf("client is nil")
	}

	if r.SenderEmail == "" {
		return "", fmt.Errorf("sender email is empty")
	}

	if to == "" {
		return "", fmt.Errorf("to is empty")
	}

	if subject == "" {
		return "", fmt.Errorf("subject is empty")
	}

	if body == "" {
		return "", fmt.Errorf("body is empty")
	}

	address := mail.Address{
		Name:    r.SenderName,
		Address: r.SenderEmail,
	}

	input := &sesv2.SendEmailInput{
		FromEmailAddress: aws.String(address.String()), // 送信元
		Destination: &types.Destination{
			ToAddresses: []string{to}, // 送信先
		},
		Content: &types.EmailContent{
			Simple: &types.Message{
				Body: &types.Body{
					Text: &types.Content{
						Data: aws.String(body), // 本文
					},
				},
				Subject: &types.Content{
					Data: aws.String(subject), // 件名
				},
			},
		},
	}

	res, err := r.Client.SendEmail(ctx, input)
	if err != nil {
		return "", err
	}

	if res == nil || res.MessageId == nil {
		return "", fmt.Errorf("failed to send email")
	}

	return *res.MessageId, nil
}
