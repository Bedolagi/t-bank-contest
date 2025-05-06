package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	reader := bufio.NewReader(os.Stdin)
	writer := bufio.NewWriter(os.Stdout)
	defer writer.Flush()

	var n int
	fmt.Fscan(reader, &n)

	a := make([]int, n)
	b := make([]int, n)
	for i := 0; i < n; i++ {
		fmt.Fscan(reader, &a[i], &b[i])
	}

	var q int
	fmt.Fscan(reader, &q)

	for i := 0; i < q; i++ {
		var t, d int
		fmt.Fscan(reader, &t, &d)
		t--

		if d < a[t] {
			fmt.Fprintln(writer, a[t])
			continue
		}

		diff := d - a[t]
		mult := diff / b[t]
		next := a[t] + (mult+1)*b[t]
		fmt.Fprintln(writer, next)
	}
}
