package in

import "example.com/yourorg/yourservice/internal/domain"

// GreetingUseCases defines inbound ports (use cases) exposed to adapters.
type GreetingUseCases interface {
	CreateGreeting(id, message string) (domain.Greeting, error)
	GetGreeting(id string) (domain.Greeting, error)
	ListGreetings() ([]domain.Greeting, error)
}
