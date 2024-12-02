package main

import (
	"fmt"
	"reflect"
)

type MyInt int

func main() {
	mi := MyInt(5)
	v := reflect.ValueOf(mi)
	fmt.Println(v.Interface())

	// or: mi, err := v.Interface().(MyInt)
}
