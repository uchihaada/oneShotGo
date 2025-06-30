package main

// type User struct {
// 	name string
// 	age  int
// }

// type Message struct {
// 	sender    User
// 	recipient User
// 	content   string
// }

// var messages []Message

// func SendMessage(sender User, recipient User, content string) {

// 	msg := Message{
// 		sender:    sender,
// 		recipient: recipient,
// 		content:   content,
// 	}

// 	messages = append(messages, msg)
// }

// func GetMessageBySender(senderName string) []Message {

// 	var result []Message

// 	for _, msg := range messages {
// 		if msg.sender.name == senderName {
// 			result = append(result, msg)
// 		}
// 	}
// 	return result
// }
// func GetMessagesByRecipient(recipientName string) []Message {

// 	var result []Message

// 	for _, msg := range messages {
// 		if msg.recipient.name == recipientName {
// 			result = append(result, msg)
// 		}
// 	}
// 	return result
// }

// func msg() {

// 	alice := User{name: "Alice", age: 30}
// 	bob := User{name: "Bob", age: 25}
// 	ada := User{name: "Ada", age: 28}
// 	golu := User{name: "Golu", age: 22}

// 	SendMessage(alice, bob, "Hello Bob from alice!")
// 	SendMessage(ada, bob, "Hi Bob from ada!")
// 	SendMessage(golu, bob, "Hey Bob from golu!")

// 	SendMessage(ada, alice, "Hi Alice from ada!")
// 	SendMessage(ada, golu, "Hey Golu from ada!")
// 	SendMessage(ada, bob, "Hello Bob from ada!")

// 	fmt.Println("Messages sent by Ada:")
// 	for _, msg := range GetMessageBySender("Ada") {
// 		fmt.Printf("From: Ada To: %s, Content: %s\n", msg.recipient.name, msg.content)
// 	}

// 	fmt.Println("Messages received by Bob:")
// 	for _, msg := range GetMessagesByRecipient("Bob") {
// 		fmt.Printf("From: %s To: Bob, Content: %s\n", msg.sender.name, msg.content)
// 	}

// }

// Anonymous Structs

// func anonymousStructExample() {
// 	myStruct := struct {
// 		name string
// 		age  int
// 	}{
// 		name: "ada",
// 		age:  28,
// 	}

// 	fmt.Println("Anonymous Struct:")
// 	fmt.Printf("Name: %s, Age: %d\n", myStruct.name, myStruct.age)

// 	// Fixed nested anonymous struct:
// 	nestedStruct := struct {
// 		person struct {
// 			name string
// 			age  int
// 		}
// 		address struct {
// 			city  string
// 			state string
// 		}
// 	}{
// 		person: struct {
// 			name string
// 			age  int
// 		}{
// 			name: "Eve",
// 			age:  35,
// 		},
// 		address: struct {
// 			city  string
// 			state string
// 		}{
// 			city:  "Wonderland",
// 			state: "Dream",
// 		},
// 	}

// 	fmt.Println("Nested Anonymous Struct:")
// 	fmt.Printf("Person: %s, Age: %d, City: %s, State: %s\n",
// 		nestedStruct.person.name,
// 		nestedStruct.person.age,
// 		nestedStruct.address.city,
// 		nestedStruct.address.state)
// }

// Anonymous structs are useful for quick, one-off data structures without needing to define a full type.
// The only reason they are used when you need only one instance of a struct and don't need to reuse the type elsewhere in your code.

// Embeded Structs

// type car struct {
// 	make  string
// 	model string
// }

// type truck struct {
// 	car          // Embedding car struct
// 	capacity int // Additional field specific to truck
// }

// func embededStructExample() {

// 	trk := truck{
// 		capacity: 4,

// 		car: car{
// 			make:  "Ford",
// 			model: "F-150",
// 		},
// 	}

// 	fmt.Println("Embedded Struct:")
// 	fmt.Printf("Truck Make: %s, Model: %s, Capacity: %d tons\n",
// 		trk.make, trk.model, trk.capacity)
// }

// Embedding allows you to include one struct within another, enabling code reuse and creating more complex data structures.
// In this example, the `truck` struct embeds the `car` struct, allowing it to access the fields of `car` directly.
// This is useful for creating hierarchical relationships between data structures, where one struct can be a specialized version of another.
// It is also like oops inheritance, where the `truck` struct inherits the properties of the `car` struct.

// Methods for the structs

// type rect struct {
// 	height int
// 	width  int
// }

// func (r *rect) area() int {
// 	return r.height * r.width
// }

// func (r *rect) SetHeight(height int) {
// 	r.height = height
// }
// func Methods() {
// 	r := rect{
// 		height: 10,
// 		width:  5,
// 	}
// 	fmt.Printf("Area of rectangle: %d\n", r.area())
// 	r.SetHeight(0)
// 	fmt.Println("Height after SetHeight:", r.height)
// 	fmt.Printf("Area of rectangle after SetHeight: %d\n", r.area())
// }
