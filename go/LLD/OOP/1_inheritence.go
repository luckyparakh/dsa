package main

type Animal struct {
	Name string
}

func (a Animal) Speak() string {
	return "I am an animal"
}

type Dog struct {
	Animal
	Name string
}

type Cat struct {
	Animal
}

func (d Dog) Speak() string {
	return "I am a dog"
}

func main() {
	dog := Dog{Animal{"dog"}, "Luke"}
	cat := Cat{Animal{"cat"}}
	println(dog.Name)
	println(dog.Speak())
	println(cat.Name)
	println(cat.Speak())

}
