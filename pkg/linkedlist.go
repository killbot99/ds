package ds

type Node[T int | string] interface {
	// Returns the next node
	Next() Node[T]

	// creates a new node and sets it as the Next of the current node
	SetNext(T) Node[T]

	// Return the value of the Node data
	Value() T
}

type node[T int | string] struct {
	data T
	next Node[T]
}

// returns the next
func (n *node[T]) Next() Node[T] {
	return n.next
}

func (n *node[T]) Value() T {
	return n.data
}

func (n *node[T]) SetNext(val T) Node[T] {
	n.next = NewNode(val)
	return n.Next()
}

type List[T int | string] interface {
	First() Node[T]
	Last() Node[T]
	Append(T)
}

func (l *list[T]) First() Node[T] {
	return l.first
}

func (l *list[T]) Last() Node[T] {
	current := l.First()

	for current.Next() != nil {
		current = current.Next()
	}

	return current
}

func (l *list[T]) Append(val T) {
	last := l.Last()
	last.SetNext(val)
}

type list[T int | string] struct {
	first Node[T]
}

func NewLinkedList[T int | string](data T) List[T] {
	n := NewNode(data)
	return &list[T]{
		first: n,
	}
}

func NewNode[T int | string](data T) Node[T] {
	n := node[T]{
		data: data,
	}
	return &n
}
