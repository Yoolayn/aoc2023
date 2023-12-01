package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

var tokens = map[string]int {
	"one":   1,
	"two":   2,
	"three": 3,
	"four":  4,
	"five":  5,
	"six":   6,
	"seven": 7,
	"eight": 8,
	"nine":  9,
}

func getLhs(s string) (string, bool) {
	var substr string
	for _, str := range s {
		substr += string(str)
		if i, err := strconv.Atoi(string(str)); err == nil {
			return strconv.Itoa(i), true
		}
		for k, v := range tokens {
			if strings.Contains(substr, k) {
				return strconv.Itoa(v), true
			}
		}
	}
	return strconv.Itoa(-1), false
}

func reverse(s string) string {
	runes := []rune(s)
	for i, j := 0, len(runes) - 1; i < j; i, j = i+1, j-1 {
		runes[i], runes[j] = runes[j], runes[i]
	}
	return string(runes)
}

func getRhs(s string) (string, bool) {
	var substr string
	for _, str := range s {
		substr += string(str)
		if i, err := strconv.Atoi(string(str)); err == nil {
			return strconv.Itoa(i), true
		}
		for k, v := range tokens {
			if strings.Contains(reverse(substr), k) {
				return strconv.Itoa(v), true
			}
		}
	}
	return strconv.Itoa(-1), false
}

func main() {
	file, err := os.Open("data/puzzleinput.txt")
	if err != nil {
		panic(err)
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)

	var sum int
	for scanner.Scan() {
		str := scanner.Text()
		lhs, lhs_ok := getLhs(str)
		rhs, rhs_ok := getRhs(reverse(str))

		if rhs_ok && lhs_ok {
			fmt.Println(str)
			num, err := strconv.Atoi(lhs + rhs)
			if err != nil {
				panic(err)
			}
			sum += num
		}
	}

	if err := scanner.Err(); err != nil {
		panic(err)
	}
	fmt.Println(sum)
}
