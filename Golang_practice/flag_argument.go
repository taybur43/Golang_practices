package main
import (
	"fmt"
	"flag"
)

func argMan(){
	fmt.Println("Found call")
	minusO := flag.Bool("en_verb",false,"Please enable to get details message")
	minusC := flag.Bool("en_sub",false,"Subtraction enable")
	minusK := flag.Int ("inter",1,"enable to add")
	flag.Parse()
	 fmt.Println("en_verb",*minusO)
	 fmt.Println("en_sub",*minusC)
	 fmt.Println("inter",*minusK)
}