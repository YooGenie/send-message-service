package send

import (
	"fmt"
	"os"

	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/sns"
	"github.com/google/uuid"
)

func SendSNSMessage(topicArn string, sess *session.Session) error {
	svc := sns.New(sess)

	msgPtr := "메시지 내용입니다."
	messageGroupId := "test"
	messageDeduplicationId := uuid.New().String()

	result, err := svc.Publish(&sns.PublishInput{
		Message:                &msgPtr,
		TopicArn:               &topicArn,
		MessageGroupId:         &messageGroupId,
		MessageDeduplicationId: &messageDeduplicationId,
	})
	if err != nil {
		fmt.Println(err.Error())
		os.Exit(1)
	}

	fmt.Println(*result.MessageId)
	return err
}
