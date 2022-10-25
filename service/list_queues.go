package service

import (
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/sqs"
)

func ListQueues(sess *session.Session) (*sqs.ListQueuesOutput, error) {

	svc := sqs.New(sess)

	result, err := svc.ListQueues(nil)
	if err != nil {
		return nil, err
	}

	return result, nil
}
