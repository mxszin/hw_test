package hw04lrucache

type List interface {
	Len() int
	Front() *ListItem
	Back() *ListItem
	PushFront(v interface{}) *ListItem
	PushBack(v interface{}) *ListItem
	Remove(i *ListItem)
	MoveToFront(i *ListItem)
}

type ListItem struct {
	Value interface{}
	Next  *ListItem
	Prev  *ListItem
}

type list struct {
	head  *ListItem
	tail  *ListItem
	count int
}

func (l *list) Len() int {
	return l.count
}

func (l *list) Front() *ListItem {
	return l.head
}

func (l *list) Back() *ListItem {
	return l.tail
}

func (l *list) PushFront(v interface{}) *ListItem {
	i := &ListItem{Value: v}
	if l.count == 0 {
		l.head = i
		l.tail = i
	} else {
		nextItem := l.head
		l.head = i

		if nextItem != nil {
			nextItem.Prev = i
		}

		i.Next = nextItem
	}

	l.count++

	return i
}

func (l *list) PushBack(v interface{}) *ListItem {
	i := &ListItem{Value: v}
	if l.count == 0 {
		l.head = i
		l.tail = i
	} else {
		prevItem := l.tail
		l.tail = i

		if prevItem != nil {
			prevItem.Next = i
		}

		i.Prev = prevItem
	}

	l.count++

	return i
}

func (l *list) Remove(i *ListItem) {
	// if it is head
	if i.Prev == nil {
		l.head = i.Next
	} else {
		i.Prev.Next = i.Next
	}

	// if it is tail
	if i.Next == nil {
		l.tail = i.Prev
	} else {
		i.Next.Prev = i.Prev
	}

	l.count--
}

func (l *list) MoveToFront(i *ListItem) {
	if l.count <= 1 {
		return
	}

	// if it is head
	if i.Prev == nil {
		return
	}
	// unbind element
	i.Prev.Next = i.Next

	// if it is tail
	if i.Next == nil {
		// set tail
		l.tail = i.Prev
	} else {
		// unbind element
		i.Next.Prev = i.Prev
	}

	// bind with head
	i.Next = l.head
	i.Prev = nil
	// set as head
	l.head.Prev = i
	l.head = i

	// l.Remove(i)
	// l.PushFront(i.Value)
}

func NewList() List {
	return new(list)
}
