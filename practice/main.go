package main

import (
	"fmt"
	"bufio"
	"os" // operating system ?
	"strconv" // for cstring conversions
	"strings" // for string functions
	"math" // used for math operations
	"time" // used for time functions
)

const aConst int = 64

func main(){

	
	// varibales()
	// inputs()
	// inputNum()
	// maths()
	dateAndtimes()

}

func varibales(){

	// go is explicit typeing this means that varibale need to be assinged a datatype

	var aString string = "This is Go"

	fmt.Println(aString)
	fmt.Printf("The varibale's type is %T\n", aString) // %T is a place holder for the varibale aString

	var anInteger int = 42

	fmt.Println(anInteger)
	fmt.Printf("The varibale's type is %T\n", anInteger) // %T is a place holder for the varibale anInteger

	var defaultInt int // a unasigned int will default to 0
	fmt.Println(defaultInt)

	// go can also use implicit typeing like JS
	var anotherString = "This is another string"

	fmt.Println(anotherString)
	fmt.Printf("The varibale's type is %T\n", anotherString) // %T is a place holder for the varibale anotherString

	var anotherInt = 53

	fmt.Println(anotherInt)
	fmt.Printf("The varibale's type is %T\n", anotherInt) // %T is a place holder for the varibale anotherInt

	// := can replace var a the beginning of an declaration and only works inside a functions scope SO NOT IN THE GLOBAL SCOPE
	myString := "This is also a string declared using :="

	fmt.Println(myString)
	fmt.Printf("The varibale's type is %T\n", myString) // %T is a place holder for the varibale myString


	fmt.Println(aConst)
	fmt.Printf("The varibale's type is %T\n", aConst) // %T is a place holder for the varibale aConst
}

func inputs(){

	reader := bufio.NewReader(os.Stdin)
	fmt.Print("Enter text: ")
	input, _ := reader.ReadString('\n')
	fmt.Println("You enterd:", input)

}

func inputNum(){

	reader := bufio.NewReader(os.Stdin)
	fmt.Print("Enter number: ")
	numInput, _ := reader.ReadString('\n')
	aFloat, err := strconv.ParseFloat(strings.TrimSpace(numInput), 64) 
	if(err != nil){
		fmt.Println(err)
	} else{
		fmt.Println("Value of number:",aFloat)
	}

}

func maths(){

	i1, i2, i3 := 12, 45 , 68
	intSum := i1 + i2 + i3

	fmt.Println("Integer sum:", intSum)

	f1, f2, f3 := 23.5, 65.1, 76.3
	floatSum := f1 + f2 + f3

	fmt.Println("Float sum:", floatSum)

	floatSum = math.Round(floatSum * 100) / 100 // multiply by 100 to round the values behinde the decimal point and devide by 100 to convert it back to the actual value
	fmt.Println("The sum is now:", floatSum)

	cicrcleRadius := 15.5
	circumference := cicrcleRadius * 2 * math.Pi
	fmt.Printf("Circumference: %.2f\n", circumference) // %.2f is a place holder for a float value with to digets behind the decimal

}

func dateAndtimes(){

	n := time.Now()
	fmt.Println("This program was run at ", n)

	t := time.Date(2009, time.November, 10, 23 ,0 ,0 ,0, time.UTC)
	fmt.Println("Go launched at ", t)
	fmt.Println(t.Format(time.ANSIC))

	parsedTime, _ := time.Parse(time.ANSIC, "Tue Nov 10 23:00:00 2009")
	fmt.Printf("The type of parsedTime is %T\n", parsedTime)
}