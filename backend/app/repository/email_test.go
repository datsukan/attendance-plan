package repository

import (
	"context"
	"errors"
	"testing"

	"github.com/aws/aws-sdk-go-v2/service/sesv2"
	"github.com/datsukan/attendance-plan/backend/infrastructure"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

type stubMailClient struct{}

func (s *stubMailClient) SendEmail(ctx context.Context, params *sesv2.SendEmailInput, optFns ...func(*sesv2.Options)) (*sesv2.SendEmailOutput, error) {
	messageID := "test-message-id"
	return &sesv2.SendEmailOutput{
		MessageId: &messageID,
	}, nil
}

type stubErrorMailClient struct{}

func (s *stubErrorMailClient) SendEmail(ctx context.Context, params *sesv2.SendEmailInput, optFns ...func(*sesv2.Options)) (*sesv2.SendEmailOutput, error) {
	return nil, errors.New("send email error")
}

func TestEmail_Send(t *testing.T) {
	senderEmail := "test@example.com"
	to := "test-to@example.com"
	subject := "test subject"
	body := "test body"

	type fields struct {
		client      infrastructure.MailClient
		senderEmail string
	}
	type args struct {
		to      string
		subject string
		body    string
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		want    string
		wantErr error
	}{
		{
			name:   "正常系",
			fields: fields{&stubMailClient{}, senderEmail},
			args:   args{to, subject, body},
			want:   "test-message-id",
		},
		{
			name:    "異常系: client が nil",
			fields:  fields{nil, senderEmail},
			wantErr: errors.New("client is nil"),
		},
		{
			name:    "異常系: senderEmail が空",
			fields:  fields{&stubMailClient{}, ""},
			wantErr: errors.New("sender email is empty"),
		},
		{
			name:    "異常系: to が空",
			fields:  fields{&stubMailClient{}, senderEmail},
			args:    args{"", subject, body},
			wantErr: errors.New("to is empty"),
		},
		{
			name:    "異常系: subject が空",
			fields:  fields{&stubMailClient{}, senderEmail},
			args:    args{to, "", body},
			wantErr: errors.New("subject is empty"),
		},
		{
			name:    "異常系: body が空",
			fields:  fields{&stubMailClient{}, senderEmail},
			args:    args{to, subject, ""},
			wantErr: errors.New("body is empty"),
		},
		{
			name:    "異常系: SendEmail がエラー",
			fields:  fields{&stubErrorMailClient{}, senderEmail},
			args:    args{to, subject, body},
			wantErr: errors.New("send email error"),
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			assert := assert.New(t)
			require := require.New(t)

			ctx := context.Background()
			r := &EmailRepositoryImpl{
				Client:      tt.fields.client,
				SenderEmail: tt.fields.senderEmail,
			}
			got, err := r.Send(ctx, tt.args.to, tt.args.subject, tt.args.body)
			if tt.wantErr != nil {
				require.Error(err)
				assert.Equal(tt.wantErr, err)
				assert.Empty(got)
				return
			}

			require.NoError(err)
			assert.Equal(tt.want, got)
		})
	}
}
