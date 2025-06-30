package main

// import "fmt"

// func mapExample() {

// 	// declare a map
// 	ages := make(map[int]string)

// 	// put values while declaring the map
// 	declAges := map[int]string{
// 		1: "golu",
// 		2: "ada",
// 		3: "urmila",
// 	}

// 	// get the length of the map
// 	fmt.Println("lenght of ages :", len(ages))
// 	fmt.Println("lenght of declAges :", len(declAges))

// 	// put values on the already created map
// 	ages[2] = "golu"

// 	// print values of the map
// 	fmt.Println("values", ages)
// 	fmt.Println("name with id 1:", declAges[1])

// 	// delete a value
// 	delete(declAges, 1)
// 	val, ok := declAges[1]
// 	if !ok {
// 		fmt.Print("success")
// 	} else {
// 		fmt.Print("not deleted properly", val)
// 	}

// }

// type of the key in map is all that is comparable

// it is simpler to use struct as key than nested maps

// one key to one value

// if associated value doesnt exist the key returns 0 value as it is a safe operation to access the key

// maps hold references
