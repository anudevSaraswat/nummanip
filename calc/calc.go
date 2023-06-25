package calc

import (
	"errors"
)

func Add(args ...int) (error, int) {

	if len(args) < 2 {
		return errors.New("minimum two arguments are required"), 0
	} else {
		sum := 0
		for _, arg := range args {
			sum += arg
		}
		return nil, sum
	}

}
