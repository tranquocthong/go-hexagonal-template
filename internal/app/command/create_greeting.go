package command

import (
	"time"

	"example.com/yourorg/yourservice/internal/domain"
	portout "example.com/yourorg/yourservice/internal/domain/port/out"
)

// CreateGreetingHandler handles creating a greeting (command)
type CreateGreetingHandler struct {
	repo portout.GreetingRepository
}

func NewCreateGreetingHandler(repo portout.GreetingRepository) CreateGreetingHandler {
	return CreateGreetingHandler{repo: repo}
}

func (h CreateGreetingHandler) Handle(id, message string) (domain.Greeting, error) {
	g := domain.Greeting{ID: id, Message: message, CreatedAt: time.Now()}
	if g.ID == "" || g.Message == "" {
		return domain.Greeting{}, domain.DomainError{Code: domain.ErrInvalid, Message: "id and message are required"}
	}
	if err := h.repo.Create(g); err != nil {
		return domain.Greeting{}, err
	}
	return g, nil
}
