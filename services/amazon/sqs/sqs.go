package sqs

import (
	"context"
	"encoding/json"
	"fmt"
	appConfig "github.com/InsuranceTech/shared/config"

	"github.com/InsuranceTech/shared/log"
	"github.com/InsuranceTech/shared/services/amazon/sqs/model"
	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/config"
	"github.com/aws/aws-sdk-go-v2/service/sqs"
	"time"
)

var (
	_log      = log.CreateTag("Sqs")
	ctx       context.Context
	cfg       *appConfig.Config
	sqsClient *sqs.Client
)

// NewAwsSqs Return new AWS Simple Queue System instance
func NewAwsSqs(context context.Context, appConfig *appConfig.Config, region string) {

	cfg_sqs, err := config.LoadDefaultConfig(ctx, config.WithRegion(region))
	if err != nil {
		panic("configuration error, " + err.Error())
	}

	client := sqs.NewFromConfig(cfg_sqs)
	println("SQS Client Initialized")

	ctx = context
	cfg = appConfig
	sqsClient = client
	return
}

func SendPushNotification(req_body model.NotifyBaseRequest) (record int64) {
	var urlResult *sqs.GetQueueUrlOutput
	var msgBody []byte
	var err error

	timezone := time.FixedZone("GMT+3", 10800)

	gQInput := &sqs.GetQueueUrlInput{
		QueueName: aws.String("push_notify"),
	}

	urlResult, err = sqsClient.GetQueueUrl(ctx, gQInput)
	if err != nil {
		_log.ErrorF("Error when try to get QueueURL : %s", err.Error())
	}
	queueURL := urlResult.QueueUrl

	/*
		msgCompose := ent.PushNormalDataWithTemplate{
			ApplicationId: "43070faa8d0a46e7b079de6fcd4044fc",
			TemplateId:    s.cfg.OneSignalTemplates.ConfirmRequest,
			Destination:   "doKmnW5N-oY:APA91bGXMP_TVtY9q_1dBcZBOvjU12sxSZn_NKZARyit8m3vhcd9mtm-xE-KayJHq6AOpqTN9sULII_hdWpQ3rK-7GowER3hv8maaYte7fxh2c-2_psVm9Yg05ppQE0OnXeYmaCib2oh",
			Data: map[string]string{
				"user_title": req_body.UserTitle,
				"type_of":    strconv.FormatInt(3, 10),
				"created_at": fmt.Sprintf("%s", time.Now().In(timezone)),
			},
			CustomData: map[string]string{
				"USER_TITLE": req_body.UserTitle,
			},
		}
	*/
	msgCompose := model.ComposePushMessage{
		ApplicationId: cfg.Sqs.APP_ID,
		Action:        "URL",
		Image:         "",
		Icon:          "",
		Data: map[string]string{
			"user_title": "test",
			"created_at": fmt.Sprintf("%s", time.Now().In(timezone)),
		},
		Priority:    "HIGH",
		Destination: req_body.FcmToken,
		Title:       req_body.Title,
		Message:     req_body.Message,
		Url:         "https://deutschesoft.net",
	}
	msgBody, err = json.Marshal(msgCompose)

	println("msgBody: ", string(msgBody))

	if err != nil {
		println("err:", err.Error())
	}

	message, err := sqsClient.SendMessage(ctx, &sqs.SendMessageInput{
		MessageBody:  aws.String(string(msgBody)),
		QueueUrl:     queueURL,
		DelaySeconds: 0,
	})
	if err != nil {
		println("err: ", err.Error())
	}
	fmt.Printf("message: %v", message)
	return
}
