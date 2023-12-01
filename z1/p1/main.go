package main

import (
	"fmt"
	"os"
	"strconv"
	"strings"
)

func isNumber(b byte) bool {
	return b >= 48 && b <= 57
}

func check(b string) (ints []int) {
	str := strings.Split(b, "\n")
	for idx, value := range str {
		if len(value) == 1 {
			str[idx] = value + value
		}
		if len(value) > 2 {
			first := value[0:1]
			last := value[len(value)-1:]
			str[idx] = first + last
		}
	}
	for _, value := range str[:len(str)-1] {
		val, err := strconv.Atoi(value)
		if err != nil {
			panic(err)
		}
		ints = append(ints, val)
	}
	return
}

func main() {
	data, err := os.ReadFile("data/puzzleinput.txt")
	if err != nil {
		panic(err)
	}

	// 10: newline
	// 48-57: 0-9

	var nums []byte
	for _, b := range data {
		if !(isNumber(b) || b == 10) {
			continue
		}
		nums = append(nums, b)
	}

	ints := check(string(nums))
	var sum int
	for _, v := range ints {
		sum += v
	}

	fmt.Println(sum)
}
