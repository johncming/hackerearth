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

	cache := make(map[int]int)
	for _, a := range arr {
		if v, ok := cache[a]; ok {
			cache[a] = v + 1
		} else {
			cache[a] = 1
		}
	}

	var K int
	fmt.Scanf("%d\n", &K)

	small := 100001
	for k, v := range cache {
		if v == K {
			if k < small {
				small = k
			}
		}
	}

	fmt.Println(small)

}
