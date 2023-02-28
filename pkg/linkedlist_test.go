package ds

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test_LinkedList(t *testing.T) {
	t.Run("single node with int", func(t *testing.T) {
		ll := NewLinkedList[int](5)
		assert.NotNil(t, ll)
		assert.Equal(t, 5, ll.First().Value())
		assert.Equal(t, 5, ll.Last().Value())
		assert.Nil(t, ll.Last().Next())
	})
	t.Run("list with three int values", func(t *testing.T) {
		ll := NewLinkedList[int](10)
		ll.Append(5)
		ll.Append(7)
		assert.Equal(t, 10, ll.First().Value())
		assert.Equal(t, 7, ll.Last().Value())
		assert.Nil(t, ll.Last().Next())
	})
}
