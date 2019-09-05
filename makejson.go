package main

import (
	"bufio"
	"encoding/json"
	"fmt"
	"os"
)

func main() {
	scanner := bufio.NewReader(os.Stdin)

	fmt.Printf("Enter Name:")
	name, err1 := scanner.ReadString('\n')
	fmt.Printf("Enter Address:")
	addr, err2 := scanner.ReadString('\n')

	if err1 == nil && err2 == nil {
		fmt.Print(name, addr)
		var idmap map[string]string
		idmap = make(map[string]string)
		idmap["name"] = name[:len(name)-1]
		idmap["address"] = addr[:len(addr)-1]
		fmt.Print(idmap)
		fmt.Print("\n")
		b_arr, err := json.Marshal(idmap)
		if err == nil{
                        fmt.Printf("Byte Array: ")
                        fmt.Print(b_arr)
                        fmt.Printf("\n JSON Object: ")
                        fmt.Printf(string(b_arr))
                        fmt.Printf("\n")
		}
	}

}
