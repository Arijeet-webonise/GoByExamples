package main

import "fmt"

func mul(nums ...int) {
	fmt.Println(nums, " ")
	prod := 1
	for _, num := range nums {
		prod *= num
	}
	fmt.Println(prod)
}

func main() {
	mul(2, 6)
	mul(2, 5, 8)

	nums := []int{2, 3, 4, 5}
	mul(nums...)
}
