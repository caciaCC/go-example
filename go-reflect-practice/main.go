package main

import (
	"fmt"
	"reflect"
)

func main() {
	v := interface{}(8888)
	fmt.Println(getValue(v))
}

func getValue(i interface{}) int {
	fmt.Println()
	dynamicType := reflect.TypeOf(i)
	//Type类型值 --动态类型值 -- 即具体实现的类型值
	// 用于提供类型值
	fmt.Println(dynamicType.Kind())
	if dynamicType.Kind() == reflect.Int {
		//todo
	}
	value := reflect.ValueOf(i)
	//Value类型值
	// 用于获取和修改类型值
	fmt.Println(value.Int())
	return 0
}
