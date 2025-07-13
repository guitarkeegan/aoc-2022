package main

import (
	"fmt"
	"iter"
	"os"
	"strings"
)

var dbg = func() func(format string, args ...any) {

	if os.Getenv("DEBUG") == "" {
		return func(string, ...any) {}
	}
	return func(format string, args ...any) {
		fmt.Printf("%s %v\n", format, args)
	}
}()

func day01(lines iter.Seq[string]) int {

	dbg("day01")

	var total int
	var letters = "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ"
	scores := make(map[rune]int)

	for i, letter := range letters {
		scores[letter] = i + 1
	}

	dbg("  scores: %v", scores)

	for line := range lines {
		firstHalf := line[:len(line)/2]
		secondHalf := line[len(line)/2:]

		for _, char := range firstHalf {
			if strings.Contains(secondHalf, string(char)) {
				total += scores[char]
				break
			}
		}

	}

	return total

}

func day02(lines []string) int {

	dbg("day02")

	var total int
	var letters = "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ"
	scores := make(map[rune]int)

	for i, letter := range letters {
		scores[letter] = i + 1
	}

	dbg("  scores: %v", scores)

	var l1, l2, l3 string

	for i, line := range lines {
		if i%3 == 0 {
			l1 = line
		} else if i%3 == 1 {
			l2 = line
		} else {
			l3 = line
			for _, ch := range l3 {
				if strings.Contains(l1, string(ch)) && strings.Contains(l2, string(ch)) {
					total += scores[ch]
					break
				}
			}
		}
	}

	return total

}

func main() {

	lines := strings.SplitSeq(input, "\n")

	linesSlice := strings.Split(input, "\n")

	fmt.Println(day01(lines))
	fmt.Println(day02(linesSlice))

}
