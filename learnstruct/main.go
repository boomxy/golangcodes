package main

import (
	"fmt"
	"log"
	"os"
)

// 一个普通的结构体
type Person struct {
	Name string
	Age  int
}

type Student struct {
	Person // 匿名字段
	Score  int
}

type Teacher struct {
	Info   Person
	Salary int
}

func simple() {
	// 第一种方式
	var p Person
	p.Name = "Tom"
	p.Age = 18
	fmt.Println(p)
	fmt.Println(p.Name, p.Age)

	// 第二种方式
	p1 := Person{
		Name: "Jerry",
		Age:  20,
	}
	fmt.Println(p1)
	fmt.Println(p1.Name, p1.Age)
}

func nested() {
	// 匿名字段，嵌套结构体
	var s Student
	s.Name = "Jack"
	s.Age = 22
	s.Score = 99
	fmt.Println(s)
	fmt.Println(s.Person.Name, s.Person.Age) // 完整的字段名
	fmt.Println(s.Name, s.Age, s.Score)      // 简写
	s2 := Student{
		Person: Person{
			Name: "Jim",
			Age:  30,
		},
		Score: 100,
	}
	fmt.Println(s2)
	fmt.Println(s2.Person.Name, s2.Person.Age, s2.Score)
	s3 := Student{ // 简写
		Person{"Jim's", 31},
		120,
	}
	fmt.Println(s3)
	fmt.Println(s3.Person.Name, s3.Person.Age, s3.Score)
}

func nested2() {
	// 具体名称字段，嵌套结构体
	var t Teacher
	t.Info.Name = "Jim" // 必须带上具体名称
	t.Info.Age = 30     // 必须带上具体名称
	t.Salary = 10000
	fmt.Println(t)
	fmt.Println(t.Info.Name, t.Info.Age, t.Salary) // 完整的字段名
	// fmt.Println(t.Name, t.Age, t.Salary) // 不能使用简写
	t2 := Teacher{
		Info: Person{
			Name: "Jim's",
			Age:  31,
		},
		Salary: 10000,
	}
	fmt.Println(t2)
	fmt.Println(t2.Info.Name, t2.Info.Age, t2.Salary)
	t3 := Teacher{ // 简写
		Person{"Jim's", 31},
		10000,
	}
	fmt.Println(t3)
	fmt.Println(t3.Info.Name, t3.Info.Age, t3.Salary)
}

func main() {
	log.SetOutput(os.Stdout)
	log.Println("Start")
	// 使用结构体
	fmt.Println("----simple()----")
	simple()
	fmt.Println("----nested()----")
	nested()
	fmt.Println("----nested2()----")
	nested2()
	log.Println("Done")
	os.Exit(0)
}
