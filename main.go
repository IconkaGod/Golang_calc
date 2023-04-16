package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

var Roman = map[string]int{
	"C":  100,
	"L":  50,
	"X":  10,
	"IX": 9,
	"V":  5,
	"IV": 4,
	"I":  1,
}

var Romana = map[int]string{
	100: "C",
	50:  "L",
	10:  "X",
	9:   "IX",
	5:   "V",
	4:   "IV",
	1:   "I",
}

func intToRoman(num int) string {
	var calc string

	for i, j := range Romana {
		for num >= i {
			num -= i
			calc += j
		}
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
