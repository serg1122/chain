package chain

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test_Node(t *testing.T) {

	v1 := "v1"
	n1 := &Node{
		prev:  nil,
		next:  nil,
		value: v1,
	}

	assert.False(t, n1.HasPrev())
	assert.False(t, n1.HasNext())

	assert.Equal(t, v1, n1.Value())

	v2 := "v2"
	n2 := &Node{
		prev:  n1,
		next:  nil,
		value: v2,
	}

	n1.setNext(n2)

	assert.True(t, n1.HasNext())
	assert.True(t, n2.HasPrev())
	assert.False(t, n2.HasNext())

	assert.IsType(t, &Node{}, n1.Next())
	assert.IsType(t, &Node{}, n2.Prev())

	assert.Equal(t, v2, n2.Value())

	assert.Equal(t, n2.Value(), n1.Next().Value())
	assert.Equal(t, n1.Value(), n2.Prev().Value())

	v3 := "v3"
	n3 := &Node{
		prev:  n2,
		next:  nil,
		value: v3,
	}

	n2.setNext(n3)

	assert.True(t, n2.HasNext())
	assert.True(t, n3.HasPrev())
	assert.False(t, n3.HasNext())

	assert.IsType(t, &Node{}, n2.Next())
	assert.IsType(t, &Node{}, n3.Prev())

	assert.Equal(t, v3, n3.Value())

	assert.Equal(t, n2.Value(), n3.Prev().Value())
	assert.Equal(t, n3.Value(), n2.Next().Value())

	assert.Equal(t, n1.Value(), n3.Prev().Prev().Value())
	assert.Equal(t, n3.Value(), n1.Next().Next().Value())

	v1updated := "v3updated"
	n1.SetValue(v1updated)

	assert.Equal(t, v1updated, n1.Value())

	v2updated := "v2updated"
	n2.SetValue(v2updated)

	assert.Equal(t, v2updated, n1.Next().Value())
}
