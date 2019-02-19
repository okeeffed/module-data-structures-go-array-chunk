package main

import (
	"fmt"
)

func arrayChunk(arr []int, size int) [][]int {
	if len(arr) <= 0 {
		return nil
	}

	if size == 0 {
		return nil
	}

	var tmp []int
	var chunkedArr [][]int
	for i, item := range arr {
		tmp = append(tmp, item)

		isFinal := len(arr)-1 == i
		if i%size == size-1 || isFinal {
			fmt.Printf("%v", tmp)
			chunkedArr = append(chunkedArr, tmp)
			tmp = nil
		}
	}

	return chunkedArr
}

func main() {
	fmt.Println("Running array chunk")
}
