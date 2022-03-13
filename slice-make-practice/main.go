package main

import "fmt"

func main() {
	intSlice := make([]int, 5)
	fmt.Println(intSlice)
	//默认值
	var slice []int
	fmt.Println(slice)
	if slice == nil {
		fmt.Println("slice is nil")
	}
	//切片、map、channel本身是引用类型
	slice = append(slice, 1)
	fmt.Println(slice)
	if slice == nil {
		fmt.Println("slice is nil")
	}
	intSlice = make([]int, 5, 6)
	fmt.Println(intSlice)
	intSlice = append(intSlice, 1)
	fmt.Println(cap(intSlice))
	bSlice := intSlice[:2]
	fmt.Println(bSlice)
	intSlice = append(intSlice, 1)
	fmt.Println(cap(intSlice))
	intSlice[0] = 1
	intSlice[1] = 1
	fmt.Println(bSlice)
	fmt.Println(intSlice)
	// *2 扩容 扩容前，bSlice为intSlice一部分，扩容后intSlice指向新地址，不再同于bSlice
}
