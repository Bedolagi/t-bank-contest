package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func main() {
	scanner := bufio.NewScanner(os.Stdin)
	scanner.Scan()
	n, _ := strconv.Atoi(scanner.Text())
	scanner.Scan()
	numsStr := strings.Split(scanner.Text(), " ")
	nums := make([]int, n)
	for i := 0; i < n; i++ {
		nums[i], _ = strconv.Atoi(numsStr[i])
	}

	used := make(map[int]bool)
	count := 0

	for _, num := range nums {
		current := num
		for {
			if !used[current] {
				used[current] = true
				count++
				break
			}
			if current == 0 {
				break
			}
			current = current / 2
		}
	}

	fmt.Println(count)
}
