package main

import "fmt"

func main() {
	a := []int{1, 2, 3, 4, 5}
	b := a[1:3]

	b[0] = 99
	b = append(b, 100)

	fmt.Println(a)
	fmt.Println(b)
}
