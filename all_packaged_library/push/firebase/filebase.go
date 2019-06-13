package firebase

import (
	"context"
	firebase "firebase.google.com/go"
	"firebase.google.com/go/messaging"
	"fmt"
	"google.golang.org/api/option"
)

var FireBaseClient *messaging.Client

func InitFireBaseClient() {
	var serviceAccountKey = []byte(`{
  "type": "123123",
  "project_id": "123123123",
  "private_key_id": "123123123123123",
  "private_key": "-----BEGIN PRIVATE KEY-----\n\n-----END PRIVATE KEY-----\n",
  "client_email": "123123123",
  "client_id": "123123123123",
  "auth_uri": "https://accounts.google.com/o/oauth2/auth",
  "token_uri": "https://oauth2.googleapis.com/token",
  "auth_provider_x509_cert_url": "https://www.googleapis.com/oauth2/v1/certs",
  "client_x509_cert_url": "123123123"
}`)
	opt := option.WithCredentialsJSON(serviceAccountKey)
	app, err := firebase.NewApp(context.Background(), nil, opt)
	if err != nil {
		panic(err)
	}
	ctx := context.Background()
	FireBaseClient, err = app.Messaging(ctx)
	if err != nil {
		panic(err)
	}
}

func FireBaseSendMsgToToken(msg string, token string) {
	notification := &messaging.Notification{
		Title: "推送标题",
		Body:  msg,
	}
	message := &messaging.Message{
		Notification: notification,
		Webpush: &messaging.WebpushConfig{
			Notification: &messaging.WebpushNotification{
				Title: "title",
				Body:  "body",
				Icon:  "icon",
			},
		},
		Token: token,
	}
	response, err := FireBaseClient.Send(context.Background(), message)
	if err != nil {
		panic(err)
	}
	fmt.Println(response)
}
