package command_test

import (
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"

	"example.com/yourorg/yourservice/internal/adapters/outbound/memory"
	"example.com/yourorg/yourservice/internal/app/command"
	"example.com/yourorg/yourservice/internal/domain"
)

func newHandler() command.CreateGreetingHandler {
	return command.NewCreateGreetingHandler(memory.NewGreetingRepository())
}

func TestCreateGreeting_Success(t *testing.T) {
	h := newHandler()
	g, err := h.Handle("id-1", "Hello")
	require.NoError(t, err)
	assert.Equal(t, "id-1", g.ID)
	assert.Equal(t, "Hello", g.Message)
	assert.False(t, g.CreatedAt.IsZero())
}

func TestCreateGreeting_EmptyID(t *testing.T) {
	h := newHandler()
	_, err := h.Handle("", "Hello")
	assert.ErrorAs(t, err, &domain.DomainError{})

	var de domain.DomainError
	require.ErrorAs(t, err, &de)
	assert.Equal(t, domain.ErrInvalid, de.Code)
}

func TestCreateGreeting_EmptyMessage(t *testing.T) {
	h := newHandler()
	_, err := h.Handle("id-1", "")
	var de domain.DomainError
	require.ErrorAs(t, err, &de)
	assert.Equal(t, domain.ErrInvalid, de.Code)
}

func TestCreateGreeting_Duplicate(t *testing.T) {
	h := newHandler()
	_, err := h.Handle("id-1", "Hello")
	require.NoError(t, err)

	_, err = h.Handle("id-1", "Hello again")
	var de domain.DomainError
	require.ErrorAs(t, err, &de)
	assert.Equal(t, domain.ErrAlreadyExists, de.Code)
}
