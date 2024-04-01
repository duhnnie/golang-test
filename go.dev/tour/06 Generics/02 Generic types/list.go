package main

import "fmt"

// List represents a singly-linked list that holds
// values of any type.
type List[T any] struct {
	next *List[T]
	val  T
}

func (l *List[T]) GetAll() []*List[T] {
	var list []*List[T]
	var current *List[T] = l

	for ; current != nil; {
		list = append(list, current)
		current = current.next
	}

	return list
}

func (l *List[T]) AddNext(element *List[T]) {
	currentNext := (*l).next

	if currentNext != nil {
		element.next = currentNext
	}

	l.next = element
}

func (l *List[T]) GetLast() *List[T] {
	current := l

	for ; current.next != nil; {
		current = current.next
	}

	return current
}

func (l *List[T]) Append(element *List[T]) {
	last := l.GetLast()
	last.next = element
}

func main() {
	firstElement := List[int]{
		val: 0,
	}

	for i := 1; i <= 10; i++ {
		newElement := List[int]{
			val: i,
		}

		firstElement.Append(&newElement)
	}

	fmt.Printf("Last in chain: %v\n", firstElement.GetLast().val)

	// --- now with strings
	stringElement := List[string] {
		val: "Hey,",
	}

	stringElement.AddNext(&List[string]{ val: "I'm" })
	stringElement.Append(&List[string]{ val: "learning" })
	stringElement.Append(&List[string]{ val: "Go" })
	stringElement.Append(&List[string]{ val: "and" })
	stringElement.Append(&List[string]{ val: "I" })
	stringElement.Append(&List[string]{ val: "love" })
	stringElement.Append(&List[string]{ val: "it!" })

	stringElement.AddNext(&List[string]{ val: "Daniel," })
	stringElement.AddNext(&List[string]{ val: "I'm" })

	allWords := stringElement.GetAll()
	wholePhrase := ""

	for i := 0; i < len(allWords); i++ {
		wholePhrase += " " + allWords[i].val
	}

	fmt.Println(wholePhrase)
}
