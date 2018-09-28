package main

import ("fmt"; "sync")
// Channel example

var wg sync.WaitGroup

func naga(s string) {
    for i:=0; i<5; i++ {
fmt.Println(s)
	}
	wg.Done()    
}
func channe(c chan int, n int) {
	defer wg.Done()
c <- n
}
func main() {
	Mychanner := make(chan int, 5)
	wg.Add(1)
	go naga("hi how are you")
	wg.Add(1)
	go naga("i am naga")
	wg.Wait()

	for i:=10; i<15; i++ {
		wg.Add(1)
channe(Mychanner, i)
	}
	wg.Wait()
	close(Mychanner)

	for item := range Mychanner {
		fmt.Println(item)
	}
}

