package main

import "fmt"

var (
	systemIDToProductIDMapping = map[int]string{
		1: "001",
		2: "002",
		3: "004",
		4: "007",
		5: "010",
	}
)

func changeBefore(systemID int) string {
	if systemID == 1 {
		return "001"
	} else if systemID == 2 {
		return "002"
	} else if systemID == 3 {
		return "004"
	} else if systemID == 4 {
		return "007"
	} else if systemID == 5 {
		return "010"
	} else {
		return "unknown"
	}
}

func changeAfter(systemID int) string {
	if val, ok := systemIDToProductIDMapping[systemID]; ok {
		return val
	}
	return "unknown"
}

func main() {
	systemID := 1
	fmt.Println(changeBefore(systemID))
	fmt.Println(changeAfter(systemID))
	//使用表驱动代替繁琐的else if结构
}
