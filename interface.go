package main

// type Notifier interface {
// 	Notify(message string)
// }

// type Email struct {
// 	address string
// }

// func (e Email) Notify(message string) {
// 	fmt.Printf("Sending email to %s: %s\n", e.address, message)
// }

// type Sms struct {
// 	phoneNumber int
// }

// func (s Sms) Notify(message string) {
// 	fmt.Printf("Sending SMS to %d: %s\n", s.phoneNumber, message)
// }

// func SendNotification(n Notifier, message string) {
// 	n.Notify(message)
// }

// func interfaceExample() {

// 	email := Email{
// 		address: "serampore",
// 	}

// 	sms := Sms{
// 		phoneNumber: 1234567890,
// 	}

// 	SendNotification(email, "Hello via Email!")
// 	SendNotification(sms, "Hello via SMS!")

// }

// A interface is a type that specifies a contract for methods that a type must implement. It allows for polymorphism, where different types can be treated uniformly if they implement the same interface.

// A type can implement multiple interfaces, and an interface can be implemented by any type that provides the required methods. This allows for flexible and reusable code.

// There is no keyword for implementing an interface in go. it happens implicitly.

// Type Assertions

type Event interface {
	GetType() string
}

type LoginEvent struct {
	User string
}

func (l LoginEvent) GetType() string {
	return "login"
}

type PurchaseEvent struct {
	User  string
	Item  string
	Price float64
}

func (p PurchaseEvent) GetType() string {
	return "purchase"
}

type LogoutEvent struct {
	User string
}

func (l LogoutEvent) GetType() string {
	return "logout"
}

// func HandleEvent(e Event) {
// 	fmt.Println("Event type:", e.GetType())

// 	// Use type assertion to access fields specific to PurchaseEvent
// 	if purchase, ok := e.(PurchaseEvent); ok {
// 		fmt.Printf("Purchase by %s: %s for $%.2f\n", purchase.User, purchase.Item, purchase.Price)
// 	} else if login, ok := e.(LoginEvent); ok {
// 		fmt.Printf("Login by %s\n", login.User)
// 	} else {
// 		fmt.Println("Unknown event type")
// 	}
// }

// func TypeAssertionExample() {
// 	HandleEvent(LoginEvent{User: "alice"})
// 	HandleEvent(PurchaseEvent{User: "bob", Item: "Book", Price: 12.99})
// }

// Type Switches

// func HandleEventWithSwitch(e Event) {
// 	switch ev := e.(type) {
// 	case LoginEvent:
// 		fmt.Printf("[Login] User: %s\n", ev.User)
// 	case PurchaseEvent:
// 		fmt.Printf("[Purchase] User: %s, Item: %s, Price: $%.2f\n", ev.User, ev.Item, ev.Price)
// 	case LogoutEvent:
// 		fmt.Printf("[Logout] User: %s\n", ev.User)
// 	default:
// 		fmt.Println("Unknown event type")
// 	}
// }

// func TypeSwitchExample() {
// 	events := []Event{
// 		LoginEvent{User: "alice"},
// 		PurchaseEvent{User: "bob", Item: "Book", Price: 12.99},
// 		LogoutEvent{User: "alice"},
// 	}

// 	for _, e := range events {
// 		HandleEventWithSwitch(e)
// 	}
// }

// interface are not class or struct, they are just a set of methods that a type must implement. They are used to define a contract for behavior without specifying how that behavior is implemented.
