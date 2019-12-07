package main

import (
	"github.com/aws/aws-lambda-go/events"
	"github.com/aws/aws-lambda-go/lambda"

	"github.com/lucasrosa/catvotes/internal/adapters/primary"
	"github.com/lucasrosa/catvotes/internal/adapters/secondary"
	"github.com/lucasrosa/catvotes/internal/domains/votes"
)

// Handler is our lambda handler invoked by the `lambda.Start` function call
func Handler(request events.APIGatewayProxyRequest) (events.APIGatewayProxyResponse, error) {
	repo := secondary.NewDynamoRepository()
	service := votes.NewService(repo)
	primary := primary.NewAPIGatewayPrimaryAdapter(service)

	return primary.HandleVote(request)
}

func main() {
	lambda.Start(Handler)
}
