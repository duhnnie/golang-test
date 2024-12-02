package main

import (
	"fmt"
	"reflect"
)

func main() {
	var f float64 = 3.4

	v := reflect.ValueOf(f)
	fmt.Println("type: ", v.Type())
	fmt.Println("kind is float32?", v.Kind() == reflect.Float32)
	fmt.Println("kind is float64?", v.Kind() == reflect.Float64)
	fmt.Println("value:", v.Float())
}
