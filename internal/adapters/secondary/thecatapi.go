package secondary

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"

	"github.com/raulinoneto/catvotes/pkg/domains/votes"
)

type votesAPI struct {
	ImageID string `json:"image_id"`
	Value   int    `json:"value"`
}

// NewDynamoRepository instantiates the repository for this adapter
func NewVotesAPI() votes.SecondaryPort {
	return &votesAPI{}
}

func (a *votesAPI) SaveVote(v votes.Vote) error {
	url := "https://api.thecatapi.com/v1/votes"
	fmt.Println("URL:>", url)

	a.ImageID = v.ImageID
	a.Value = 0
	if v.Vote {
		a.Value = 1
	}

	jsonStr, err := json.Marshal(*a)
	if err != nil {
		return err
	}
	req, err := http.NewRequest("POST", url, bytes.NewBuffer(jsonStr))
	if err != nil {
		return err
	}

	req.Header.Set("x-api-key", "3a5d1212-97f3-4ac3-99fe-84b003c6f590")
	req.Header.Set("Content-Type", "application/json")

	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		return err
	}
	defer resp.Body.Close()

	fmt.Println("response Status:", resp.Status)
	fmt.Println("response Headers:", resp.Header)
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return err
	}
	fmt.Println("response Body:", string(body))

	return nil
}
