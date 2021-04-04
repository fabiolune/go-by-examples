package main

import "fmt"

func main() {
	var customers = GetCustomers()
	fmt.Println(customers)
}

func getUserName(userId int) (name string) {
	return "fabio lonegro"
}
