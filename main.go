package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

var Roman = map[string]int{
	"I":  1,
	"IV": 4,
	"V":  5,
	"IX": 9,
	"X":  10,
}

func intToRoman(num int) string {
	var calc string

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

	for k, v := range first {
		if k < len(first)-1 && Roman[string(first[k+1])] > Roman[string(first[k])] {
			firstRom -= Roman[string(v)]
		} else {
			firstRom += Roman[string(v)]
		}
	}

	for k, v := range second {
		if k < len(first)-1 && Roman[string(first[k+1])] > Roman[string(first[k])] {
			secondRom -= Roman[string(v)]
		} else {
			secondRom += Roman[string(v)]
		}
	}

	if action == "-" && firstRom < secondRom {
		return "Error"
	}

	if firstRom > 0 && firstRom <= 10 && secondRom > 0 && secondRom <= 10 {
		switch action {
		case "+":
			firstRom = firstRom + secondRom
		case "-":
			firstRom = firstRom - secondRom
		case "/":
			firstRom = firstRom / secondRom
		case "*":
			firstRom = firstRom * secondRom
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
		return "Error"
	}
	action := string(arr[pos])

	first := arr[0:pos]
	second := arr[pos+1 : len(arr)-2]

	if checkRoman(first, second) {
		return calcRoman(first, second, action)
	} else {
		if len(first) == 0 || len(second) == 0 {
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
			return "Error"
		}
	}

	return calc
}

func main() {
	reader := bufio.NewReader(os.Stdin)
	for {
		str, _ := reader.ReadString('\n')

		fmt.Print(calculate(str), "\n")

	}

}
