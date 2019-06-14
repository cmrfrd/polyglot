package main

import "fmt"

func main() {

	var _string string = ""

	for i := 1 ; i <=150 ; i++ {
		if i%3 == 0 {
			_string += "fizz"
		}
		if i%5 == 0 {
			_string += "buzz"
		}
		if i%7 == 0 {
			_string += "baz"
		}
	}
	

	fmt.Println(_string);
}
