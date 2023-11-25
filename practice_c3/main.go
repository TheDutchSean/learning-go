/*

	the new() function does not initialize memory and you cannot add a key value pair to it

	m := new(map[string]int)
	m["theAnswer"] = 42
	fmt.Println(m)

	Will return 
	invalid operation: m["theAnswer"]
	(type *map[string]int) does not support indexing)
	

	make() should be used instead if you want to add data immediately to the object

	the make() function does initialize memory and can accept values

	m := make(map[string]int)
	m["theAnswer"] = 42
	fmt.Println(m)

	Will return
	map[theAnswer:42]

*/

package main

import (
	"fmt"
	"sort"
)

func main(){

	// pointer()
	// arrays()
	// slices()
	// maps()
	// structs()
	var ary = []int16{12, 7, 4, 67, 82}
ary = append(ary[2:len(ary)])
fmt.Println(ary[2])
	challenge()

}


func pointer(){
	anInt := 42
	var p = &anInt
	fmt.Println("VAlue of p:", *p)

	value1 := 42.13
	pointer1 := &value1
	fmt.Println("Value of p:", *pointer1)

	*pointer1 = *pointer1 / 31
	fmt.Println("Pointer 1:", *pointer1)
	fmt.Println("Value 1:", value1)

}


func arrays(){

	//arrays cant be used to order or modifyed durring run time
	// an array size must be declared
	var colors[3]string
	colors[0] = "Red"
	colors[1] = "Green"
	colors[2] = "Blue"
	fmt.Println(colors)
	fmt.Println(colors[0])

	var numbers = [5]int {5,3,1,2,4}
	fmt.Println(numbers)
	
	fmt.Println("Number of colors:", len(colors))
	fmt.Println("Number of numbers:", len(numbers))
}

func slices(){
	// an slices is an array just without its size declared
	var colors = [] string {"Red", "Green", "Blue"}
	fmt.Println(colors)

	colors = append(colors, "Purple")
	fmt.Println(colors)

	// to remove items from a slice you use:
	colors = append(colors[1:len(colors)])
	fmt.Println(colors)

	// using make([]int, 5 ,5) the slice is limited to 5 items leaving out the last number removes the limiter make([]int, 5)
	numbers := make([]int, 5)
	numbers[0] = 134
	numbers[1] = 72
	numbers[2] = 32
	numbers[3] = 12
	numbers[4] = 146
	fmt.Println(numbers)
	
	numbers = append(numbers, 236)
	fmt.Println(numbers)

	sort.Ints(numbers)
	fmt.Println(numbers)
}

func maps(){

	states := make(map[string]string)
	fmt.Println(states)
	states["WA"] = "Washington"
	states["OR"] = "Oregon"
	states["CA"] = "Califronia"
	fmt.Println(states)

	california := states["CA"]
	fmt.Println(california)

	delete(states, "OR")

	states["NY"] = "New York"
	fmt.Println(states)

	for key, value := range states {
		fmt.Printf("%v: %v\n", key, value)
	}

	keys := make([]string, len(states))
	i := 0
	for key := range states {
		keys[i] = key
		i++
	}

	fmt.Println(keys)
	// sort the keys
	sort.Strings(keys)
	fmt.Println(keys)


	for i := range keys {
		fmt.Println(states[keys[i]])
	}

}


func structs(){

	// structs can be seen as classes only like other langues the cannot inhert properties or methods the are all inviduals

	poodle := Dog{"Poodle", 12}
	fmt.Println(poodle)
	fmt.Printf("%v\n", poodle)
	fmt.Printf("Breed:%v\nWieght:%v\n", poodle.Breed, poodle.Weight)
	poodle.Weight = 10
	fmt.Printf("Breed:%v\nWieght:%v\n", poodle.Breed, poodle.Weight)
}

type Dog struct {
	// by using a Uppercase starting letter a variable becomes part of the global scope "its exported"
	Breed string
	Weight int
}



func challenge(){

	slice := []string{"apple", "banana", "orange"}
	
	


	result := ConvertToMap(slice)

	fmt.Println(result)



}

// Converts a slice of strings to a map object.
func ConvertToMap(items []string) map[string]float64 {

	var price float64 = 100 / float64(len(items))
	// Create a map object
	result := make(map[string]float64)
	for i := range items {
		result[items[i]] = price
	}
	// Your code goes here
	return result
}
