package main

import (
	"fmt"
)

func main() {
	fmt.Println("Enter 1 for true, 0 for false")
	var n int
  fmt.Scan(&n)
	if (n == 1) {
		fmt.Println("True")
	} else if (n == 0) {
		fmt.Println("False")
	} else {
		fmt.Println("Invalid choice")
	}
}
