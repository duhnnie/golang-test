package main

import (
	"fmt"
	"reflect"
)

func main() {
	var x float64 = 34.56

	v := reflect.ValueOf(x)
	fmt.Println("Settability of v:", v.CanSet())

	p := reflect.ValueOf(&x)
	fmt.Println("Settability of p:", p.CanSet())

	pv := p.Elem()
	fmt.Println("Settability of p.Elem()", pv.CanSet())

	pv.SetFloat(23.76)
	fmt.Println(x)
	fmt.Println(pv.Interface())
}
