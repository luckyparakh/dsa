package main

/*
The Strategy design pattern is commonly used in situations where you need to dynamically switch the algorithm or strategy based on the conditions of the program. Here are some common use cases:

1. **Sorting Algorithms**: You might have a list of items that can be sorted in different ways. For example, a list of products could be sorted by price, by name, or by rating. Each of these sorting methods can be implemented as a separate strategy.

2. **Payment Methods**: In an e-commerce application, you might support multiple payment methods such as credit card, PayPal, or bank transfer. Each payment method can be implemented as a separate strategy.

3. **Compression Algorithms**: If you have a program that needs to compress files, you might support different compression algorithms like ZIP, RAR, or TAR. Each compression algorithm can be implemented as a separate strategy.

4. **Routing Algorithms**: In a navigation or mapping application, you might have different algorithms for calculating the best route, such as shortest distance, least traffic, or fewest turns. Each routing algorithm can be implemented as a separate strategy.

5. **Validation Strategies**: If you have different types of validation for user input, each validation could be a strategy. For example, you might have different validation strategies for email addresses, phone numbers, or social security numbers.

In all these cases, the Strategy pattern allows you to switch algorithms or methods at runtime without modifying the code that uses them. This makes your code more flexible and easier to maintain.


Simple Explanation:
Sure, let's imagine you're playing a video game where your character can attack enemies. Now, there are different ways your character can attack. 
It can use a sword, throw a fireball, or use a bow and arrow.
In this situation, the Strategy pattern is like choosing which attack method to use. 
You can easily switch between using the sword, the fireball, or the bow and arrow depending on what you think is best for the current situation in the game.
So, the Strategy pattern is like having a backpack full of different toys (strategies), and you can choose and swap any toy you want to play with at any time.
*/
type NewAnimal struct {
	Name   string
	walker Walker
}

type Walker interface {
	Walk() string
}

type fast struct{}

func (f fast) Walk() string {
	return "I walk fast"
}

type fastest struct{}

func (f fastest) Walk() string {
	return "I walk very fast"
}

type NewDog struct {
	Animal
}

// Problem is that we have to override the Walk method in each subclass of Animal, even when code is same in few classes.
// Due to this, we are violating DRY principle.
type NewTiger struct {
	Animal
}

type NewLeopard struct {
	Animal
}

func main() {
	dog := NewAnimal{Name: "dog", walker: fast{}}
	println(dog.walker.Walk())
	leopard := NewAnimal{Name: "leopard", walker: fastest{}}
	println(leopard.walker.Walk())
	tiger := NewAnimal{Name: "tiger", walker: fastest{}}
	println(tiger.walker.Walk())
	tiger.walker = fast{}
	println(tiger.walker.Walk())

}
