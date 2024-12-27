/*
Sure, let's imagine you have a plain, simple ice cream cone. Now, you want to make it more exciting and tasty. So, you decide to add some chocolate syrup on top. Then, you think it would be even better with some sprinkles. And finally, you put a cherry on top. Now, your simple ice cream cone is a lot more exciting and delicious!

In this situation, the Decorator pattern is like adding toppings to your ice cream. The ice cream cone is your original object, and each topping (chocolate syrup, sprinkles, cherry) is a decorator that adds some new flavor or feature to your ice cream.

So, the Decorator pattern is like having a basic thing and adding more features to it one by one, making it more complex and feature-rich step by step. And the best part is, you can choose exactly what features you want to add!

The Decorator pattern is commonly used in situations where you need to add new functionality to an object dynamically, or when subclassing to extend functionality is not the best solution. Here are some common examples:

1. **Graphical User Interface (GUI) Toolkits**: In GUI toolkits, decorators can be used to add behaviors to individual UI components, like adding scrollbars, borders, or shadows.

2. **Java I/O Classes**: The Java I/O classes use the Decorator pattern extensively. For example, you can wrap a `BufferedReader` around a `FileReader` to read text from a file more efficiently.

3. **Middleware in Web Frameworks**: In web development frameworks like Express.js for Node.js, middleware functions that modify the request or response objects can be seen as decorators.

4. **Python Decorators**: Python's decorator feature is a direct application of the Decorator pattern. It allows you to wrap a function or method with another function to extend its behavior.

5. **Adding Features to Classes**: If you have a class and you want to add features to it, but only for certain objects and not all, you can use the Decorator pattern. For example, in a game, you might have a basic character that can be decorated with additional abilities or attributes.

6. **Logging and Monitoring**: The Decorator pattern can be used to add logging or monitoring functionality to a method without changing its code. The decorator can log the start and end of the method, calculate the time it took to run, or count how many times it was called.

In all these cases, the Decorator pattern allows you to add functionality to objects in a way that is flexible and easy to maintain.

Class explosion is a term used in object-oriented programming to describe a situation where the number of classes in a system becomes excessively large due to the improper use of inheritance and other object-oriented features. This can make the system difficult to understand, maintain, and extend.

Class explosion often occurs when developers try to model every real-world object as a separate class, or when they create a new subclass for every possible variation of a class's behavior. This can lead to a deep and complex inheritance hierarchy, with many small, specialized classes that are used only once.

One way to mitigate class explosion is to use design patterns that favor object composition over class inheritance, such as the Strategy pattern or the Decorator pattern. These patterns allow you to add behavior to objects dynamically, without needing to create a new subclass for each variation.
*/
package main

import (
	"errors"
	"fmt"
)

// It has both has-a and is-a relationship.
// It has-a relationship because it has a reference to the object it decorates.
// It is-a relationship because it is of the same type as the object it decorates.

type IPizza interface {
	GetPrice() (int, error)
}

type Pizza struct{}

func (p *Pizza) GetPrice() (int, error) {
	return 7, nil
}

type CheeseDecorator struct {
	pizza IPizza
}

func (c *CheeseDecorator) GetPrice() (int, error) {
	if c.pizza == nil {
		return 0, errors.New("Pizza is nil")
	}
	price, err := c.pizza.GetPrice()
	if err != nil {
		return 0, err
	}
	return price + 5, nil
}

type TomatoDecorator struct {
	pizza IPizza
}

func (t *TomatoDecorator) GetPrice() (int, error) {
	if t.pizza == nil {
		return 0, errors.New("Pizza is nil")
	}
	price, err := t.pizza.GetPrice()
	if err != nil {
		return 0, err
	}
	return price + 2, nil
}

func main() {
	pizza := &Pizza{}
	pizzaWithCheese := &CheeseDecorator{pizza}
	pizzaWithCheeseAndTomato := &TomatoDecorator{pizzaWithCheese}
	fmt.Println(pizzaWithCheeseAndTomato.GetPrice()) // 14
}
