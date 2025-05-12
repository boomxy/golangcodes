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

type School struct {
	Name     string
	Addr     string
	Teachers []Teacher
	Leader   *Person
	Students *[]Student
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

func nestedWithPointer() {
	s1 := Student{ // 简写
		Person{"Jim's1", 21},
		121,
	}
	s2 := Student{ // 简写
		Person{"Jim's2", 22},
		122,
	}
	s3 := Student{ // 简写
		Person{"Jim's3", 23},
		123,
	}
	students := &[]Student{s1, s2, s3}
	t1 := Teacher{
		Info: Person{
			Name: "Jim't1",
			Age:  31,
		},
		Salary: 10001,
	}
	t2 := Teacher{
		Info: Person{
			Name: "Jim't1",
			Age:  32,
		},
		Salary: 10002,
	}
	t3 := Teacher{
		Info: Person{
			Name: "Jim't1",
			Age:  33,
		},
		Salary: 10003,
	}
	teachers := []Teacher{t1, t2, t3}
	leader := &Person{
		Name: "Jim'l",
		Age:  51,
	}
	var s School
	s.Name = "BJ"
	s.Addr = "Beijing"
	s.Teachers = teachers
	s.Students = students
	s.Leader = leader
	fmt.Println(s)

	fmt.Println("The school is:", s.Name, ",Address is:", s.Addr)
	fmt.Println("The leader is:", s.Leader, "==Name is:", s.Leader.Name, ",Age is:", s.Leader.Age) // 这种方式更常用
	// 注意： (*s.Leader).Name 和 s.Leader.Name 是等价的， s.Leader.Age 和 (*s.Leader).Age 是等价的
	fmt.Println("The leader is:", s.Leader, "==Name is:", (*s.Leader).Name, ",Age is:", (*s.Leader).Age) 
	fmt.Println("The teachers are:", s.Teachers)
	for _, v := range s.Teachers { // 遍历， 注意s.Teachers
		fmt.Println("The teacher is:", v, "==Name is:", v.Info.Name, ",Age is:", v.Info.Age, ",Salary is:", v.Salary)
	}
	fmt.Println("The students are:", s.Students)

	for _, v := range *s.Students { // 遍历, 注意*s.Students
		fmt.Println("The student is:", v, "==Name is:", v.Name, ",Age is:", v.Age, ",Score is:", v.Score)
	}
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
	fmt.Println("----nestedWithPointer()----")
	nestedWithPointer()
	log.Println("Done")
	os.Exit(0)
}
