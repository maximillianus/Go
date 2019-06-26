package main

import (
	"fmt"
	"time"
)

type TheStruct struct {
	the_time time.Time
}

func main() {
	t := TheStruct{time.Now()}
	tt := time.Now().Format("2006-01-02T15:04:05")

	fmt.Printf("Time: %s\n", tt)
	fmt.Printf("%+v\n", t)
	fmt.Printf("%v\n", t)

	the_struct_2 := TheStruct{time.Unix(1501200000, 200)}
	//63636796800
	fmt.Println(the_struct_2)
	fmt.Printf("%+v\n", the_struct_2)
	fmt.Println(the_struct_2.the_time)
}
