package main

func main() {
	badStudent := BadStudent{}

	badStudent.Person.Name = "yidane"
	badStudent.Student.Name = "eachen"

	badStudent.Person.Hello()
	badStudent.Student.Hello()
	badStudent.Hello()
}
