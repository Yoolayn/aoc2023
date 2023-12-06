package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
	"sync"
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

type answer struct {
	lhs int
	rhs int
	ok bool
}

func liner(line string, ch chan<- int, wg *sync.WaitGroup) {
	defer wg.Done()
	nums := parse(line)
	if nums.ok {
		ch <- nums.lhs * 10 + nums.rhs
	}
}

func parse(line string) answer {
	var acc string
	var answer answer
	for _, char := range line {
		if num, err := strconv.Atoi(string(char)); err == nil {
			acc = ""
			if answer.ok {
				answer.rhs = num
			} else {
				answer.lhs = num
				answer.rhs = num
				answer.ok = true
			}
		}
		acc += string(char)
		for k, v := range tokens {
			if strings.Contains(acc, k) {
				acc = ""
				if answer.ok {
					answer.rhs = v
				} else {
					answer.lhs = v
					answer.rhs = v
					answer.ok = true
				}
			}
		}
	}
	return answer
}

func main() {
	file, err := os.Open("data/puzzleinput.txt")
	if err != nil {
		panic(err)
	}
	defer file.Close()

	var wg sync.WaitGroup
	ch := make(chan int)

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		wg.Add(1)
		line := scanner.Text()
		go liner(line, ch, &wg)
	}
	if err := scanner.Err(); err != nil {
		panic(err)
	}

	go func() {
		wg.Wait()
		close(ch)
	}()

	var sum int
	for ans := range ch {
		sum += ans
	}
	fmt.Println(sum)
}
