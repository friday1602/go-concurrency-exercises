package main

import "fmt"

func main() {
	ch := make(chan int)
	go func() {
		for i := 0; i < 6; i++ {
			// TODO: send iterator over channel
			ch <- i
		}
		close(ch)
	}()

	// TODO: range over channel to recv values
	for v := range ch {
		fmt.Println(v)
	}

}
