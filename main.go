package main

import (
	"log"
	"send-message-service/config"
	"send-message-service/send"

	"github.com/aws/aws-sdk-go/aws/session"
)

func main() {
	config.ConfigureEnvironment("./")

	sess := session.Must(session.NewSessionWithOptions(session.Options{
		SharedConfigState: session.SharedConfigEnable,
	}))

	//queueURL := ""
	SNSURL := ""

	err := send.SendSNSMessage(SNSURL, sess)
	if err != nil {
		log.Println("메시지 전송 에러")
	}

	//err := send.SendSqsMessage(sess, &queueURL, "emil", "메시지 내용")
	//if err != nil {
	//	log.Println("메시지 전송 에러")
	//}
}
