package main

import (
	"fmt"
)

func main() {
	var N, M int
	fmt.Scanf("%d %d\n", &N, &M)

	var result []int
	for i := 0; i < N; i++ {
		var v int
		fmt.Scanf("%d", &v)
		result = append(result, v)
	}

	var ret int
	for i, it := range result {
		if it == M {
			ret = i
		}
	}
	fmt.Println(ret + 1)
}
