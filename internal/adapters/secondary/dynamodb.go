package secondary

import (
	"context"
	"fmt"
	"os"

	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/dynamodb"
	"github.com/aws/aws-sdk-go/service/dynamodb/dynamodbattribute"
	"github.com/raulinoneto/catvotes/pkg/domains/votes"
)

type votesRepository struct{
	ctx context.Context
}

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

	type ImageKey struct {
		ImageID string `json:"image_id"`
	}
	type VoteInc struct {
		Increment int `json:":val"`
	}

	key, err := dynamodbattribute.MarshalMap(ImageKey{
		ImageID: v.ImageID,
	})
	if err != nil {
		return err
	}
	increment, _ := dynamodbattribute.MarshalMap(VoteInc{
		Increment: 1,
	})

	updateExpression := aws.String("set votes = votes - :val")
	if v.Vote {
		updateExpression = aws.String("set votes = votes + :val")
	}

	input := &dynamodb.UpdateItemInput{
		Key:                       key,
		TableName:                 aws.String(os.Getenv("TABLE_NAME")),
		UpdateExpression:          updateExpression,
		ExpressionAttributeValues: increment,
		ReturnValues:              aws.String("UPDATED_NEW"),
	}

	result, err := svc.UpdateItem(input)

	if err != nil {
		fmt.Println(err.Error())
		return err
	}

	fmt.Println("updateitem result", result)
	if err != nil {
		fmt.Println("Error while putting message to db", err)
	} else {
		fmt.Println("Success while putting message to db")
	}

	return err
}
