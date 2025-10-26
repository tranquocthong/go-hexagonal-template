package out

import "example.com/yourorg/yourservice/internal/domain"

// GreetingRepository defines outbound port for persisting and fetching greetings.
type GreetingRepository interface {
	Create(g domain.Greeting) error
	GetByID(id string) (domain.Greeting, error)
	List() ([]domain.Greeting, error)
}
