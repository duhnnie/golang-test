package main

import (
	"fmt"
	"reflect"
)

func main() {
	var f float64 = 3.4
	fmt.Println("type:", reflect.TypeOf(f))
}
