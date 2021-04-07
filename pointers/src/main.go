package main

import "fmt"

func main() {

	value_types()

	slices()

	str := "initial"

	// with the pointer the parameter is passed by reference and is changed
	functions_with_pointers(&str)
	fmt.Println(str)

	// without the pointer the parameter is passed by value and is not actually changed
	functions_without_pointers(str)
	fmt.Println(str)
}

func value_types() {
	x := 7

	y := &x
	fmt.Println(x, y)

	*y = 9
	fmt.Println(x, y)
}

func slices() {
	x := []int{1, 2, 3}

	y := &x
	(*y)[0] = 100
	fmt.Println(x, y, *y)

}

func functions_with_pointers(value *string) {
	*value = "changed"
}

func functions_without_pointers(value string) {
	value = "changed again"
}
