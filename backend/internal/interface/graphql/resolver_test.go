package graphql

import (
	"context"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestQueryResolver_Ping(t *testing.T) {
	resolver := &Resolver{}
	queryResolver := resolver.Query()

	resp, err := queryResolver.Ping(context.Background())
	assert.NoError(t, err)
	assert.Equal(t, "pong", resp)
}
