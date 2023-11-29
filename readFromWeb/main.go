/*

	Chapter 4

*/

package main

import (
	"fmt"
	"net/http"
	"io"
	"encoding/json"
	"strings"
)

const url = "http://services.explorecalifornia.org/json/tours.php"

func main(){



	client := http.Client{}
	req , err := http.NewRequest("GET", url, nil)
	checkerror(err)

	req.Header = http.Header{
		"Host": {"services.explorecalifornia.org"},
		`User-Agent`: {`Go/2.1 package-http`},
	}

	response , err := client.Do(req)
	checkerror(err)

	defer response.Body.Close()

	bytes, err := io.ReadAll(response.Body)
	checkerror(err)

	content := string(bytes)

	// fmt.Println(content)
	tours := toursFromJson(content)
	for _, tour := range tours {
		fmt.Println(tour.Name)
	} 

	challenge()

}

func checkerror(err error){
	if err != nil {
		panic(err)
	}
}


type Tour struct {
	Name, Price string
}

func toursFromJson(content string) []Tour {

	tours := make([]Tour, 0, 20)
	
	decoder := json.NewDecoder(strings.NewReader(content))
	_, err := decoder.Token()
	checkerror(err)

	var tour Tour

	for decoder.More() {
		err := decoder.Decode(&tour)
		checkerror(err)
		tours = append(tours, tour)
	}

	return tours

}

// challange

type cartItem struct {
	Name  string `json:"name"`
	Price float64 `json:"price"`
	Quantity int `json:"quantity"`
}

func challenge(){
	
	jsonString := 
	`[{"name":"apple","price":2.99,"quantity":3},
	{"name":"orange","price":1.50,"quantity":8},
	{"name":"banana","price":0.49,"quantity":12}]`

	result := getCartFromJson(jsonString)
	fmt.Println(result)
	fmt.Println(getCartFromJsonV2(jsonString))
}

	// getCartFromJson() returns a slice containing cartItem objects.
	func getCartFromJson(jsonString string) []cartItem {
		
		var cart []cartItem
	
		decoder := json.NewDecoder(strings.NewReader(jsonString))
		_, err := decoder.Token()
		checkerror(err)
		
		var item cartItem

		for decoder.More() {
			err := decoder.Decode(&item)
			checkerror(err)
			cart = append(cart, item)
		}
	
		// Your code goes here.
		return cart
	}


	func getCartFromJsonV2(jsonString string) []cartItem {
		
		var cart []cartItem
	
		err := json.Unmarshal([]byte(jsonString), &cart)
		checkerror(err)
		
		return cart
	}