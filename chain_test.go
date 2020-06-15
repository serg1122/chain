package chain

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test_Chain(t *testing.T) {

	v0 := "v0"
	v1 := "v1"
	v2 := "v2"

	chain := New()

	assert.True(t, chain.IsEmpty())

	assert.Nil(t, chain.First())
	assert.Nil(t, chain.Last())

	chain.Append(&Node{value: v1})

	assert.False(t, chain.IsEmpty())

	assert.IsType(t, &Node{}, chain.First())
	assert.IsType(t, &Node{}, chain.Last())

	assert.Equal(t, v1, chain.First().Value())
	assert.Equal(t, chain.First().Value(), chain.Last().Value())

	chain.Append(&Node{value: v2})

	assert.False(t, chain.IsEmpty())

	assert.IsType(t, &Node{}, chain.Last())

	assert.Equal(t, v1, chain.First().Value())
	assert.Equal(t, v2, chain.Last().Value())

	chain.Prepend(&Node{value: v0})

	assert.False(t, chain.IsEmpty())

	assert.IsType(t, &Node{}, chain.First())

	assert.Equal(t, v0, chain.First().Value())
	assert.Equal(t, v2, chain.Last().Value())
}
