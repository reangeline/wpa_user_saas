package main

import (
	"context"
	"log"

	"github.com/aws/aws-sdk-go-v2/service/dynamodb"
	"github.com/reangeline/wpa_user_saas/internal/infra/http"

	"github.com/aws/aws-sdk-go-v2/config"
)

func main() {

	cfg, err := config.LoadDefaultConfig(context.TODO(), config.WithRegion("us-east-1"))
	if err != nil {
		log.Fatalf("unable to load SDK config, %v", err)
	}

	svc := dynamodb.NewFromConfig(cfg)

	server := http.NewServerLambda(svc)

	server.ServerHttp()

}
