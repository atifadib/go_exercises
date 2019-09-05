package main

import (
	"fmt"
	"sort"
)

func main() {
	var num int
	sli := make([]int, 0, 3)
	for true {
		_, err := fmt.Scanf("%d", &num)
		if err == nil {
			sli = append(sli, num)
			sort.Ints(sli)
			fmt.Print(sli)
			fmt.Printf("\n")
		} else {
			break
		}
	}
	fmt.Print(sli)
	fmt.Printf("\n")
}
