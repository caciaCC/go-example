package main

import "fmt"

func modifyString1(str string) string {
	byteStr := []byte(str)
	for i, _ := range byteStr {
		byteStr[i] = byte('0')
	}
	return string(byteStr)
}
func modifyString2(str string) string {
	runeStr := []rune(str)
	for i, _ := range runeStr {
		runeStr[i] = rune('0')
	}
	return string(runeStr)
}

func main() {
	str := "你好,世界!"
	fmt.Println(str)
	fmt.Println(modifyString1(str))
	fmt.Println(modifyString2(str))
	//a chinese character present 3 byte in UTF-8
}
