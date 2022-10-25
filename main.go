package main

import (
	"fmt"
	"github.com/aws/aws-sdk-go/aws/session"
	"send-message-service/service"
	"strings"
)

func main() {
	sess := session.Must(session.NewSessionWithOptions(session.Options{
		SharedConfigState: session.SharedConfigEnable,
	}))

	queues, err := service.ListQueues(sess)
	if err != nil {
		fmt.Println("Got an error retrieving queue URLs:")
	}

	var queueURL string
	for _, url := range queues.QueueUrls {
		if strings.Contains(*url, "sqs-test.fifo") {
			queueURL = *url
		}
	}

	err = service.SendMessage(sess, &queueURL)
	if err != nil {
		fmt.Println("메시지 전송 에러")
	}
}
