package calc

import (
	"errors"
	"fmt"
)

func Add(args ...int) error {
	if len(args) < 2 {
		return errors.New("minimum two arguments are required")
	} else {
		sum := 0
		for _, arg := range args {
			sum += arg
		}
		fmt.Println("Sum:", sum)
	}

	return nil

}
