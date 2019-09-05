package main

import (
"fmt"
"math"
"os"
)

func GenDisplaceFn(acc,vel,dis float64) func(float64) float64{
    new_function := func(time float64) float64{
              var disp float64 = (0.5*acc*math.Pow(time,2)) + (vel*time) + dis
              return disp
    }
    return new_function
}


func main(){
fmt.Printf("Enter Initial Acceleration:")
var acc float64
_,err1 := fmt.Scanf("%f",&acc)
if err1 != nil {
    os.Exit(1)
}
fmt.Printf("Enter Initial Velocity:")
var vel float64
_,err2 := fmt.Scanf("%f",&vel)
if err2 != nil {
   os.Exit(2)
}
fmt.Printf("Enter Initial Displacement:")
var dis float64
_,err3 := fmt.Scanf("%f",&dis)
if err3 != nil{
   os.Exit(3)
}
compute_dist_at_t := GenDisplaceFn(acc,vel,dis)

fmt.Printf("Enter Time step:")
var time float64
_,err4 := fmt.Scanf("%f",&time)
if err4==nil{
fmt.Printf("Displacement at time %f: %f\n",time,compute_dist_at_t(time))
}else{
fmt.Println("Error Occured")
}
}
