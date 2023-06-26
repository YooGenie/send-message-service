package send

import (
	"encoding/json"
	"send-message-service/config"

	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/sqs"
	"github.com/google/uuid"
)

func SendSqsMessage(sess *session.Session, queueURL *string, eventType string, content interface{}) error {
	svc := sqs.New(sess)

	messageGroupId := eventType
	messageDeduplicationId := uuid.New().String()

	messageBody, _ := json.Marshal(content)
	dataType := "String"

	var _, err = svc.SendMessage(&sqs.SendMessageInput{
		MessageAttributes: map[string]*sqs.MessageAttributeValue{
			"EventType": {
				DataType:    &dataType,
				StringValue: &eventType,
			},
			"Env": {
				DataType:    &dataType,
				StringValue: &config.Config.Environment,
			},
			"Module": {
				DataType:    &dataType,
				StringValue: &config.Config.Module,
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
