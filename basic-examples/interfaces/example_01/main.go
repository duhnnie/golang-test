package main

import (
	"fmt"
	"strconv"
)

type Stringer interface {
	String() string
}

type Booleano bool

func (b Booleano) String() string {
	if b {
		return "verdadero"
	}

	return "falso"
}

func (b Booleano) Get() bool {
	return bool(b)
}

func toString(v interface{}) string {
	if stringer, ok := v.(Stringer); ok {
		return stringer.String()
	}

	switch t := v.(type) {
	case int:
		return strconv.Itoa(t)
	case float64:
		return fmt.Sprintf("%f", t)
	case string:
		return t
	}

	return "???"
}

// Basic example
func main() {
	s := "hola a todos"
	msg := toString(s)
	fmt.Println(msg)

	b := Booleano(true)
	msg = toString(b)
	fmt.Println(msg)
}
