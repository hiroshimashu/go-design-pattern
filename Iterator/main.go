package main

// Aggregate defines the property of iterator
type Aggregate interface {
	iterator() Iterator
}

// Iterator defines methods
type Iterator interface {
	hasNext() bool
	next() interface{}
}

type book struct {
	name string
}

func NewBook(name string) *book {
	return &book{
		name
	}
}

func main() {

}
