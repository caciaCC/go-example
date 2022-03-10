package main

func main() {
	intChan := make(chan int, 2)
	intChan <- 1
	<-intChan
	<-intChan
}
