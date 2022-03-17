package main

import (
	"fmt"
	"reflect"
)

var m map[reflect.Kind]func(...interface{}) interface{}

func init() {
	m = make(map[reflect.Kind]func(...interface{}) interface{})

	m[reflect.Int] = func(x ...interface{}) interface{} {
		return func(a int) int {
			return a + 100
		}(x[0].(int))
	}
	m[reflect.String] = func(x ...interface{}) interface{} {
		return func(a, b string) string {
			return a + b
		}(x[0].(string), x[1].(string))
	}

}

func main() {

	var num = 10
	tp := reflect.TypeOf(num)
	fmt.Println(m[tp.Kind()](num))

	st1 := "aaa"
	st2 := "bbb"

	tp = reflect.TypeOf(st1)
	fmt.Println(m[tp.Kind()](st1, st2))

}
