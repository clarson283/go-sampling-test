package main

import (
	"fmt"
	"math/rand"
)

func main() {
	s := make([]int, 50)

	for i := range s {
		s[i] = rand.Intn(100)
	}

	n := make([]int, 10)
	x := &s

	for j := range n {
		n[j] = s[rand.Intn(len(s))]
	}

	fmt.Println(*x)
	fmt.Println(n)
}
