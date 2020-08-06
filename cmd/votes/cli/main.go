package main

import (
	"fmt"

	"github.com/raulinoneto/catvotes/internal/adapters/primary"
	"github.com/raulinoneto/catvotes/internal/adapters/secondary"
	"github.com/raulinoneto/catvotes/pkg/domains/votes"
)

func main() {
	repo := secondary.NewVotesAPI()
	service := votes.NewService(repo)
	primary := primary.NewCLIPrimaryAdapter(service)

	result, err := primary.HandleVote()
	if err != nil {
		fmt.Println(err)
	}

	fmt.Println(result)
}
