package main

import(
 "fmt"
)

func main(){
var num float64

fmt.Printf("Enter a floating point number: ")
_,err := fmt.Scanf("%f",&num)

if err== nil{
var trunc int = int(num)
//trunc = math.Trunc(num)
fmt.Printf("Integer Number: %d\n",trunc)
}

}

