package main

import (
	"fmt"
	"io"
	"os"
)

func permutations(arr []int) [][]int {
	var helper func([]int, int)
	res := [][]int{}

	helper = func(arr []int, n int) {
		if n == 1 {
			tmp := make([]int, len(arr))
			copy(tmp, arr)
			if arr[0]*10+arr[1] < 24 && arr[2]*10+arr[3] < 60 {
				res = append(res, tmp)
			}
		} else {
			for i := 0; i < n; i++ {
				helper(arr, n-1)
				if n%2 == 1 {
					tmp := arr[i]
					arr[i] = arr[n-1]
					arr[n-1] = tmp
				} else {
					tmp := arr[0]
					arr[0] = arr[n-1]
					arr[n-1] = tmp
				}
			}
		}
	}
	helper(arr, len(arr))
	return res
}

func main() {
	arr := []int{1, 2, 3, 4}
	resArr := [][]int{}
	resArr = permutations(arr)
	fmt.Println("Valid times: ")
	for _, perm := range resArr {
		s := fmt.Sprintf("%d%d:%d%d\n", perm[0], perm[1], perm[2], perm[3])
		io.WriteString(os.Stdout, s)
	}
	fmt.Println(fmt.Sprintf("Return: %d", len(resArr)))
}
