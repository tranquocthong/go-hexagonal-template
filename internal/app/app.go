package app

import (
	"time"

	"example.com/yourorg/yourservice/internal/domain"
	portin "example.com/yourorg/yourservice/internal/domain/port/in"
	portout "example.com/yourorg/yourservice/internal/domain/port/out"
)

type Application struct {
	Greeting GreetingUseCase
}

func NewApplication(greetingRepo portout.GreetingRepository) *Application {
	return &Application{
		Greeting: GreetingUseCase{repo: greetingRepo},
	}
}

type GreetingUseCase struct {
	repo portout.GreetingRepository
}

// Compile-time assertion that GreetingUseCase implements inbound GreetingUseCases port.
var _ portin.GreetingUseCases = (*GreetingUseCase)(nil)

func (u GreetingUseCase) CreateGreeting(id, message string) (domain.Greeting, error) {
	g := domain.Greeting{ID: id, Message: message, CreatedAt: time.Now()}
	if g.ID == "" || g.Message == "" {
		return domain.Greeting{}, domain.DomainError{Code: domain.ErrInvalid, Message: "id and message are required"}
	}
	if err := u.repo.Create(g); err != nil {
		return domain.Greeting{}, err
	}
	return g, nil
}

func (u GreetingUseCase) GetGreeting(id string) (domain.Greeting, error) {
	return u.repo.GetByID(id)
}

func (u GreetingUseCase) ListGreetings() ([]domain.Greeting, error) {
	return u.repo.List()
}
