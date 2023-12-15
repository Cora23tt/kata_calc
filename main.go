package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

var romanNumerals = map[string]int{
	"I":    1,
	"II":   2,
	"III":  3,
	"IV":   4,
	"V":    5,
	"VI":   6,
	"VII":  7,
	"VIII": 8,
	"IX":   9,
	"X":    10,
}

func isRoman(input string) bool {
	for roman := range romanNumerals {
		if input == roman {
			return true
		}
	}
	return false
}

func romanToArabic(input string) int {
	return romanNumerals[input]
}

func arabicToRoman(number int) string {
	if number <= 0 || number > 10 {
		return ""
	}

	romanNumerals := []struct {
		Value  int
		Symbol string
	}{
		{10, "X"},
		{9, "IX"},
		{8, "VIII"},
		{7, "VII"},
		{6, "VI"},
		{5, "V"},
		{4, "IV"},
		{3, "III"},
		{2, "II"},
		{1, "I"},
	}

	var result strings.Builder

	for _, numeral := range romanNumerals {
		for number >= numeral.Value {
			result.WriteString(numeral.Symbol)
			number -= numeral.Value
		}
	}

	return result.String()
}

func calculate(a, b int, operator string) int {
	switch operator {
	case "+":
		return a + b
	case "-":
		return a - b
	case "*":
		return a * b
	case "/":
		return a / b
	default:
		return 0
	}
}

func main() {
	reader := bufio.NewReader(os.Stdin)

	fmt.Println("Калькулятор. Введите выражение:")

	input, _ := reader.ReadString('\n')
	input = strings.TrimSpace(input)

	values := strings.Fields(input)
	if len(values) != 3 {
		fmt.Println("Ошибка: введено неверное выражение.")
		return
	}

	var a, b int
	var err error
	if isRoman(values[0]) && isRoman(values[2]) {
		a = romanToArabic(values[0])
		b = romanToArabic(values[2])
	} else {
		a, err = strconv.Atoi(values[0])
		if err != nil {
			fmt.Println("Ошибка: введено неверное число.")
			return
		}

		b, err = strconv.Atoi(values[2])
		if err != nil {
			fmt.Println("Ошибка: введено неверное число.")
			return
		}
	}

	if a < 1 || a > 10 || b < 1 || b > 10 {
		fmt.Println("Ошибка: числа должны быть от 1 до 10 включительно.")
		return
	}

	result := calculate(a, b, values[1])

	if isRoman(values[0]) && isRoman(values[2]) {
		if result <= 0 {
			fmt.Println("Ошибка: в римской системе нет неположительных чисел.")
			return
		}
		fmt.Printf("Результат: %s\n", arabicToRoman(result))
	} else {
		fmt.Printf("Результат: %d\n", result)
	}
}
