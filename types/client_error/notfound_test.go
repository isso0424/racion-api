package client_error_test

import (
	"errors"
	"isso0424/racion-api/types/client_error"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestIsNotFoundError(t *testing.T) {
	err := client_error.CreateNotFound("type", "target", "condition")

	assert.Equal(t, "type", err.Type)
	assert.Equal(t, "target", err.Target)
	assert.Equal(t, "condition", err.Condition)

	result := client_error.IsNotFound(err)
	assert.Equal(t, true, result)

	e := errors.New("")
	result = client_error.IsNotFound(e)
	assert.Equal(t, false, result)

	e = errors.New("not found")
	result = client_error.IsNotFound(e)
	assert.Equal(t, true, result)
}
