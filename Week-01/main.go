package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

// Nhập vào biểu thức, tính giá trị biểu thức => làm cách thường và kí pháp balan

func main() {
	var scanner = bufio.NewScanner(os.Stdin)

	fmt.Print("> Enter expression: ")
	for scanner.Scan() {
		var expression = scanner.Text()

		eval(expression)

		fmt.Print("> Enter expression: ")
	}
}

// Reserve Polish Notation
func rpnEval(expression string) {
	var stack []string

	fmt.Println(stack)
}

func eval(expression string) {
	var chunks []string = strings.Split(expression, " ")

	if len(chunks) != 3 {
		fmt.Println("Invalid expression")
		return
	}

	var x, error_x = strconv.ParseFloat(chunks[0], 10)
	var y, error_y = strconv.ParseFloat(chunks[2], 10)

	if error_x != nil {
		fmt.Println("Invalid x - ", error_x)
		return
	}

	if error_y != nil {
		fmt.Println("Invalid y - ", error_y)
		return
	}

	switch chunks[1] {
	case "+":
		fmt.Printf("> x + y = %.2f\n", x+y)
	case "-":
		fmt.Printf("> x - y = %.2f\n", x-y)
	case "*":
		fmt.Printf("> x * y = %.2f\n", x*y)
	case "/":
		fmt.Printf("> x / y = %.2f\n", x/y)
	default:
		fmt.Println("Invalid operator")
	}
}
