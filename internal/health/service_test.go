package health

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestGetHealth(t *testing.T) {
	actual := GetHealth()

	assert.Equal(t, "ok", actual.Status)
	assert.Equal(t, "dev", actual.Env)
	assert.Equal(t, "1.0.0", actual.Version)
	assert.NotEmpty(t, actual.Quote)

	// Verify the quote is in our list
	found := false
	for _, q := range quotes {
		if q == actual.Quote {
			found = true
			break
		}
	}
	assert.True(t, found, "Quote should be one of the Family Guy quotes")
}
