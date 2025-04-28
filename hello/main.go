package main

import (
	"fmt"
)

func main() {
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
