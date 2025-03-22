package pointer

import (
	"fmt"
	"testing"
)

type Person struct {
	Name string
	Age  int
}

func modifyPerson(p Person) {
	fmt.Printf("Address of modifyPerson person: %p\n", &p)
	p.Age = 30
}

func modifyPerson1(p *Person) {
	fmt.Printf("Address of modifyPerson1 person: %p\n", &p)
	p.Age = 60
}

func Test(t *testing.T) {
	person := Person{Name: "Alice", Age: 25}
	fmt.Printf("Address of person: %p\n", &person)
	modifyPerson(person)
	fmt.Println(person.Age) // 输出: 25

	person2 := Person{Name: "Alice", Age: 50}
	fmt.Printf("Address of person2: %p\n", &person2)
	modifyPerson1(&person2)
	fmt.Println(person2.Age) // 输出: 60
}
