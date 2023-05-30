package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

var RomanNumerals = map[string]int{
	"I": 1,
	"V": 5,
	"X": 10,
	"i": 1,
	"v": 5,
	"x": 10,
}

func RomanToInt(str string) int {
	sum := 0
	greatest := 0
	for i := len(str) - 1; i >= 0; i-- {
		letter := str[i]
		num := RomanNumerals[string(letter)]
		if num >= greatest {
			greatest = num
			sum = sum + num
			continue
		}
		sum = sum - num
	}

	return sum
}

func IntToRoman(num int) string {
	var str string

	for num >= 100 {
		num -= 100
		str += "C"
	}

	for num >= 90 {
		num -= 90
		str += "XC"
	}

	for num >= 50 {
		num -= 50
		str += "L"
	}

	for num >= 40 {
		num -= 40
		str += "XL"
	}

	for num >= 10 {
		num -= 10
		str += "X"
	}

	for num >= 9 {
		num -= 9
		str += "IX"
	}

	for num >= 5 {
		num -= 5
		str += "V"
	}

	for num >= 4 {
		num -= 4
		str += "IV"
	}

	for num >= 1 {
		num -= 1
		str += "I"
	}

	return str
}

func checkRoman(str string) bool {
	for i := 0; i < len(str); i++ {
		if str[i] == 'I' || str[i] == 'V' || str[i] == 'X' || str[i] == 'i' || str[i] == 'v' || str[i] == 'x' {
			continue
		} else {
			return false
		}
	}
	return true
}

func calc(action string, first, second int) int {
	if first < 1 || first > 10 || second < 1 || second > 10 {
		swt = false
	}

	switch action {
	case "+":
		first += second
	case "-":
		first -= second
	case "*":
		first *= second
	case "/":
		first /= second
	}

	return first
}

func calculate(str string) string {
	pos := 0
	countOperands := 0
	var answer string

	for i := 0; i < len(str); i++ {
		if str[i] == '+' || str[i] == '-' || str[i] == '/' || str[i] == '*' {
			pos += i
			countOperands++
		}
	}

	if countOperands != 1 {
		swt = false
		return "Error"
	}

	action := string(str[pos])

	firstNum := str[0:pos]
	secondNum := str[pos+1 : len(str)]

	if !checkRoman(firstNum) && !checkRoman(secondNum) {
		first, _ := strconv.Atoi(firstNum)
		second, _ := strconv.Atoi(secondNum)

		answer = strconv.Itoa(calc(action, first, second))
	} else if checkRoman(firstNum) && checkRoman(secondNum) {
		if calc(action, RomanToInt(firstNum), RomanToInt(secondNum)) == 0 || calc(action, RomanToInt(firstNum), RomanToInt(secondNum)) < 0 {
			swt = false
			return "Error"
		}

		answer = IntToRoman(calc(action, RomanToInt(firstNum), RomanToInt(secondNum)))
	} else {
		swt = false
		return "Error"
	}

	return answer
}

var swt = true

func main() {
	reader := bufio.NewReader(os.Stdin)
	for swt {
		fmt.Println("Input num:")
		str, _ := reader.ReadString('\n')
		str = strings.TrimSpace(str)
		str = strings.ReplaceAll(str, " ", "")
		fmt.Print(calculate(str), "\n")
	}

}
