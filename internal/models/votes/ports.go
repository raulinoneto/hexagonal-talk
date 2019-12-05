package ports

type Primary interface {	
	Vote(image_id string, vote bool) error 
}

type Secondary interface {
	SaveVote(image_id string, vote bool) error
    GetImage(image_id string) error
}
