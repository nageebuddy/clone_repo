package main

import "fmt"

type Naga struct {
	first,last string
	age Age
}
type Age struct {
	value map[int]int
}

func main() {
	naga := Naga{}
	naga.first = "nagarajan"
	naga.last = "manokaran"
	naga.age = Age{value: map[int]int{}}
	naga.age.value[1] = 2345

fmt.Println(naga.age.value[1])
}