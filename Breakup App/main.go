package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func main() {
	var N int
	fmt.Scanf("%d\n", &N)

	cache := map[byte]map[string]int{
		'G': make(map[string]int),
		'M': make(map[string]int),
	}

	r := bufio.NewReader(os.Stdin)
	for i := 0; i < N; i++ {
		s, _ := r.ReadString('\n')
		s = strings.TrimRight(s, "\n")

		var c map[string]int
		switch s[0] {
		case 'G':
			c = cache['G']
		case 'M':
			c = cache['M']
		}

		items := strings.Split(s, " ")
		for _, it := range items {
			if len(it) > 0 && it[0] >= '0' && it[0] <= '9' {
				if v, ok := c[it]; ok {
					c[it] = v + 1
				} else {
					c[it] = 1
				}
			}
		}
	}

	var g, m int
	for k, v := range cache['G'] {
		if k == "19" || k == "20" {
			g = g + 4*v
		}
	}
	for k, v := range cache['M'] {
		if k != "19" && k != "20" {
			m = m + 3*v
		}
	}

	if g > m {
		fmt.Println("Date")
	} else {
		fmt.Println("No Date")
	}
}
