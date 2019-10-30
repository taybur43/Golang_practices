package main
import (
	"errors"
	//"fmt"
	//"log"
)
func division(x, y int) (int, error, error) {
	if y == 0 {
		return 0, nil, errors.New("cannot divide by zero")
	}
	if x%y != 0 {
		remainder := errors.New("chere is a remainder")
		return x / y, remainder, nil
	} 
	return x / y, nil, nil
	
}