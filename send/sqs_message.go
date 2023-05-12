package send

import (
	"encoding/json"
	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/sqs"
	"github.com/google/uuid"
	"send-message-service/config"
)

func SendSqsMessage(sess *session.Session, queueURL *string, eventType string, content interface{}) error {
	svc := sqs.New(sess)

	messageGroupId := eventType
	messageDeduplicationId := uuid.New().String()

	messageBody, _ := json.Marshal(content)

	var _, err = svc.SendMessage(&sqs.SendMessageInput{
		MessageAttributes: map[string]*sqs.MessageAttributeValue{
			"EventType": {
				DataType:    aws.String("String"),
				StringValue: aws.String(eventType),
			},
			"Env": {
				DataType:    aws.String("String"),
				StringValue: aws.String(config.Config.Environment),
			},
			"Module": {
				DataType:    aws.String("String"),
				StringValue: aws.String(config.Config.Module),
			},
		},
		MessageBody:            aws.String(string(messageBody)),
		QueueUrl:               queueURL,
		MessageGroupId:         &messageGroupId,
		MessageDeduplicationId: &messageDeduplicationId,
	})

	if err != nil {
		return err
	}

	return nil
}
