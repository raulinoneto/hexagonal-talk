package secondary

import (
	"errors"
	"fmt"
	"github.com/raulinoneto/catvotes/pkg/domains/votes"
)

type Printer struct {
	print bool
}

func NewPrinter(print bool) *Printer{
	return &Printer{print}
}

func (p *Printer) SaveVote(vote votes.Vote) error {
	if p.print {
		fmt.Printf("Vote: %+v\n", vote)
		return nil
	}
	return errors.New("vote not printed")
}