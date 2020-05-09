package main

import "fmt"

func main() {
	i := 1
	j := 2

	sum := 0
	for j < 4000000 {
		if j % 2 ==0 {
			sum += j
		}
		next := i + j
		i = j
		j = next

	}
	fmt.Println(sum)
}
