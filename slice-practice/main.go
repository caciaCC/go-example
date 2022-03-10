package main

import "fmt"

type Customer struct {
}

var CustomerList *[]Customer

func main() {
	*CustomerList = append(*CustomerList, Customer{})
	fmt.Println(CustomerList)
}
