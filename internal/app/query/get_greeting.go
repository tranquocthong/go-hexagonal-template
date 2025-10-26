package query

import (
	"example.com/yourorg/yourservice/internal/domain"
	portout "example.com/yourorg/yourservice/internal/domain/port/out"
)

// GetGreetingHandler handles fetching a greeting by id (query)
type GetGreetingHandler struct {
	repo portout.GreetingRepository
}

func NewGetGreetingHandler(repo portout.GreetingRepository) GetGreetingHandler {
	return GetGreetingHandler{repo: repo}
}

func (h GetGreetingHandler) Handle(id string) (domain.Greeting, error) {
	return h.repo.GetByID(id)
}
