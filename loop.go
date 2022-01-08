package main

import (
	"fmt"
)

func main() {
	fmt.Println("Enter loop count")
	var n int
  fmt.Scan(&n)
	for i := 0; i < n; i++ {
		fmt.Printf("%d\n", i)
	}
}
