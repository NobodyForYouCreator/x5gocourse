package main

import (
	"fmt"
)

func Category(i int) string {
	text := "err"
	if 0 <= i && i < 10_000 {
		text = "small"
	} else if 10_000 <= i && i <= 49_999 {
		text = "regular"
	} else if 50_000 <= i {
		text = "large"
	}

	return text
}

func main() {
	fmt.Println(Category(-10000))
}
