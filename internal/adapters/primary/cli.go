package primary

import (
	"flag"
	"fmt"

	"github.com/raulinoneto/catvotes/pkg/domains/votes"
)

type CLIPrimaryAdapter struct {
	service votes.PrimaryPort
}

func NewCLIPrimaryAdapter(s votes.PrimaryPort) *CLIPrimaryAdapter {
	return &CLIPrimaryAdapter{
		s,
	}
}

// PlaceOrder receives the request, processes it and returns a Response or an error
func (a *CLIPrimaryAdapter) HandleVote() (string, error) {
	var clientVote int
	message := "Vote increased"
	v := votes.Vote{}
	v.Vote = true

	flag.StringVar(&v.ImageID, "image", "", "image that will be voted")
	flag.IntVar(&clientVote, "vote", 1, "increase vote to image")
	flag.Parse()

	if clientVote == 0 {
		v.Vote = false
		message = "Vote Decreased"
	}

	if !v.Vote && clientVote != 0 {
		return "", fmt.Errorf("Vote only allows 1 or 0")
	}

	// Processing vote
	err := a.service.Vote(v.ImageID, v.Vote)
	if err != nil || clientVote > 1 {
		return "", err
	}

	return message, nil
}
