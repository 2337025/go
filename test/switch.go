package main

import "fmt"

func main() {
	switch 3 {
		case 1:
			fmt.Print("1")
		case 2:
			fmt.Print("2")
		default:
			fmt.Print("其他")
	}
}
