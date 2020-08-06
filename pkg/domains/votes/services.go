package votes

import "context"

type port struct {
	repo SecondaryPort
	ctx context.Context
}

// NewService receives a Secondary Port of domain and insantiates a Primary Port
func NewService(repo SecondaryPort) PrimaryPort {
	return &port{
		repo,
		context.Background(),
	}
}

func (p *port) Vote(image_id string, vote bool) error {
	v := Vote{image_id, vote}
	err := p.repo.SaveVote(v)
	return err
}
