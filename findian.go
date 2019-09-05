package main

import (
	"fmt"
	"strings"
)

func main() {
	var input_str string
	i := "i"
	a := "a"
	n := "n"
	fmt.Printf("Enter a String:")
	_, err := fmt.Scanf("%s", &input_str)
	fmt.Printf("You entered %s\n", input_str)
	if err == nil {
		if strings.HasPrefix(input_str, i) && strings.HasSuffix(input_str, n) && strings.Contains(input_str, a) {
			fmt.Printf("Found\n")
		} else {
			fmt.Printf("Not Found\n")
		}
	} else {
		fmt.Print(err)
	}
}
