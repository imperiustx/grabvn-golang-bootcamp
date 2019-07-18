package main

import (
	"bufio"
	"os"
	"strconv"
	"strings"
)

func Div(x, y int) int {
	if y == 0 {
		println("must not 0")
	}
	return x / y
}

func contains(arr []string, str string) bool {
	for _, a := range arr {
		if a == str {
			return true
		}
	}
	return false
}

func Arithmetic(text string) string {
	opertors := []string{"+", "-", "*", "/"}
	splitText := strings.Split(text, " ")
	num1, err := strconv.Atoi(splitText[0])
	if err != nil {
		println("num 1 not int")
		return "there's a error"
	}
	num2, err := strconv.Atoi(splitText[2])
	if err != nil {
		println("num 2 not int")
		return "there's a error"
	}

	operator := splitText[1]
	c := contains(opertors, operator)
	if c == false {
		println("not a operator")
		return "Not a operator"
	}
	var result int

	switch operator {
	case "+":
		result = num1 + num2
	case "-":
		result = num1 - num2
	case "*":
		result = num1 * num2
	case "/":
		result = Div(num1, num2)
	}
	return text + " = " + strconv.Itoa(result)
}

func main() {
	reader := bufio.NewScanner(os.Stdin)
	reader.Scan()
	text := reader.Text()
	println(Arithmetic(text))
}
