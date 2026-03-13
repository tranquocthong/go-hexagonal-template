package jwt_test

import (
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"

	"example.com/yourorg/yourservice/pkg/auth/jwt"
)

const testSecret = "test-secret-32-chars-long-enough!"

func TestGenerate_Validate_RoundTrip(t *testing.T) {
	tok, err := jwt.Generate("user@example.com", testSecret, time.Hour)
	require.NoError(t, err)
	require.NotEmpty(t, tok)

	claims, err := jwt.Validate(tok, testSecret)
	require.NoError(t, err)
	assert.Equal(t, "user@example.com", claims.Subject)
}

func TestValidate_ExpiredToken(t *testing.T) {
	tok, err := jwt.Generate("user@example.com", testSecret, -time.Second)
	require.NoError(t, err)

	_, err = jwt.Validate(tok, testSecret)
	assert.Error(t, err)
}

func TestValidate_WrongSecret(t *testing.T) {
	tok, err := jwt.Generate("user@example.com", testSecret, time.Hour)
	require.NoError(t, err)

	_, err = jwt.Validate(tok, "wrong-secret")
	assert.Error(t, err)
}

func TestValidate_TamperedToken(t *testing.T) {
	tok, err := jwt.Generate("user@example.com", testSecret, time.Hour)
	require.NoError(t, err)

	_, err = jwt.Validate(tok+"tampered", testSecret)
	assert.Error(t, err)
}
