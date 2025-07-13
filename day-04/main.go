package main

import (
	"fmt"
	"iter"
	"strconv"
	"strings"
)

func countPairs(lines iter.Seq[string], comp func(l, r []int) int) int {

	var total int

	for line := range lines {
		var left, right, _ = strings.Cut(line, ",")
		var lArr = strings.Split(left, "-")
		var rArr = strings.Split(right, "-")

		var l []int
		var r []int

		for _, str := range lArr {
			num, _ := strconv.Atoi(str)
			l = append(l, num)
		}
		for _, str := range rArr {
			num, _ := strconv.Atoi(str)
			r = append(r, num)
		}
		total += comp(l, r)
	}
	return total
}

// return 1 if one interval overlaps the other, else return 0
func overlaps(l, r []int) int {

	if l[0] >= r[0] && l[1] <= r[1] || r[0] >= l[0] && r[1] <= l[1] {
		return 1
	}
	return 0

}

func overlapsPartial(l, r []int) int {

	// left starts
	if l[0] < r[0] && l[1] < r[0] ||
		// right starts
		r[0] < l[0] && r[1] < l[0] {
		return 0
	}

	return 1
}

func main() {

	lines := strings.SplitSeq(input, "\n")

	// fmt.Println(countPairs(lines, overlaps))
	fmt.Println(countPairs(lines, overlapsPartial))

}
