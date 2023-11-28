/*

	Chapter 4

*/

package main

import (
	"fmt"
	"math/rand"
	"time"
	"math"
)

func main(){

	// conditional()
	// switchStatement()
	// loops()
	// challenge()
	var out int
for j:=0; j < 20; j++ {
  out=j*j + out
  if out > 12 {goto theEnd }  }
theEnd: fmt.Println(out)
}

func conditional(){

	theAnswer := 42
	var result string

	if theAnswer < 0 {
		result = "Less then zero"
	} else if theAnswer == 0 {
		result = "Equal then zero"
	} else{
		result = "Greater than zero"
	}

	fmt.Println(result)

	// here the variable is declared in the if statement
	if x := -42; x < 0 {
		result = "Less then zero"
	} else if x == 0 {
		result = "Equal then zero"
	} else{
		result = "Greater than zero"
	}

	fmt.Println(result)
}

func switchStatement(){
	// https://pkg.go.dev/math/rand@master#Seed
	source := rand.NewSource(time.Now().Unix())
	rng := rand.New(source)
	dow := rng.Intn(7) + 1
	fmt.Println("Day", dow)

	var result string
	switch dow {
		case 1: 
			result = "It's Sunday!"
			// fallthrough will also execute the cases below
		case 2: 
			result = "It's Monday!"
		case 3: 
			result = "It's Tuesday!"
		case 4: 
			result = "It's Wednesday!"
		case 5: 
			result = "It's Thursday!"
		case 6: 
			result = "It's Friday!"
		case 7: 
			result = "It's Saturday!"
		default:
			result = "It's a Funky day"
	}

	fmt.Println(result)

}


func loops(){

	colors := [] string {"Red", "Green", "Blue"}
	fmt.Println(colors)

	for i := 0; i < len(colors); i++{
		fmt.Println(colors[i])
	}

	for i := range colors {
		fmt.Println(colors[i])
	}

	for index, color := range colors{
		fmt.Printf("#%v:%v\n", index, color)
	}

	value := 1
	for value < 10 {
		fmt.Println("Value:", value)
		value++
	} 

	// break and continue can be used to
	sum := 1
	for sum < 10 {
		sum += sum
		fmt.Println("Sum:", sum)
		if sum > 200 {
			goto theEnd
		}
	} 

	// this is a label to jump to
	theEnd : fmt.Println("End of program")

}


type cartItem struct{
	name string
	price float64
	quantity int
}

func challenge(){



	var cart []cartItem
	var apples = cartItem{"apple", 2.99, 3}
	var oranges = cartItem{"orange", 1.50, 8}
	var bananas = cartItem{"banana", .49, 12}
	cart = append(cart, apples, oranges, bananas)
	
	
	result := calculateTotal(cart)

	fmt.Println(result)

}

func calculateTotal(cart []cartItem) float64 {
	
	priceTotal := 0.00

	for _, item := range cart{
		itemPrice := item.price * float64(item.quantity)
		priceTotal += itemPrice
	}
	
	priceTotal = math.Round(priceTotal * 100) / 100
	return priceTotal
}