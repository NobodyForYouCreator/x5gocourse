package main

import "fmt"

func stat(shop []int) (int, int, int, float64) {
	n := len(shop)
	if n == 0 {
		return 0, 0, 0, 0.0
	}
	mn := 1000000
	mx := -1
	sm := 0

	for _, i := range shop {
		mn = min(mn, i)
		mx = max(mx, i)
		sm += i
	}

	return mn, mx, sm, float64(sm) / float64(n)
}

func main() {
	shop := [...]int{}
	fmt.Println(stat(shop[:]))
}
