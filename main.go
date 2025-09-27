package main

import "fmt"

func main() {
	fmt.Println("Hello, World!")
}

type Person struct {
	Name string
	Age  int
}

func (p Person) Greet() string {
	return fmt.Sprintf("Hello, my name is %s and I am %d years old.", p.Name, p.Age)
}
