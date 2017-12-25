package main

import (
	"fmt"
)

func main() {
	var nodes, edges int
	fmt.Scanf("%d %d\n", &nodes, &edges)

	var matrix [][]bool
	for i := 0; i < nodes; i++ {
		var tmp []bool
		for j := 0; j < nodes; j++ {
			tmp = append(tmp, false)
		}
		matrix = append(matrix, tmp)
	}

	for i := 0; i < edges; i++ {
		var v, w int
		fmt.Scanf("%d %d\n", &v, &w)
		matrix[v][w] = true
	}

	var queries int
	fmt.Scanf("%d\n", &queries)
	for i := 0; i < queries; i++ {
		var v, w int
		fmt.Scanf("%d %d\n", &v, &w)
		if matrix[v][w] {
			fmt.Println("YES")
		} else {
			fmt.Println("NO")
		}
	}

}
