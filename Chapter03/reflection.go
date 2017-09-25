package main

import (
	"fmt"
	"reflect"
)

func main() {

	type t1 int
	type t2 int

	x1 := t1(1)
	x2 := t2(1)
	x3 := 1

	st1 := reflect.ValueOf(&x1).Elem()
	st2 := reflect.ValueOf(&x2).Elem()
	st3 := reflect.ValueOf(&x3).Elem()

	typeOfX1 := st1.Type()
	typeOfX2 := st2.Type()
	typeOfX3 := st3.Type()

	fmt.Printf("X1 Type: %s\n", typeOfX1)
	fmt.Printf("X2 Type: %s\n", typeOfX2)
	fmt.Printf("X3 Type: %s\n", typeOfX3)

	type aStructure struct {
		X    uint
		Y    float64
		Text string
	}

	x4 := aStructure{123, 3.14, "A Structure"}
	st4 := reflect.ValueOf(&x4).Elem()
	typeOfX4 := st4.Type()

	fmt.Printf("X4 Type: %s\n", typeOfX4)
	fmt.Printf("The fields of %s are:\n", typeOfX4)

	for i := 0; i < st4.NumField(); i++ {
		fmt.Printf("%d: Field name: %s ", i, typeOfX4.Field(i).Name)
		fmt.Printf("Type: %s ", st4.Field(i).Type())
		fmt.Printf("and Value: %v\n", st4.Field(i).Interface())
	}
}
