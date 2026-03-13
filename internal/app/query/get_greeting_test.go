package query_test

import (
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"

	"example.com/yourorg/yourservice/internal/adapters/outbound/memory"
	"example.com/yourorg/yourservice/internal/app/query"
	"example.com/yourorg/yourservice/internal/domain"
)

func TestGetGreeting_SeededRecord(t *testing.T) {
	h := query.NewGetGreetingHandler(memory.NewGreetingRepository())
	g, err := h.Handle("hello")
	require.NoError(t, err)
	assert.Equal(t, "hello", g.ID)
	assert.Equal(t, "Hello, World!", g.Message)
}

func TestGetGreeting_NotFound(t *testing.T) {
	h := query.NewGetGreetingHandler(memory.NewGreetingRepository())
	_, err := h.Handle("nonexistent")
	var de domain.DomainError
	require.ErrorAs(t, err, &de)
	assert.Equal(t, domain.ErrNotFound, de.Code)
}

func TestListGreetings_ReturnsSeedData(t *testing.T) {
	h := query.NewListGreetingsHandler(memory.NewGreetingRepository())
	items, err := h.Handle()
	require.NoError(t, err)
	assert.Len(t, items, 1)
	assert.Equal(t, "hello", items[0].ID)
}
