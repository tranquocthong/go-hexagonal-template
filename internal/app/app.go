package app

import (
	portin "example.com/yourorg/yourservice/internal/domain/port/in"
	portout "example.com/yourorg/yourservice/internal/domain/port/out"

	"example.com/yourorg/yourservice/internal/app/command"
	"example.com/yourorg/yourservice/internal/app/query"
	"example.com/yourorg/yourservice/internal/domain"
)

// Application holds all application layer handlers (CQRS)
type Application struct {
	Queries  Queries
	Commands Commands
}

// Queries bundles all query handlers
type Queries struct {
	GetGreetingHandler   query.GetGreetingHandler
	ListGreetingsHandler query.ListGreetingsHandler
}

// Commands bundles all command handlers
type Commands struct {
	CreateGreetingHandler command.CreateGreetingHandler
}

// NewApplication wires command/query handlers with outbound ports
func NewApplication(greetingRepo portout.GreetingRepository) *Application {
	return &Application{
		Queries: Queries{
			GetGreetingHandler:   query.NewGetGreetingHandler(greetingRepo),
			ListGreetingsHandler: query.NewListGreetingsHandler(greetingRepo),
		},
		Commands: Commands{
			CreateGreetingHandler: command.NewCreateGreetingHandler(greetingRepo),
		},
	}
}

// CreateGreeting implements portin.GreetingUseCases
func (a *Application) CreateGreeting(id, message string) (domain.Greeting, error) {
	return a.Commands.CreateGreetingHandler.Handle(id, message)
}

// GetGreeting implements portin.GreetingUseCases
func (a *Application) GetGreeting(id string) (domain.Greeting, error) {
	return a.Queries.GetGreetingHandler.Handle(id)
}

// ListGreetings implements portin.GreetingUseCases
func (a *Application) ListGreetings() ([]domain.Greeting, error) {
	return a.Queries.ListGreetingsHandler.Handle()
}

// compile-time check that *Application satisfies GreetingUseCases
var _ portin.GreetingUseCases = (*Application)(nil)
