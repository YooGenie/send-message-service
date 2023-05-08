package main

import (
	"github.com/aws/aws-sdk-go/aws/session"
	"log"
	"send-message-service/send"
)

func main() {
	sess := session.Must(session.NewSessionWithOptions(session.Options{
		SharedConfigState: session.SharedConfigEnable,
	}))

	queueURL := ""

	err := send.SendSqsMessage(sess, &queueURL)
	if err != nil {
		log.Println("메시지 전송 에러")
	}
}
