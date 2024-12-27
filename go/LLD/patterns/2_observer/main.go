package main

import "fmt"

// Observer pattern is used when there is one-to-many relationship between objects such as if one object is modified, its dependent objects are to be notified automatically.
// Observer pattern falls under behavioral pattern category.


/*
The Observer pattern is widely used in programming, especially in event-driven systems. Here are some practical use cases:

1. **GUI Tools**: In graphical user interface (GUI) tools, the Observer pattern is used to listen for events like button clicks, mouse movements, key presses, etc. When an event occurs, all registered observers are notified.

2. **Real-Time Data Monitoring**: In real-time data monitoring systems, the Observer pattern can be used to notify interested parties when the data changes. For example, a stock market application might use the Observer pattern to notify users when the price of a particular stock changes.

3. **Social Media Platforms**: Social media platforms like Facebook or Twitter use the Observer pattern to notify users of new posts or messages. When a user posts a new message, all their followers (observers) are notified.

4. **Email and SMS Notification Systems**: The Observer pattern is used in email and SMS notification systems. When a specific event occurs (like a new product launch), all registered users are notified via email or SMS.

5. **Model-View-Controller (MVC) Architectures**: In MVC architectures, the model uses the Observer pattern to notify the view whenever its state changes, so that the view can update itself.


Sure, let's imagine you and your friends have a favorite tree in a park that you love to climb. But sometimes, the tree has ripe fruits that you all love to eat. Now, instead of going to the park every day to check if the fruits are ripe, you ask the park's gardener to tell you when the fruits are ready.

In this situation, the gardener is like the "Subject" in the Observer pattern, and you and your friends are the "Observers". The gardener (Subject) knows when the tree has ripe fruits, and he tells (notifies) you and your friends (Observers) when it's time to come to the park to eat them.

So, the Observer pattern is like having someone to watch something for you and tell you when something interesting happens. This way, you can go on and do other fun things without worrying about missing out.
*/

/*
In this pattern, the WeatherNotifier and ScoreNotifier act as subjects,
while the WeatherMobileObserver, WeatherEMailObserver, and ScoreEMailObserver act as observers.
*/

type Observable interface {
	Update()
}
type WeatherMobileObserver struct {
	notifier Notifier
	Mobile   int
}

func (o *WeatherMobileObserver) Update() {
	value, _ := o.notifier.GetData().(string)
	// Update observer
	fmt.Printf("ID %d updated with weather: %v\n", o.Mobile, value)
}

type WeatherEMailObserver struct {
	notifier Notifier
	EMail    string
}

func (o *WeatherEMailObserver) Update() {
	value, _ := o.notifier.GetData().(string)
	// Update observer
	fmt.Printf("EMail %s updated with weather: %s\n", o.EMail, value)
}

type ScoreEMailObserver struct {
	notifier Notifier
	EMail    string
}

func (o *ScoreEMailObserver) Update() {
	value, _ := o.notifier.GetData().(int)
	// Update observer
	fmt.Printf("EMail %s updated with score: %d\n", o.EMail, value)
}

type Notifier interface {
	Add(Observable)
	Remove(Observable)
	Notify()
	GetData() any
}

type WeatherNotifier struct {
	observers map[Observable]bool
	data      string
}

func (w *WeatherNotifier) Add(o Observable) {
	// Add observer to list
	w.observers[o] = true
}

func (w *WeatherNotifier) Remove(o Observable) {
	// Remove observer from list
	delete(w.observers, o)
}

func (w *WeatherNotifier) Notify() {
	// Notify all observers
	for o := range w.observers {
		o.Update()
	}
}

func (w *WeatherNotifier) SetData(data any) {
	d, _ := data.(string)
	w.data = d
	w.Notify()
}

func (w *WeatherNotifier) GetData() any {
	return w.data
}

type ScoreNotifier struct {
	observers map[Observable]bool
	data      int
}

func (s *ScoreNotifier) Add(o Observable) {
	// Add observer to list
	s.observers[o] = true
}

func (s *ScoreNotifier) Remove(o Observable) {
	// Remove observer from list
	delete(s.observers, o)
}

func (s *ScoreNotifier) Notify() {
	// Notify all observers
	for o := range s.observers {
		o.Update()
	}
}

func (s *ScoreNotifier) SetData(data any) {
	d, _ := data.(int)
	s.data = d
	s.Notify()
}

func (s *ScoreNotifier) GetData() any {
	return s.data
}

func main() {
	weatherNotifier := WeatherNotifier{observers: make(map[Observable]bool)}
	observer1 := WeatherMobileObserver{notifier: &weatherNotifier, Mobile: 1}
	observer2 := WeatherEMailObserver{notifier: &weatherNotifier, EMail: "rp@g.com"}
	weatherNotifier.Add(&observer1)
	weatherNotifier.Add(&observer2)
	weatherNotifier.SetData("Sunny")

	scoreNotifier := ScoreNotifier{observers: make(map[Observable]bool)}
	observer3 := ScoreEMailObserver{notifier: &scoreNotifier, EMail: "gp@q.com"}
	observer4 := ScoreEMailObserver{notifier: &scoreNotifier, EMail: "abc@q.com"}
	scoreNotifier.Add(&observer3)
	scoreNotifier.Add(&observer4)
	scoreNotifier.SetData(100)
}
