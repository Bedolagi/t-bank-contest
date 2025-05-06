package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func min(a, b int64) int64 {
	if a < b {
		return a
	}
	return b
}

func main() {
	reader := bufio.NewReader(os.Stdin)
	var n int
	var a, b int64
	fmt.Fscanf(reader, "%d %d %d\n", &n, &a, &b)

	s, _ := reader.ReadString('\n')
	s = strings.TrimSpace(s)

	open := 0
	close := 0
	balance := 0

	for _, c := range s {
		if c == '(' {
			balance++
		} else {
			if balance > 0 {
				balance--
			} else {
				close++
			}
		}
	}
	open = balance

	total := open + close
	pairs := min(int64(open), int64(close))
	cost := min(
		int64(total)*b,
		pairs*min(a, 2*b)+int64(total-2*int(pairs))*b,
	)

	fmt.Println(cost)
}
