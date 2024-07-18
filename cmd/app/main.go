package main

import (
	"github.com/reangeline/wpa_user_saas/internal/infra/http"

	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/dynamodb"
)

func main() {
	sess := session.Must(session.NewSessionWithOptions(session.Options{
		SharedConfigState: session.SharedConfigEnable,
	}))

	svc := dynamodb.New(sess)

	server := http.NewServerLambda(svc)
	server.ServerHttp()
}
