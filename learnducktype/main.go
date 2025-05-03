package main

import (
	"fmt"
	"os"
	"log"
)

// 定义接口
type Person interface {
	job()
	growUp()
}

type Human struct {
	name string
	age int
}

// Human实现了Person接口
func (h Human) job() {
	log.Println("I'm "+h.name+ ",job is engineer")
}

// Human实现了Person接口
func (h *Human) growUp() {
	h.age++
}

type Student struct {
	name string
	age int
}

// Student实现了Person接口
func (s Student) job() {
	log.Println("I'm "+s.name + ",job is student")
}

// Student实现了Person接口
func (s *Student) growUp() {
	s.age++
}


func whatJob(p Person) {
	// p是接口类型，可以接收任何实现了Person接口的实例
	p.job()
}

func growUp(p Person) {
	// p是接口类型，可以接收任何实现了Person接口的实例
	p.growUp()
}

func main() {
	h := Human{"Joe", 20}
	s := Student{"Bob", 11}
	// whatJob(h)
	// whatJob(s)
	whatJob(&s)
	whatJob(&h)
	fmt.Println(h.age, s.age)
	growUp(&s)
	growUp(&h)
	fmt.Println(h.age, s.age)

	os.Exit(0)

}