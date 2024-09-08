package main

import (
	"fmt"
	"math"
	"strconv"
	"strings"
)

func main() {
	fmt.Println(DecodeMorse(".... . -.--   .--- ..- -.. ."))
}

//26,8

func DecodeMorse(morseCode string) interface{} {
	morseWords := strings.Split(morseCode, "   ")
	convertedWords := []string{}

	for _, word := range morseWords {
		convertedWord := []string{}
		letters := strings.Split(word, " ")
		for _, letter := range letters {
			//convertedWord = append(convertedWord, MORSE_CODE[letter])
			convertedWord = append(convertedWord, letter)
		}
		convertedWords = append(convertedWords, strings.Join(convertedWord, ""))
	}

	res := []string{}
	for _, el := range convertedWords {
		if el == "" {
			continue
		} else {
			res = append(res, el)
		}
	}

	return strings.Join(res, " ")
}

func GetMinBase(n uint64) uint64 {
	var r uint64
	if n%2 == 0 {
		r = 1
	}
	if n%2 != 0 {
		r = 2
	}
	copuOfN := n

	for copuOfN >= r {
		ostatok := int(copuOfN % r)
		copuOfN /= r
		if copuOfN < r && copuOfN != 1 {
			r += 2
			copuOfN = n
			continue
		}
		if ostatok != 1 {
			r += 2
			copuOfN = n
			continue
		}
	}
	return r

}

func convertSingleRgb(value int) []int {
	if value < 0 {
		value = 0
	}
	if value > 255 {
		value = 255
	}
	var x int
	var y int
	x = value / 16
	y = value - 16*x
	return []int{x, y}
}

func RGB(r, g, b int) string {
	convertedValues := []int{}
	convertedValues = append(convertedValues, convertSingleRgb(r)...)
	convertedValues = append(convertedValues, convertSingleRgb(g)...)
	convertedValues = append(convertedValues, convertSingleRgb(b)...)
	res := []string{}

	for _, el := range convertedValues {
		switch el {
		case 10:
			res = append(res, "A")
		case 11:
			res = append(res, "B")
		case 12:
			res = append(res, "C")
		case 13:
			res = append(res, "D")
		case 14:
			res = append(res, "E")
		case 15:
			res = append(res, "F")
		default:
			res = append(res, strconv.Itoa(el))
		}
	}

	return strings.Join(res, "")
}

func SumOfIntervals(intervals [][2]int) int {
	sum := 0
	for _, inteval := range intervals {
		sum += inteval[len(inteval)-1] - inteval[0]
	}
	return sum
}

type dateType struct {
	name  string
	value int64
}

func FormatDuration(seconds int64) string {
	res := []string{}
	data := []dateType{}

	years := seconds / 31536000
	data = append(data, dateType{name: "years", value: years})
	seconds -= years * 31536000
	days := seconds / 86400
	data = append(data, dateType{name: "days", value: days})
	seconds -= days * 86400
	hours := seconds / 3600
	data = append(data, dateType{name: "hours", value: hours})
	seconds -= hours * 3600
	minutes := seconds / 60
	data = append(data, dateType{name: "minutes", value: minutes})
	seconds -= minutes * 60
	data = append(data, dateType{name: "seconds", value: seconds})

	realLength := len(data)

	for i := 0; i < len(data); i++ {
		if data[i].value == 0 {
			realLength--
		}
	}

	for i := 0; i < len(data); i++ {
		if data[i].value != 0 {
			if data[i].value == 1 {
				data[i].name = data[i].name[:len(data[i].name)-1]
			}
			if realLength == 1 {
				return fmt.Sprintf("%v %v", data[i].value, data[i].name)
				break
			}
			res = append(res, fmt.Sprintf("%v %v, ", data[i].value, data[i].name))
		}
	}
	lastEl := res[len(res)-1]
	res[len(res)-1] = lastEl[:len(lastEl)-2]

	if len(res) > 1 {
		res[len(res)-1] = fmt.Sprintf(" and %v", res[len(res)-1])
		prevEl := res[len(res)-2]
		res[len(res)-2] = res[len(res)-2][:len(prevEl)-2]
	}

	return strings.Join(res, "")
}

func HumanReadableTime(seconds int) string {
	res := []string{}

	hours := seconds / 3600
	if len(strconv.Itoa(hours)) == 1 {
		res = append(res, fmt.Sprintf("0%v", strconv.Itoa(hours)))
	} else {
		res = append(res, strconv.Itoa(hours))
	}

	seconds -= hours * 3600
	minutes := seconds / 60
	if len(strconv.Itoa(minutes)) == 1 {
		res = append(res, fmt.Sprintf("0%v", strconv.Itoa(minutes)))
	} else {
		res = append(res, strconv.Itoa(minutes))
	}

	seconds -= minutes * 60
	if len(strconv.Itoa(seconds)) == 1 {
		res = append(res, fmt.Sprintf("0%v", strconv.Itoa(seconds)))
	} else {
		res = append(res, strconv.Itoa(seconds))
	}
	return strings.Join(res, ":")
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

func ListSquared(m, n int) [][]int {
	res := [][]int{}

	for i := m; i <= n; i++ {
		squaredDivisors := []int{}
		for d := 1; d <= i; d++ {
			if i%d == 0 {
				squaredDivisors = append(squaredDivisors, i/d)
			}
		}
		sum := 0
		for _, el := range squaredDivisors {
			sum += el * el
		}
		root := math.Sqrt(float64(sum))
		if root == float64(int(root)) {
			res = append(res, []int{i, sum})
		}
	}

	return res
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
