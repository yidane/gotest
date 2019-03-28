package main

import "fmt"

type Person struct {
	Name string
}

func (person Person) Hello() {
	fmt.Println("Person.Hi")
}

type Student struct {
	Name string
}

func (student Student) Hello() {
	fmt.Println("Student.Hello")
}

type BadStudent struct {
	Student
	Person
}

func (badStudent BadStudent) Hello() {
	fmt.Println("BadStudent.Hello")
}
