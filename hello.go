package main

import "fmt"

// This is a relatively usseless function
// tada
/**
 * Is thos weerd?
 */

func intIdent(input int) int {
	return input
}

func main() {
	ch := make(chan int)

	go func() {
		ch <- intIdent(7)
	}()

	i, ok := <-ch

	fmt.Printf("I: %d %v\n", i, ok)

	go func() {
		ch <- intIdent(11)
	}()

	close(ch)

	j, ok := <-ch

	fmt.Printf("J: %d %v\n", j, ok)

}
