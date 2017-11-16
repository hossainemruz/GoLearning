package GoLearning

import (
	"fmt"
	"time"
)
func pinger(c chan int)  {
	var n int =0
	for i:=0; ; i++{
		fmt.Println("wrinting on channel ",n," ",c);
		c <- n
		n++;
		if n>100{
			break
		}
	}
}

func ponger(c chan string)  {
	for i:=0; ; i++{
		c <- "pong"
		time.Sleep(time.Microsecond*250)
	}
}

func printer(c chan int)  {
	for{
	message:= <- c
	fmt.Println("From printer 1",message)
	time.Sleep(time.Millisecond*200)
	}
}

func pinger2(c chan string)  {
	for i:=0; ; i++{
		c <- "ping"
	}
}

func ponger2(c chan string)  {
	for i:=0; ; i++{
		c <- "pong"
		time.Sleep(time.Microsecond*250)
	}
}

func printer2(c chan int)  {
	for{
		message:= <- c
		fmt.Println("From printer 2 ", message)
		time.Sleep(time.Microsecond*800)
	}
}
func main()  {
	var c chan  int = make(chan int)
	go pinger(c)
	//go ponger(c)
	go printer(c)
	//var c2 chan  string = make(chan string)
	//go pinger(c2)
	//go ponger(c2)
	go printer2(c)
	var input string
	fmt.Scanln(&input)
	fmt.Println(input)

}