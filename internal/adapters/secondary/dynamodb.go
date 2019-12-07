package secondary

import (
	"fmt"
	"os"

	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/dynamodb"
	"github.com/aws/aws-sdk-go/service/dynamodb/dynamodbattribute"
	"github.com/lucasrosa/catvotes/internal/domains/votes"
)

type votesRepository struct{}

// NewDynamoRepository instantiates the repository for this adapter
func NewDynamoRepository() votes.SecondaryPort {
	return &votesRepository{}
}

func (r *votesRepository) SaveVote(v votes.Vote) error {
	fmt.Println("saving vote", v)

	sess, err := session.NewSession(&aws.Config{
		Region: aws.String("us-east-1")},
	)

	svc := dynamodb.New(sess)

	// persistedOrder := PersistedOrder{
	// 	ID:        order.ID,
	// 	Email:     order.Email,
	// 	Amount:    order.Amount,
	// 	Currency:  order.Currency,
	// 	ProductID: order.ProductID,
	// }
	// fmt.Println("Persisting order:", persistedOrder)

	// Marshall the Item into a Map DynamoDB can deal with
	av, err := dynamodbattribute.MarshalMap(v)
	if err != nil {
		fmt.Println("Got error marshalling map:")
		fmt.Println(err.Error())
		return err
	}

	// Create Item in table and return
	input := &dynamodb.PutItemInput{
		Item:      av,
		TableName: aws.String(os.Getenv("TABLE_NAME")),
	}

	_, err = svc.PutItem(input)
	if err != nil {
		fmt.Println("Error while putting message to db", err)
	} else {
		fmt.Println("Success while putting message to db")
	}

	return err
}
