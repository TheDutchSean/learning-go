/*

	Chapter 4

*/

package main

import (
	"fmt"
	"strconv"
	"strings"
)

func main(){

	// functions();
	// poodle := Dog{"Poodle", 12, "Woof"}
	// fmt.Println(poodle)
	// poodle.Speak()
	// poodle.Sound = "Arf!"
	// poodle.Speak()
	challenge()


}




func functions(){
	doSometing()
	sum := parameters(5, 8)
	fmt.Println("The sums is",sum)
	fmt.Println("The total is",addAllvalues(4, 7, 8, 9))
	multiSum, multiCount := returnMoreValues(4, 7, 8, 9)
	fmt.Println("The sums is",multiSum)
	fmt.Println("The length is",multiCount)
}

func doSometing(){
	fmt.Println("Doing something")
}

func parameters(value1 int, value2 int) int {
	return value1 + value2
}

// by using ... after a declaration you can define  a arbitray number of properties of teh same type
func addAllvalues(values...int) int {

	total := 0
	for _, value := range values {
		total += value
	}

	return total
}

// returning more than 1 value can de done by declaring the values in the output

func returnMoreValues(values...int) (int, int) {

	total := 0
	for _, value := range values {
		total += value
	}

	return total, len(values)
}


type Dog struct {
	// by using a Uppercase starting letter a variable becomes part of the global scope "its exported"
	Breed string
	Weight int
	Sound string
}

// methodes are part of structs (d Dog) define which struc will be the  methode parent
func (d Dog) Speak(){
	fmt.Println(d.Sound);
}


// methodes are part of structs (d Dog) define which struc will be the  methode parent
func (d Dog) SpeakThreeTimes(){ 
	d.Sound = fmt.Sprintf("%v %v %v",d.Sound, d.Sound, d.Sound);
	fmt.Println(d.Sound);
}

func challenge(){

	value1 := "100.0"
	value2 := "2"
	operation := "/"
	result := calculate(value1, value2, operation)
	fmt.Println(result);
}

// calculate() returns the result of the requested operation.
func calculate(input1 string, input2 string, operation string) float64 {
    // Your code goes here.
	value1 := convertInputToValue(input1)
	value2 := convertInputToValue(input2)

	if(operation == "+"){
		return addValues(value1, value2)
	} else if(operation == "-"){
		return subtractValues(value1, value2)
	} else if(operation == "*"){
		return multiplyValues(value1, value2)
	} else if(operation == "/"){
		return divideValues(value1, value2)
	}

	return 0
}

func convertInputToValue(input string) (float64){

	value, err := strconv.ParseFloat(strings.TrimSpace(input), 64)

	if(err != nil){
		return 0
	}
	
	return value
}

func addValues(value1, value2 float64) float64 {
	return value1 + value2
}

func subtractValues(value1, value2 float64) float64 {
	return value1 - value2
}

func multiplyValues(value1, value2 float64) float64 {
	return value1 * value2
}

func divideValues(value1, value2 float64) float64 {
	return value1 / value2
}