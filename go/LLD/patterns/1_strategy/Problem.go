package main

type Animal struct {
	Name string
}

func (a Animal) Walk() string {
	return "I walk fast"
}

type Dog struct {
	Animal
}

// Problem is that we have to override the Walk method in each subclass of Animal, even when code is same in few classes.
// Due to this, we are violating DRY principle.
type Tiger struct {
	Animal
}

func (t Tiger) Walk() string {
	return "I walk very fast"
}

type Leopard struct {
	Animal
}

func (l Leopard) Walk() string {
	return "I walk very fast"
}

func main() {
	dog := Dog{Animal{"dog"}}
	println(dog.Walk())
	tiger:= Tiger{Animal{"tiger"}}
	println(tiger.Walk())
	leopard:= Leopard{Animal{"leopard"}}
	println(leopard.Walk())
}
