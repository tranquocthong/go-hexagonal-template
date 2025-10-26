package query

import (
	"example.com/yourorg/yourservice/internal/domain"
	portout "example.com/yourorg/yourservice/internal/domain/port/out"
)

// ListGreetingsHandler handles listing greetings (query)
type ListGreetingsHandler struct {
	repo portout.GreetingRepository
}

func NewListGreetingsHandler(repo portout.GreetingRepository) ListGreetingsHandler {
	return ListGreetingsHandler{repo: repo}
}

func (h ListGreetingsHandler) Handle() ([]domain.Greeting, error) {
	return h.repo.List()
}
