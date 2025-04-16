package handler

import (
	"testing"

	"github.com/stretchr/testify/require"
)

func TestHandlersJsonError(t *testing.T) {
	msg := "Hello Json"
	result := jsonError(msg)
	require.Equal(t, `{"message":"Hello Json"}`, string(result))
}
