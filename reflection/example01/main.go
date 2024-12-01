package main

import (
	"fmt"
	"reflect"
)

func main() {
	var x float64 = 2.76

	refValuePtr := reflect.ValueOf(&x)
	refTypePtr := reflect.TypeOf(&x)

	fmt.Println(refValuePtr)
	fmt.Println(refValuePtr.String())

	fmt.Println(refTypePtr)
	fmt.Println(refTypePtr.String())

	// -------
	refValue := refValuePtr.Elem()
	fmt.Println(refValue.Float())

	// ------
	refValue.SetFloat(4.5)
	fmt.Println(x)
	refValue.Set(reflect.ValueOf(9.99))
	fmt.Println(x)

	i := refValue.Interface().(float64)
	println(i)
}
