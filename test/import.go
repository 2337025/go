package main

import (
	"fmt"
	loovee "test/show2"
)

func init(){
	fmt.Print("main init\n")
}

func main(){
	loovee.Show2();

	fmt.Print("main main\n")
}
