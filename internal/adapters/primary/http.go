package primary

import (
	"fmt"
	"github.com/labstack/echo"
	"net/http"

	"github.com/raulinoneto/catvotes/pkg/domains/votes"
)

type HttpPrimaryAdapter struct {
	service votes.PrimaryPort
}

func NewHttpPrimaryAdapter(s votes.PrimaryPort) *HttpPrimaryAdapter {
	return &HttpPrimaryAdapter{
		s,
	}
}

// PlaceOrder receives the request, processes it and returns a Response or an error
func (a *HttpPrimaryAdapter) HandleVote(c echo.Context) error {

	// Verifying the body of the request
	v := new(votes.Vote)
	err := c.Bind(v)
	if err != nil {
		return c.JSON(http.StatusBadRequest, fmt.Sprintf(`{message: %s}`, err.Error()))
	}

	// Processing vote
	err = a.service.Vote(v.ImageID, v.Vote)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, fmt.Sprintf(`{message: %s}`, err.Error()))
	}

	return c.JSON(http.StatusCreated, `{message: "vote has been saved"}`)
}
