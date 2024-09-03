package main

import (
	"fmt"
	"strconv"
	"strings"
)

func main() {
	fmt.Println(Solution([]int{-6, -3, -2, -1, 0, 1, 3, 4, 5, 7, 8, 9, 10, 11, 14, 15, 17, 18, 19, 20}))
}

func Solution(arr []int) string {
	list := append(arr, 0)
	res := []string{}
	streak := 0
	for i := 0; i < len(list)-1; i++ {
		if list[i] == list[i+1]-1 {
			streak++
		} else {
			if streak > 1 {
				res = append(res, fmt.Sprintf("%v-%v", list[i-streak], list[i]))
			} else {
				if streak > 0 {
					res = append(res, strconv.Itoa(list[i-1]))
				}
				res = append(res, strconv.Itoa(list[i]))
			}
			streak = 0
		}
	}
	return strings.Join(res, ",")
}

func MoveZeros(arr []int) []int {
	zeroesCounter := 0
	res := []int{}
	for _, el := range arr {
		if el == 0 {
			zeroesCounter++
		} else {
			res = append(res, el)
		}
	}
	for i := 0; i < zeroesCounter; i++ {
		res = append(res, 0)
	}
	return res
}

func Snail(snaipMap [][]int) []int {
	res := []int{}
	if len(snaipMap) == 0 {
		return res
	}
	if len(snaipMap) == 1 {
		return append(make([]int, 0), snaipMap[0]...)
	}

	concatedSnaip := make([][]int, len(snaipMap)-2)
	for i := range concatedSnaip {
		concatedSnaip[i] = make([]int, len(snaipMap)-2)
	}

	for i := 0; i < len(snaipMap)-2; i++ {
		for j := 0; j < len(snaipMap)-2; j++ {
			concatedSnaip[i][j] = snaipMap[i+1][j+1]
		}
	}

	for i := 0; i < len(snaipMap[0]); i++ {
		res = append(res, snaipMap[0][i])
	}
	for i := 1; i < len(snaipMap); i++ {
		res = append(res, snaipMap[i][len(snaipMap)-1])
	}
	for i := len(snaipMap) - 2; i >= 0; i-- {
		res = append(res, snaipMap[len(snaipMap)-1][i])
	}
	for i := len(snaipMap) - 2; i > 0; i-- {
		res = append(res, snaipMap[i][0])
	}

	res = append(res, Snail(concatedSnaip)...)

	fmt.Println(concatedSnaip)
	return res
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
	if len(sentence) == 0 {
		return ""
	}
	var res []string
	wordsMap := map[int]string{}

	for _, word := range strings.Split(sentence, " ") {
		letters := strings.Split(word, "")
		for _, letter := range letters {
			if conv, err := strconv.Atoi(letter); err == nil {
				wordsMap[conv-1] = word
			}
		}
	}

	for i := 0; i < len(wordsMap); i++ {
		res = append(res, wordsMap[i])
	}
	return strings.Join(res, " ")
}
