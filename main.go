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
	"L": 50,
	"C": 100,
	"D": 500,
	"M": 1000,
}

func romanToInt(s string) int {
	sum := 0
	greatest := 0
	for i := len(s) - 1; i >= 0; i-- {
		letter := s[i]
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

func intToRoman(num int) string {
	var calc string

	if num <= 0 {
		swt = false
		return "Error"
	}

	for num >= 100 {
		num -= 100
		calc += "C"
	}

	for num >= 90 {
		num -= 90
		calc += "XC"
	}

	for num >= 50 {
		num -= 50
		calc += "L"
	}

	for num >= 40 {
		num -= 40
		calc += "XL"
	}

	for num >= 10 {
		num -= 10
		calc += "X"
	}

	for num >= 9 {
		num -= 9
		calc += "IX"
	}

	for num >= 5 {
		num -= 5
		calc += "V"
	}

	for num >= 4 {
		num -= 4
		calc += "IV"
	}

	for num >= 1 {
		num -= 1
		calc += "I"
	}

	return calc
}

func calcRoman(first, second, action string) string {
	var firstRom = 0
	var secondRom = 0

	firstRom = romanToInt(first)
	secondRom = romanToInt(second)

	if firstRom > 10 || secondRom > 10 {
		swt = false
		return "Error"
	}

	if action == "-" && firstRom < secondRom {
		swt = false
		return "Error"
	}

	if firstRom > 0 && firstRom <= 10 && secondRom > 0 && secondRom <= 10 {
		switch action {
		case "+":
			firstRom += secondRom
		case "-":
			firstRom -= secondRom
		case "/":
			firstRom /= secondRom
		case "*":
			firstRom *= secondRom
		}
	}

	return intToRoman(firstRom)
}

func checkRoman(first, second string) bool {
	for i := 0; i < len(first); i++ {
		if first[i] == 'I' || first[i] == 'X' || first[i] == 'V' {
			continue
		} else {
			return false
		}
	}

	for i := 0; i < len(second); i++ {
		if second[i] == 'I' || second[i] == 'X' || second[i] == 'V' {
			continue
		} else {
			return false
		}
	}

	return true
}

func calculate(arr string) string {
	var calc string
	pos := 0
	countOperands := 0

	for i := 0; i < len(arr); i++ {
		if arr[i] == '+' || arr[i] == '-' || arr[i] == '/' || arr[i] == '*' {
			pos += i
			countOperands++
		}
	}

	if countOperands != 1 {
		swt = false
		return "Error"
	}
	action := string(arr[pos])

	first := arr[0:pos]
	second := arr[pos+1 : len(arr)-2]

	first = strings.ReplaceAll(first, " ", "")
	second = strings.ReplaceAll(second, " ", "")

	if checkRoman(first, second) {
		return calcRoman(first, second, action)
	} else {
		if len(first) == 0 || len(second) == 0 {
			swt = false
			return "Error"
		}

		a, _ := strconv.Atoi(first)
		b, _ := strconv.Atoi(second)

		if a > 0 && a <= 10 && b > 0 && b <= 10 {
			switch action {
			case "+":
				calc = strconv.Itoa(a + b)
			case "-":
				calc = strconv.Itoa(a - b)
			case "/":
				calc = strconv.Itoa(a / b)
			case "*":
				calc = strconv.Itoa(a * b)
			}
		} else {
			swt = false
			return "Error"
		}
	}

	return calc
}

var swt = true

func main() {
	reader := bufio.NewReader(os.Stdin)
	for swt {
		str, _ := reader.ReadString('\n')
		fmt.Print(calculate(str), "\n")
	}

}
