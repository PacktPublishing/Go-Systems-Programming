package main

import (
	"fmt"
	"reflect"
)

func main() {

	type message struct {
		X     int
		Y     int
		Label string
	}

	p1 := message{23, 12, "A Message"}
	p2 := message{}
	p2.Label = "Message 2"

	s1 := reflect.ValueOf(&p1).Elem()
	s2 := reflect.ValueOf(&p2).Elem()
	fmt.Println("S2= ", s2)

	typeOfT := s1.Type()
	fmt.Println("P1=", p1)
	fmt.Println("P2=", p2)

	for i := 0; i < s1.NumField(); i++ {
		f := s1.Field(i)
		fmt.Printf("%d: %s ", i, typeOfT.Field(i).Name)
		fmt.Printf("%s = %v\n", f.Type(), f.Interface())
	}

}
