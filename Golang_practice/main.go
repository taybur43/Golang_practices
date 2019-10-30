package main
import (
	"fmt"
)

func main(){
	fmt.Println("Starting from here")
	test,err1,err2 := division(6,3)
	fmt.Println(test)
	fmt.Println(err1)
	fmt.Println(err2)
}