package hw04lrucache

import (
	"testing"

	"github.com/stretchr/testify/require"
)

func TestList(t *testing.T) {
	t.Run("empty list", func(t *testing.T) {
		l := NewList()

		require.Equal(t, 0, l.Len())
		require.Nil(t, l.Front())
		require.Nil(t, l.Back())
	})

	t.Run("push front to empty list", func(t *testing.T) {
		l := NewList()

		l.PushFront(10) // [10]
		require.Equal(t, 1, l.Len())

		first := l.Front()
		require.Equal(t, 10, first.Value)
		require.Equal(t, (*ListItem)(nil), first.Next)
		require.Equal(t, (*ListItem)(nil), first.Prev)
	})

	t.Run("push back to empty list", func(t *testing.T) {
		l := NewList()

		l.PushBack(10) // [10]
		require.Equal(t, 1, l.Len())

		first := l.Front()
		require.Equal(t, 10, first.Value)
		require.Equal(t, (*ListItem)(nil), first.Next)
		require.Equal(t, (*ListItem)(nil), first.Prev)
	})

	t.Run("push", func(t *testing.T) {
		l := NewList()

		l.PushFront(10) // [10]
		l.PushFront(20) // [20, 10]
		l.PushBack(30)  // [20, 10, 30]
		require.Equal(t, 3, l.Len())
		first := l.Front() // 20
		last := l.Back()   // 30

		require.Equal(t, 20, first.Value)
		require.Equal(t, 10, first.Next.Value)
		require.Equal(t, (*ListItem)(nil), first.Prev)

		require.Equal(t, 30, last.Value)
		require.Equal(t, (*ListItem)(nil), last.Next)
		require.Equal(t, 10, last.Prev.Value)
	})

	t.Run("remove", func(t *testing.T) {
		l := NewList()

		l.PushFront(10) // [10]
		require.Equal(t, 1, l.Len())

		first := l.Front()
		l.Remove(first)

		require.Equal(t, 0, l.Len())

		l.PushFront(20) // [20]
		l.PushBack(30)  // [20, 30]
		require.Equal(t, 2, l.Len())
		first = l.Front() // 20
		l.Remove(first)   // [30]
		first = l.Front() // 30
		last := l.Back()  // 30
		require.Equal(t, 1, l.Len())

		require.Equal(t, 30, first.Value)
		require.Equal(t, (*ListItem)(nil), first.Next)
		require.Equal(t, (*ListItem)(nil), first.Prev)
		require.Equal(t, 30, last.Value)
		require.Equal(t, (*ListItem)(nil), last.Next)
		require.Equal(t, (*ListItem)(nil), last.Prev)
	})

	t.Run("move to front", func(t *testing.T) {
		l := NewList()

		l.PushFront(30) // [30]
		l.PushFront(20) // [20 30]
		l.PushFront(10) // [10 20 30]
		require.Equal(t, 3, l.Len())

		l.MoveToFront(l.Front()) // [10 20 30]
		require.Equal(t, 10, l.Front().Value)
		require.Equal(t, 20, l.Front().Next.Value)
		require.Equal(t, 30, l.Front().Next.Next.Value)

		l.MoveToFront(l.Front().Next) // [20 10 30]
		require.Equal(t, 20, l.Front().Value)
		require.Equal(t, 10, l.Front().Next.Value)
		require.Equal(t, 30, l.Front().Next.Next.Value)

		l.MoveToFront(l.Front().Next.Next) // [30 20 10]
		require.Equal(t, 30, l.Front().Value)
		require.Equal(t, 20, l.Front().Next.Value)
		require.Equal(t, 10, l.Front().Next.Next.Value)
	})

	t.Run("complex", func(t *testing.T) {
		l := NewList()

		l.PushFront(10) // [10]
		l.PushBack(20)  // [10, 20]
		l.PushBack(30)  // [10, 20, 30]
		require.Equal(t, 3, l.Len())

		middle := l.Front().Next // 20
		l.Remove(middle)         // [10, 30]
		require.Equal(t, 2, l.Len())

		for i, v := range [...]int{40, 50, 60, 70, 80} {
			if i%2 == 0 {
				l.PushFront(v)
			} else {
				l.PushBack(v)
			}
		} // [80, 60, 40, 10, 30, 50, 70]

		require.Equal(t, 7, l.Len())
		require.Equal(t, 80, l.Front().Value)
		require.Equal(t, 70, l.Back().Value)

		l.MoveToFront(l.Front()) // [80, 60, 40, 10, 30, 50, 70]
		l.MoveToFront(l.Back())  // [70, 80, 60, 40, 10, 30, 50]

		elems := make([]int, 0, l.Len())
		for i := l.Front(); i != nil; i = i.Next {
			elems = append(elems, i.Value.(int))
		}
		require.Equal(t, []int{70, 80, 60, 40, 10, 30, 50}, elems)
	})
}
