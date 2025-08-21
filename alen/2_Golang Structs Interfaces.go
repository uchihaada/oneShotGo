package main

import (
	"encoding/json"
	"fmt"
)

// Golang Practice Set 4: Structs & Interfaces

type student struct {
	age  int
	name string
}

func (s *student) setAge(n int) {
	s.age = n
}

func (s student) getAge() {
	fmt.Println(s.age)
}

type human struct {
	eligible bool
	student
}

type animal interface {
	call() string
}

func callfunc(a animal) {
	fmt.Println(a.call())
}

type dog struct {
}

func (d *dog) call() string {
	return fmt.Sprintln("dog barks")
}

type bird struct {
}

func (b *bird) call() string {
	return fmt.Sprintln("bird cookos!")

}

type oneint interface {
	addOne() int
}

func add1(o oneint) {
	fmt.Println(o.addOne())
}

type twoint interface {
	addtwo() int
}

func add2(o twoint) {
	fmt.Println(o.addtwo())
}

type thirdint interface {
	oneint
	twoint
}

func add3(o thirdint) {
	fmt.Println("Using thirdint (embedded) interface:")
	fmt.Printf("  - addOne: %d\n", o.addOne())
	fmt.Printf("  - addtwo: %d\n", o.addtwo())
}

type interfaceembedding struct {
	a int
}

func (i *interfaceembedding) addOne() int {
	return i.a + 1
}
func (i *interfaceembedding) addtwo() int {
	return i.a + 2
}

type Address struct {
	city    string
	pincode int
}
type anonm struct {
	string
	int
	bool
	Address
}

type design interface {
	color() string
}

func col(c design) {
	fmt.Println(c.color())
}

type base struct {
}

func (b *base) color() string {
	return fmt.Sprintln("color is white")
}

type override struct {
	base
}

func (o *override) color() string {
	return "color is black"
}

type Builder struct {
	value string
	num   int
}

func (b *Builder) setNum() int {
	b.num = 2
	return b.num
}
func (b Builder) setNumval() int {
	b.num = 3
	return b.num
}
func (b *Builder) setName(str string) *Builder {
	b.value = "name :" + str
	return b
}

func (b *Builder) setAge(age int) *Builder {
	b.value += fmt.Sprintf("age: %d", age)
	return b
}

func (b *Builder) Build() string {
	return b.value
}

type Marshal struct {
	A int    `json:"id"`
	B string `json:"name"`
	C bool   `json:"eligible"`
}

func StructsInterfaces() {

	// 1. Basic struct

	s := student{
		age:  1,
		name: "rup",
	}

	fmt.Println(s.age, s.name)

	// 2. Struct with method
	s.setAge(5)
	fmt.Println(s.age)

	// 3. Embedded struct
	s1 := human{
		eligible: true,
		student: student{
			age:  19,
			name: "abhijit",
		},
	}
	fmt.Println(s1.eligible, s1.age, s1.name)

	// 4. Interface implementation
	d := dog{}
	callfunc(&d)

	// 5. Interface with multiple implementations
	d1 := dog{}
	callfunc(&d1)
	b := bird{}
	callfunc(&b)

	// 6. Type assertion
	a := 2
	b1 := "ada"
	assert(a)
	assert(b1)

	// 7. Type switch
	switchres(a)
	switchres(b1)

	// 8. Interface as parameter
	// this is same as assertion and switch example where interface is passed as type independent variable

	// 9. Pointer receiver method
	s2 := student{
		age:  23,
		name: "golu",
	}
	s2.setAge(24)
	s2.getAge()

	// 10. Interface embedding
	n := interfaceembedding{
		a: 2,
	}
	//calling directly
	fmt.Println(n.a)
	fmt.Println(n.addOne())
	fmt.Println(n.addtwo())
	//calling through individual interface
	add1(&n)
	add2(&n)
	//calling through embedded
	add3(&n)

	// 11. Struct literal
	S3 := student{20, "ADA"}
	fmt.Println(S3.age, S3.name)
	s4 := student{
		name: "ABHIJIT",
	}
	fmt.Println(s4.name)
	s5 := student{}
	fmt.Println(s5.name, s5.age)
	s6 := &student{age: 30, name: "hanuman"}
	fmt.Println(s6.age, s6.name)

	// 12. Anonymous field
	an := anonm{
		"ada",
		32,
		true,
		Address{
			"serampore",
			712202,
		},
	}

	fmt.Println(an.string)
	fmt.Println(an.int)
	fmt.Println(an.bool)

	fmt.Println(an.pincode)
	fmt.Println(an.city)

	// 13. Method overriding with interface
	ba := base{}
	col(&ba)
	ov := override{}
	col(&ov)
	// Use the embedded base field to avoid unused field error
	_ = ov.base

	// 14. Interface nil check
	res := check()
	fmt.Println(res == nil)
	res = check2()
	fmt.Println(res == nil)

	// 15. Method chaining
	build := Builder{}
	fmt.Println(build.setName("ada").setAge(23).Build())

	// 16. Value vs pointer receiver
	newb := Builder{"s", 10}
	fmt.Println(newb.num)
	fmt.Println(newb.setNum())
	fmt.Println(newb.num)
	fmt.Println(newb.setNumval())
	fmt.Println(newb.num)

	// 17. JSON marshal using struct
	ms := Marshal{
		A: 2,
		B: "df",
		C: true,
	}
	jsondaata, err := json.Marshal(ms)

	if err == nil {
		fmt.Println(string(jsondaata))
	}

	// 18. Interface nil trap

	// 19. Interface slice

	var anim []animal

	anim = append(anim, &d)
	anim = append(anim, &b)

	for _, r := range anim {
		callfunc(r)
	}

	// 20. Struct equality

	// data types that are non-comparable if inside of a struct then they
	// are non-comparable
	// All corresponding fields have equal values
	// Same struct type (field names, types, and order matter)
}

func check() interface{} {
	var i interface{}
	return i

}

func check2() interface{} {
	var p *animal = nil
	var i interface{}
	i = p
	return i
}
func assert(data interface{}) {
	if str, ok := data.(string); ok {
		fmt.Println(str)
	}
	if int, ok := data.(int); ok {
		fmt.Println(int)
	}

}

func switchres(data interface{}) {
	switch v := data.(type) {
	case int:
		fmt.Println(v)
	case string:
		fmt.Println(v)
	}
}
