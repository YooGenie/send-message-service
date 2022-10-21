package send

import (
	"flag"
	"fmt"
	"github.com/aws/aws-sdk-go/service/sns"
	"os"
)

func SendMessage(arn string, svc *sns.SNS) {
	msgPtr := flag.String("m", "테스트입니다ㅋㅋㅋ", "The message to send to the subscribed users of the topic")
	topicPtr := flag.String("t", arn, "The ARN of the topic to which the user subscribes")
	flag.Parse()

	result, err := svc.Publish(&sns.PublishInput{
		Message:  msgPtr,
		TopicArn: topicPtr,
	})
	if err != nil {
		fmt.Println(err.Error())
		os.Exit(1)
	}

	fmt.Println(*result.MessageId)
}
