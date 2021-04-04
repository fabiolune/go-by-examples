package main

type (
	Customer struct {
		FirstName string
		LastName  string
	}
)

func GetCustomers() (customers []Customer) {
	return append(customers,
		Customer{FirstName: "fabio", LastName: "lonegro"},
		Customer{FirstName: "giulia", LastName: "dalmasso"},
	)
}
