package votes

type PrimaryPort interface {
	Vote(image_id string, vote bool) error
}

type SecondaryPort interface {
	SaveVote(vote Vote) error
	//GetImage(image_id string) error
}
