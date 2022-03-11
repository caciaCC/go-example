package main

import "fmt"

func modify(a []int) {
	a[0] = 1
	a[1] = 2
}
func modify2(a []int) []int {
	a = append(a, 3)
	return a
}

func main() {
	var slice []int
	slice = make([]int, 2)
	modify(slice)
	fmt.Println(slice)
	slice = modify2(slice)
	fmt.Println(slice)
}
