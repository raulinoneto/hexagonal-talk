package main

import (
	"fmt"

	"github.com/lucasrosa/catvotes/internal/adapters/primary"
	"github.com/lucasrosa/catvotes/internal/adapters/secondary"
	"github.com/lucasrosa/catvotes/internal/domains/votes"
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
