package main

import (
"fmt"
"strings"
)


type Animal struct{
    name string
    food string
    locomotion string
    noise string
}

func (animal *Animal) Eat(){
    fmt.Printf("The %v eats %v\n",animal.name,animal.food)
}

func (animal *Animal) Move(){
    fmt.Printf("The %v can %v\n",animal.name,animal.locomotion)
}

func (animal *Animal) Speak(){
    fmt.Printf("The %v goes %v\n",animal.name,animal.locomotion)
}

func (animal *Animal) call(action string) {
     switch action {
       case "eat" : animal.Eat()
       case "move" : animal.Move()
       case "speak" : animal.Speak()
       default : fmt.Println("Invalid Action")
     }
}

func main(){
    cow := Animal{"cow","grass","walk","moo"}
    bird := Animal{"bird","worms","fly","peep"}
    snake := Animal{"snake","mice","slither","hsss"}

//    cow.Eat()
//    cow.Move()
//    cow.Speak()
    take := true
    for take {
    var action string
    fmt.Printf("Enter the animal and action:")
    _,err := fmt.Scanf("%s",&action)
    if err == nil {
    fmt.Println(action)
    input_arr := strings.Split(action,",")
    animal_name,action_name := input_arr[0],input_arr[1]
    //fmt.Printf("%v and %v\n",animal_name,action_name)
    if animal_name == "cow"{
       cow.call(action_name)
    }else if animal_name == "snake"{
       snake.call(action_name)
    }else if animal_name == "bird"{
       bird.call(action_name)
    }else{
       fmt.Println("Invalid Animal Name")
    }
    }else{
    take = false
    }
    }
}
