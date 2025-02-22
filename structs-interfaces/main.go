package main

import (
	"fmt"
	"reflect"
)

func main() {
	p := Person{Name: "Alice", Age: 30}

	fmt.Println("Name:", p.Name)
	fmt.Println("Age:", p.Age)

	fmt.Println("------------------------------------------------------------------------")

	emp := Employee{
		Person:  Person{Name: "Alice", Age: 30},
		Company: "Tech Corp",
		Salary:  70000,
	}

	fmt.Println("Name:", emp.Name)
	fmt.Println("Age:", emp.Age)
	fmt.Println("Company:", emp.Company)
	fmt.Println("Salary:", emp.Salary)

	fmt.Println("------------------------------------------------------------------------")

	emp.Greet()

	fmt.Println("------------------------------------------------------------------------")

	var s Speaker
	s = Person{Name: "Tasnif"}

	fmt.Println(s.Speak())

	fmt.Println("------------------------------------------------------------------------")

	makeSound(emp)
	makeSound(Raf{})
	fmt.Println("------------------------------------------------------------------------")

	validatePerson(emp)
	validatePerson(p)
	fmt.Println("------------------------------------------------------------------------")

	detectType(emp)
	detectType(p)
	detectType(Raf{})
	detectType(5)
	fmt.Println("------------------------------------------------------------------------")

	p1 := Person{Name: "Alice", Age: 25}
	p2 := Person{Name: "Alice", Age: 25}
	p3 := Person{Name: "Bob", Age: 30}

	fmt.Println(map[bool]string{true: "Same person", false: "Different person"}[p1 == p2])
	fmt.Println(map[bool]string{true: "Same person", false: "Different person"}[p1 == p3])
	fmt.Println("------------------------------------------------------------------------")
}

var _ Speaker = (*Person)(nil)

func detectType(i interface{}) {
	switch v := i.(type) {
	case Person:
		fmt.Println("it's a person")
	case Employee:
		fmt.Println("it's a employee")
	case Raf:
		fmt.Println("it's a rafiur")
	default:
		fmt.Println("Invalid type", reflect.TypeOf(v))
	}
}

func validatePerson(i interface{}) {
	v, ok := i.(Person)
	if ok {
		fmt.Println("Valid person", v)
	} else {
		fmt.Println("Invalid person", v)
	}
}

type Person struct {
	Name  string `json:"name"`
	Age   int    `json:"age"`
	Email string `xml:"email"`
}

type Employee struct {
	Person
	Company string
	Salary  int
}

func (p Person) Greet() {
	fmt.Println("Hello, my name is", p.Name)
}

type Speaker interface {
	Speak() string
}

func (p Person) Speak() string {
	return "Hello my name is " + p.Name
}

type Raf struct{}

func (r Raf) Speak() string {
	return "Pera"
}

func makeSound(s Speaker) {
	fmt.Println(s.Speak())
}
