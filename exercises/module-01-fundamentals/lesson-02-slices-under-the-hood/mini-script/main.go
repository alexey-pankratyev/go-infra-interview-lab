package main

import "fmt"

// safeAppend returns a new slice with all elements of src appended to a copy of dst.
// The original dst must not be modified.
func safeAppend(dst, src []int) []int {
	// TODO: implement me.
	return nil
}

// chunk splits s into consecutive chunks of the given size.
// Each chunk must be an independent copy — modifying one must not affect others or s.
func chunk(s []int, size int) [][]int {
	// TODO: implement me.
	return nil
}

func main() {
	dst := []int{1, 2, 3}
	src := []int{4, 5}

	result := safeAppend(dst, src)
	fmt.Println("dst after safeAppend:", dst)
	fmt.Println("result of safeAppend:", result)

	s := []int{1, 2, 3, 4, 5}
	chunks := chunk(s, 2)
	// mutate first chunk — must not affect s
	if len(chunks) > 0 && len(chunks[0]) > 0 {
		chunks[0][0] = 99
	}
	fmt.Println("chunks:", chunks)
	fmt.Println("original after chunk mutation:", s)
}
