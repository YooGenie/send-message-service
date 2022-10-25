package service

import (
	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/sqs"
)

func SendMessage(sess *session.Session, queueURL *string) error {
	svc := sqs.New(sess)

	_, err := svc.SendMessage(&sqs.SendMessageInput{
		MessageGroupId: aws.String("donation"),
		MessageBody:    aws.String("테스트입니다."),
		QueueUrl:       queueURL,
	})

	if err != nil {
		return err
	}

	return nil
}
