package main

import (
	"fmt"
	"reflect"
)

func main() {
	strVar := "500"
	intValue := 0
	_, err := fmt.Sscan(strVar, &intValue)
	fmt.Println(intValue, err, reflect.TypeOf(intValue))
}