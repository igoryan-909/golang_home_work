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
	len   int
	front *ListItem
	back  *ListItem
}

func (l *list) Len() int {
	return l.len
}

func (l *list) Front() *ListItem {
	return l.front
}

func (l *list) Back() *ListItem {
	return l.back
}

func (l *list) PushFront(v interface{}) *ListItem {
	current := l.front
	l.len++
	l.front = &ListItem{
		Value: v,
		Next:  current,
		Prev:  nil,
	}

	if l.len == 1 {
		l.back = l.front
	} else {
		current.Prev = l.front
	}
	return l.front
}

func (l *list) PushBack(v interface{}) *ListItem {
	l.len++
	current := l.back
	l.back = &ListItem{
		Value: v,
		Next:  nil,
		Prev:  current,
	}

	if l.len == 1 {
		l.front = l.back
	} else {
		current.Next = l.back
	}
	return l.back
}

func (l *list) Remove(i *ListItem) {
	if i.Prev != nil {
		i.Prev.Next = i.Next
	} else {
		l.front = i.Next
	}
	if i.Next != nil {
		i.Next.Prev = i.Prev
	} else {
		l.back = i.Prev
	}
	l.len--
}

func (l *list) MoveToFront(i *ListItem) {
	if l.len == 1 || i.Prev == nil {
		return
	}
	l.Remove(i)
	l.len++
	l.front.Prev = i
	i.Next = l.front
	i.Prev = nil
	l.front = i
}

func NewList() List {
	return new(list)
}
