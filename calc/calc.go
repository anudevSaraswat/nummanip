package calc

import(
  "fmt"
"errors"
)

func Add(args ...int) error {
	if len(args) < 0 {
return errors.New("minimum two arguments are required")
} else {
sum := 0
for _, arg := range args {
sum += arg
}
fmt.Println("Sum:", sum)
}

}
