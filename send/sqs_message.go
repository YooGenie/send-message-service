package send

import (
	"fmt"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/sqs"
)

func SendSqsMessage(sess *session.Session, queueURL *string) error {
	svc := sqs.New(sess)
	fmt.Println(svc)
	return nil
}
