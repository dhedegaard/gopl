package main

import (
	"fmt"
	"time"
)

func main() {
	before := time.Now()

	for i := 0; i < 10000000; i++ {
		i = i
	}

	fmt.Println("Increment loop:", time.Now().Sub(before))

	rang := []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}

	before = time.Now()

	for i := 0; i < 1000000; i++ {
		for _, a := range rang {
			a = a
		}
	}

	fmt.Println("Range loop:", time.Now().Sub(before))
}
