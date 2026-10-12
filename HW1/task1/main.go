package main

import (
	"fmt"
	"strconv"
)

func FizzBuzz(i int) string {
	text := strconv.Itoa(i)

	if i%15 == 0 {
		text = "FizzBuzz"
	} else if i%3 == 0 {
		text = "Fizz"
	} else if i%5 == 0 {
		text = "Buzz"
	}

	return text
}

func main() {
	for i := 1; i <= 30; i++ {
		fmt.Println(FizzBuzz(i))
	}
}
