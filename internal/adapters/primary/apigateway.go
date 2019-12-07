package primary

import (
	"encoding/json"

	"github.com/aws/aws-lambda-go/events"
	"github.com/lucasrosa/catvotes/internal/domains/votes"
)

// APIGatewayAdapter is the interface that defines the entrypoints to this adapter
// type APIGatewayAdapter interface {
// 	PlaceOrder(request events.APIGatewayProxyRequest) (Response, error)
// }

type APIGatewayPrimaryAdapter struct {
	service votes.PrimaryPort
}

func NewAPIGatewayPrimaryAdapter(s votes.PrimaryPort) APIGatewayPrimaryAdapter {
	return &apiGatewayPrimaryAdapter{
		s,
	}
}

// Response is of type APIGatewayProxyResponse since we're leveraging the
// AWS Lambda Proxy Request functionality (default behavior)
//
// https://serverless.com/framework/docs/providers/aws/events/apigateway/#lambda-proxy-integration
type Response events.APIGatewayProxyResponse

// PlaceOrder receives the request, processes it and returns a Response or an error
func (a *APIGatewayPrimaryAdapter) HandleVote(request events.APIGatewayProxyRequest) (Response, error) {

	// Verifying the body of the request
	v := votes.Vote{}
	err := json.Unmarshal([]byte(request.Body), &order)
	if err != nil {
		return Response{StatusCode: 400}, nil
	}

	// Processing vote
	err = a.service.Vote(&v)
	if err != nil {
		return Response{StatusCode: 502}, err
	}

	return successfulResponse(), nil
}

func successfulResponse() Response {
	return Response{
		StatusCode:      201,
		IsBase64Encoded: false,
		Headers: map[string]string{
			"Content-Type":                     "application/json",
			"Access-Control-Allow-Credentials": "true",
			"Access-Control-Allow-Origin":      "*",
			"Access-Control-Allow-Methods":     "POST",
			"Access-Control-Allow-Headers":     "application/json",
		},
	}
}
