package main

import (
	"fmt"
	"strings"
)

type Person struct {
	Name string
	Age  int
}

type Car struct {
	Name  string
	Price int
}

type stringPrint func(s string)


func (sp stringPrint) String() {
	sp("hello worlddd")
}

func CommonString(s string) {
	fmt.Println(s + "ZZzzZ")
}

func main() {

	sp:=stringPrint(CommonString)
	sp.String()



	C := &Car{
		Name:  "BMW",
		Price: 1000000,
	}
	fmt.Println(C)
	P := Person{
		Name: "John",
		Age:  30,
	}
	fmt.Println(P)
	fmt.Println("hello world", "hello world")

	// ss := "hello world"
	fmt.Println("ss" + strings.Join([]string{"hello", "world"}, "**") + "ss")

	type Size struct {
		Height float32 `json:"height"`
		Weight float32 `json:"weight"`
	}

	type Rectangle struct {
		Size `json:"size"`
	}

	var rect = Rectangle{
		Size: Size{
			Height: 10,
			Weight: 20,
		},
	}
	fmt.Println("hello world", rect)

}
