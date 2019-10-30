package main
import (
	"fmt"
	"log"
)

func main(){
	fmt.Println("Starting from here")
	test,rem,err := division(1,3)
    if err != nil {
		log.Fatal(err)
	}else{
		fmt.Println("The result is = ", test)
	}
	if rem != nil{
		fmt.Println(rem)
	}
}