package main

import (
	"fmt"
	//"time"
)

func main(){
	//for i:=1;i<=10;i++{
	//	fmt.Print(i)
	//	fmt.Print("\n")
	//	time.Sleep(1*time.Second)
	//}

	a := []string{"香蕉","苹果","雪梨"}
	for key,value:=range a{
		fmt.Print("key:")
		fmt.Print(key)
		fmt.Print("value:"+value+"\n")
		fmt.Print("\n")
	}

}
