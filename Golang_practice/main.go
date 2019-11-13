package main
import (
	"fmt"
	"log"
	"github.com/rs/zerolog"
    
)

func main(){
	fmt.Println("Starting from here")
	zerolog.TimeFieldFormat = zerolog.TimeFormatUnix
    log.Print("hello world")
	test,rem,err := division(6,3)
    if err != nil {
		log.Fatal(err)
	}else{
		fmt.Println("The result is = ", test)
	}
	if rem != nil{
		fmt.Println(rem)
	}
	fmt.Println("Going to call .....")
	argMan()
}