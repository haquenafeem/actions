package main

import (
	"fmt"

	"github.com/haquenafeem/actions/calculator"
)

func main() {
	fmt.Println("3 + 4 =", calculator.Add(3, 4))
	fmt.Println("25 - 6 =", calculator.Subtract(25, 6))
	fmt.Println("15 * 6 =", calculator.Multiply(15, 6))
	fmt.Println("15 * 4 =", calculator.Multiply(15, 4))
	// xyz
}
