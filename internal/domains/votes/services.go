package votes

type port struct {
	repo SecondaryPort
}

// NewService receives a Secondary Port of domain and insantiates a Primary Port
func NewService(repo SecondaryPort) PrimaryPort {
	return &port{
		repo,
	}
}

func (p *port) Vote(image_id string, vote bool) error {
	// TODO: check if vote exists
	v := Vote{image_id, vote}
	err := p.repo.SaveVote(v)
	return err
}
