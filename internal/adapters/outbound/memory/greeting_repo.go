package memory

import (
	"sync"
	"time"

	"example.com/yourorg/yourservice/internal/domain"
	portout "example.com/yourorg/yourservice/internal/domain/port/out"
)

type greetingRepo struct {
	mu   sync.RWMutex
	byID map[string]domain.Greeting
}

func NewGreetingRepository() portout.GreetingRepository {
	r := &greetingRepo{byID: make(map[string]domain.Greeting)}
	// Seed demo data at construction time
	r.byID["hello"] = domain.Greeting{ID: "hello", Message: "Hello, World!", CreatedAt: time.Now()}
	return r
}

func (r *greetingRepo) Create(g domain.Greeting) error {
	r.mu.Lock()
	defer r.mu.Unlock()
	if _, exists := r.byID[g.ID]; exists {
		return domain.DomainError{Code: domain.ErrAlreadyExists, Message: "greeting already exists"}
	}
	r.byID[g.ID] = g
	return nil
}

func (r *greetingRepo) GetByID(id string) (domain.Greeting, error) {
	r.mu.RLock()
	defer r.mu.RUnlock()
	g, ok := r.byID[id]
	if !ok {
		return domain.Greeting{}, domain.DomainError{Code: domain.ErrNotFound, Message: "greeting not found"}
	}
	return g, nil
}

func (r *greetingRepo) List() ([]domain.Greeting, error) {
	r.mu.RLock()
	defer r.mu.RUnlock()
	out := make([]domain.Greeting, 0, len(r.byID))
	for _, g := range r.byID {
		out = append(out, g)
	}
	return out, nil
}

var _ portout.GreetingRepository = (*greetingRepo)(nil)
