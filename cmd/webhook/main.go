package main

import (
	"os"

	"github.com/aws/aws-lambda-go/lambda"
	"github.com/k8s-devops-tools/krew-release-bot/pkg/releaser"
)

func main() {
	ghToken := os.Getenv("GH_TOKEN")
	releaser := releaser.New(ghToken)

	lambda.Start(releaser.HandleActionLambdaWebhook)
}
