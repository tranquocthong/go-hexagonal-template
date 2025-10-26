package app

import (
	portout "example.com/yourorg/yourservice/internal/domain/port/out"

	"example.com/yourorg/yourservice/internal/app/command"
	"example.com/yourorg/yourservice/internal/app/query"
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
