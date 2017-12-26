package main

import (
	"fmt"
)

func main() {
	var N int
	fmt.Scanf("%d\n", &N)

	var arr []int
	for i := 0; i < N; i++ {
		var tmp int
		fmt.Scanf("%d", &tmp)
		arr = append(arr, tmp)
	}

	var K int
	fmt.Scanf("%d\n", &K)

	var index int
	for i, it := range arr {
		if it == K {
			index = i
			break
		}
	}

	fmt.Println(index)
}
