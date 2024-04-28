package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

const (
	ErrInvalidOperationFormat      = "Ошибка! Неверный формат операции"
	ErrMixOfRomanAndArabicNumerals = "Ошибка! Введены одновременно разные системы счисления"
	ErrRomanNegativeResult         = "Ошибка! В римской системе нет отрицательных чисел"
	ErrOperandOutOfRange           = "Ошибка! Значение выходит за рамки моей работы: %s"
)

func main() {
	fmt.Println("Добро пожаловать в мой первый калькулятор на Go!")

	scanner := bufio.NewScanner(os.Stdin)

	for {
		fmt.Print("Ввод: ")
		scanner.Scan()
		input := scanner.Text()

		parts := strings.Split(input, " ")
		if len(parts) != 3 || len(parts) == 1 && strings.ContainsAny(parts[0], "+-*/") {
			panic(ErrInvalidOperationFormat)
		}

		operand1, operator, operand2 := parts[0], parts[1], parts[2]

		if isRomanNumber(operand1) && !isRomanNumber(operand2) || !isRomanNumber(operand1) && isRomanNumber(operand2) {
			panic(ErrMixOfRomanAndArabicNumerals)
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
			fmt.Println("Вывод:", arabicToRoman(result))
		} else {
			fmt.Println("Вывод:", result)
		}
	}
}

func parseOperand(operand string) (int, error) {
	if isRomanNumber(operand) {
		return romanToArabic(operand)
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
		if char != 'I' && char != 'V' && char != 'X' {
			return false
		}
	}
	return true
}

func romanToArabic(roman string) (int, error) {
	romanNumerals := map[rune]int{'I': 1, 'V': 5, 'X': 10}
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
		1: "I", 2: "II", 3: "III", 4: "IV", 5: "V",
		6: "VI", 7: "VII", 8: "VIII", 9: "IX", 10: "X",
	}

	var result strings.Builder
	for arabic > 0 {
		for num := 10; num > 0; num-- {
			if arabic >= num {
				result.WriteString(romanNumerals[num])
				arabic -= num
				break
			}
		}
	}
	return result.String()
}
