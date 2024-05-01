package main

import (
	"bufio"
	"fmt"
	"os"
	"sort"
	"strconv"
	"strings"
)

const (
	ErrInvalidOperationFormat      = "Ошибка! Неверный формат операции"
	ErrMixOfRomanAndArabicNumerals = "Ошибка! Введены одновременно разные системы счисления"
	ErrRomanNegativeResult         = "Ошибка! В римской системе нет отрицательных чисел"
	ErrOperandOutOfRange           = "Ошибка! Значение выходит за рамки моей работы: %s"
	ErrInvalidOperator             = "Ошибка! Недопустимый математический оператор"
)

func main() {
	fmt.Println("Добро пожаловать в мой первый калькулятор на Go!")

	scanner := bufio.NewScanner(os.Stdin)

	for {
		fmt.Print("Ввод: ")
		scanner.Scan()
		input := scanner.Text()

		parts := strings.Split(input, " ")

		if len(parts) != 3 {
			panic(ErrInvalidOperationFormat)
		}

		operand1, operator, operand2 := parts[0], parts[1], parts[2]

		if isRomanNumber(operand1) && !isRomanNumber(operand2) || !isRomanNumber(operand1) && isRomanNumber(operand2) {
			panic(ErrMixOfRomanAndArabicNumerals)
		}

		if !isValidOperator(operator) {
			panic(ErrInvalidOperator)
		}

		a, err := parseOperand(operand1)
		if err != nil {
			fmt.Println(err)
			continue
		}

		b, err := parseOperand(operand2)
		if err != nil {
			fmt.Println(err)
			continue
		}

		var result int
		switch operator {
		case "+":
			result = a + b
		case "-":
			result = a - b
			if isRomanNumber(operand1) && result < 0 {
				panic(ErrRomanNegativeResult)
			}
		case "*":
			result = a * b
		case "/":
			result = a / b
		}

		if isRomanNumber(operand1) {
			if result < 1 || result > 100 {
				panic(fmt.Sprintf(ErrOperandOutOfRange, arabicToRoman(result)))
			}
			fmt.Println("Вывод:", arabicToRoman(result))
		} else {
			fmt.Println("Вывод:", result)
		}
	}
}

func parseOperand(operand string) (int, error) {
	if isRomanNumber(operand) {
		num, err := romanToArabic(operand)
		if err != nil {
			return 0, err
		}
		if num < 1 || num > 10 {
			panic(fmt.Sprintf(ErrOperandOutOfRange, operand))
		}
		return num, nil
	}

	num, err := strconv.Atoi(operand)
	if err != nil {
		return 0, fmt.Errorf("invalid operand: %s", operand)
	}

	if num < 1 || num > 10 {
		panic(fmt.Sprintf(ErrOperandOutOfRange, operand))
	}
	return num, nil
}

func isRomanNumber(operand string) bool {
	for _, char := range operand {
		if char != 'I' && char != 'V' && char != 'X' && char != 'L' && char != 'C' {
			return false
		}
	}
	return true
}

func romanToArabic(roman string) (int, error) {
	romanNumerals := map[rune]int{'I': 1, 'V': 5, 'X': 10, 'L': 50, 'C': 100}
	var result int
	var prevValue int
	for i := len(roman) - 1; i >= 0; i-- {
		value := romanNumerals[rune(roman[i])]
		if value < prevValue {
			result -= value
		} else {
			result += value
		}
		prevValue = value
	}
	return result, nil
}

func arabicToRoman(arabic int) string {
	romanNumerals := map[int]string{
		1: "I", 4: "IV", 5: "V", 9: "IX", 10: "X",
		40: "XL", 50: "L", 90: "XC", 100: "C",
	}

	var result strings.Builder
	var keys []int
	for key := range romanNumerals {
		keys = append(keys, key)
	}
	sort.Sort(sort.Reverse(sort.IntSlice(keys)))

	for _, num := range keys {
		for arabic >= num {
			result.WriteString(romanNumerals[num])
			arabic -= num
		}
	}
	return result.String()
}

func isValidOperator(operator string) bool {
	validOperators := map[string]bool{"+": true, "-": true, "*": true, "/": true}
	return validOperators[operator]
}
