package main

import (
	"fmt"
	"reflect"
)

func main() {
	var ui8 uint8 = 'A'

	v := reflect.ValueOf(ui8)
	fmt.Println("type", v.Type())
	fmt.Println("kind id uint8?:", v.Kind() == reflect.Uint8)
	ui8 = uint8(v.Uint()) // It always works with the largest type: uint8 => uint64, so we need to cast it.
}
