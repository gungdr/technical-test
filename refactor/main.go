package main

import (
	"fmt"
	"refactor/refactor"
)

func main() {
	result := refactor.FindString("(foo)")
	fmt.Printf("Result : %s\n", result)
}
