package main

import ( "fmt")

type CheckNaga struct {
	naga map[int]Event
	raj string
}
type Event struct {
	value int
}

func main() {

	checknaga := CheckNaga{naga: map[int]Event{}}
	event := checknaga.naga[1]
	fmt.Println(event)

	event.value = 100

	fmt.Println(event)
}