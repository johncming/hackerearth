package main

import (
	"fmt"
)

func main() {
	var T int
	fmt.Scanf("%d\n", &T)

	vowels := map[rune]struct{}{
		'A': struct{}{},
		'E': struct{}{},
		'I': struct{}{},
		'O': struct{}{},
		'U': struct{}{},

		'a': struct{}{},
		'e': struct{}{},
		'i': struct{}{},
		'o': struct{}{},
		'u': struct{}{},
	}

	for i := 0; i < T; i++ {
		var s string
		cnt := 0
		fmt.Scanf("%s\n", &s)
		for _, c := range s {
			if _, ok := vowels[c]; ok {
				cnt++
			}
		}
		fmt.Println(cnt)
	}
}
