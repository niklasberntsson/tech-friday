package health

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestGetHealth(t *testing.T) {
	expected := HealthResponse{
		Status:  "ok",
		Env:     "dev",
		Version: "1.0.0",
	}

	actual := GetHealth()

	assert.Equal(t, expected, actual)
}
