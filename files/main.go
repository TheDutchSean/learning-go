/*

	Chapter 4

*/

package main

import (
	"fmt"
	"io"
	"os"
)


func main(){


	// os.Create will create a file
	path := "./fromString.txt"
	file := createFile(path)
	length := writeToFile(file, "Hello from Go")
	fmt.Printf("Wrote a file with %v characters\n", length)
	// defer wait till evey operation has competed before continuing on
	defer file.Close() // dont forget to cole your files!!
	// before you want to read the file you have to make sure to close it and to make sure its closed we need to use the defer keyword
	defer fmt.Println("Text read from file:",string(readFile(path)))


}

func checkerror(err error){
	if err != nil {
		panic(err)
	}
}


func createFile(path string) *os.File {
	// os.Create will create a file
	file , err := os.Create("./fromString.txt")
	
	checkerror(err)

	return file
}

func writeToFile(file *os.File, content string) int {
	// io.WriteString will write an string to an file
	length, err := io.WriteString(file, content)

	checkerror(err)

	return length

}

func readFile(fileName string) []byte {

	data, err := os.ReadFile(fileName)
	checkerror(err)

	return data

}