package util

import "strconv"

type Person struct {
	Name string
	Age  int
}
type Employee struct {
	EmployeeId int
	Person     Person
}

func (e Employee) PrintInfo() string {
	return "EmployeeId: " + strconv.Itoa(e.EmployeeId) + ", Name: " + e.Person.Name + ", Age: " + strconv.Itoa(e.Person.Age)
}
