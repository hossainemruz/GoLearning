package GoLearning

import (
	"fmt"
	"time"
	"math/rand"
)

func secondRoutine() {

	for i := 0; i < 100; i++ {
		fmt.Println("From secondRoutine: ", i)
		randomSleepTime := time.Duration(rand.Intn(100))
		time.Sleep(time.Millisecond*randomSleepTime)
	}
}
func main() {
	go secondRoutine()
	for i := 0; i < 100; i++ {
		fmt.Println("From main: ", i)
		randomSleepTime:=time.Duration(rand.Intn(250))
		time.Sleep(time.Millisecond * randomSleepTime)
	}
}
