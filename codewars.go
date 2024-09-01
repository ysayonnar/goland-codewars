package main

import (
	"fmt"
	"strings"
)

func main() {
	fmt.Println(ToCamelCase("The_Stealth_Warrior"))
}

func ArrayDiff(a, b []int) []int {
	for aIndex, aElement := range a {
		for _, bElement := range b {
			if aElement == bElement {
				a[aIndex] = -3000000
			}
		}
	}

	resultArr := make([]int, 0, len(a))

	for _, aElement := range a {
		if aElement != -3000000 {
			resultArr = append(resultArr, aElement)
		}
	}
	return resultArr

}

func SpinWords(str string) string {
	words := strings.Split(str, " ")

	for i, word := range words {
		if len(word) < 5 {
			continue
		}
		letters := strings.Split(word, "")
		newLetters := make([]string, len(letters))
		for j := len(letters) - 1; j >= 0; j-- {
			newLetters = append(newLetters, letters[j])
		}
		words[i] = strings.Join(newLetters, "")
	}

	return strings.Join(words, " ")
}

func Zeros(n int) int {
	res := 0
	for n > 0 {
		n /= 5
		res += n
	}
	return res
}

func ToCamelCase(s string) string {
	if len(s) == 0 {
		return ""
	}
	letters := strings.Split(s, "")
	res := make([]string, len(s))

	for i := 0; i < len(s); i++ {
		letter := letters[i]
		if letter == "-" || letter == "_" {
			res = append(res, strings.ToUpper(letters[i+1]))
			i++
			continue
		} else {
			res = append(res, letter)
		}
	}

	return strings.Join(res, "")
}

func Parse(data string) []int {
	var res []int
	value := 0
	commands := strings.Split(data, "")

	for _, cmd := range commands {
		switch cmd {
		case "i":
			value++
		case "d":
			value--
		case "s":
			value *= value
		case "o":
			res = append(res, value)
		}
	}

	return res
}

func FindOutlier(integers []int) int {
	oddArr := []int{}
	evenArr := []int{}

	for _, num := range integers {
		if num%2 == 0 {
			evenArr = append(evenArr, num)
		}
		if num%2 != 0 {
			oddArr = append(oddArr, num)
		}
	}

	if len(oddArr) == 1 {
		return oddArr[0]
	}
	if len(evenArr) == 1 {
		return evenArr[0]
	}
	return 0
}

func Order(sentence string) string {
	return sentence
}
