package main

import (
	"fmt"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/sns"
	"send-message-service/send"
	"strings"
)

func main() {
	sess := session.Must(session.NewSessionWithOptions(session.Options{
		SharedConfigState: session.SharedConfigEnable,
	}))

	svc := sns.New(sess)

	//토픽 리스트 보여주기
	topics, err := svc.ListTopics(nil)
	if err != nil {
		fmt.Println(err.Error())
	}
	arnURL := ""
	for _, v := range topics.Topics {
		if strings.Contains(*v.TopicArn, "email-test") { //test-ysh.fifo
			arnURL = *v.TopicArn
		}
	}

	send.SendMessage(arnURL, svc)

}
