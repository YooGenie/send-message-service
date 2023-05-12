package main

import (
	"github.com/aws/aws-sdk-go/aws/session"
	"log"
	"send-message-service/config"
	"send-message-service/send"
)

func main() {
	config.ConfigureEnvironment("./")

	sess := session.Must(session.NewSessionWithOptions(session.Options{
		SharedConfigState: session.SharedConfigEnable,
	}))

	queueURL := ""

	err := send.SendSqsMessage(sess, &queueURL, "emil", "메시지 내용")
	if err != nil {
		log.Println("메시지 전송 에러")
	}
}
