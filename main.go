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
	SQSURL := "https://sqs.ap-northeast-2.amazonaws.com/400281609678/SQS_TEST2.fifo"
	eventType := "donationRegistered"
	message := map[string]interface{}{
		"donationId": 1,
		"memberId":   1,
		"mobile":     "01000000000",
		"content":    "메시지 내용입니다.",
	}
	err := send.SendSqsMessage(sess, &SQSURL, eventType, message)
	if err != nil {
		log.Println("메시지 전송 에러")
	}

}
