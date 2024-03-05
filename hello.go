package main

func main() {
	ch := make(chan int)

	go func() {
		ch <- 7
	}()

	i, ok := <-ch

	print(i, ok)

	go func() {
		ch <- 11
	}()

	close(ch)

	j, ok := <-ch

	print(j, ok)

}
