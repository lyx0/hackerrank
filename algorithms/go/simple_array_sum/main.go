package main

import "fmt"

type input []int

func main() {
	xi := input{1, 2, 3, 4, 10, 11}
	xi2 := input{6}

	s := sum(xi)
	fmt.Println(s)

	s = sum(xi2)
	fmt.Println(s)
}

func sum(in []int) int {
	ret := 0
	for _, v := range in {
		ret += v
		// println(i, v)
	}
	return ret
}
