package main

import (
	"bufio"
	"fmt"
	"myapp/logger"
	"os"
	"time"
)

func main() {
	reader := bufio.NewReader(os.Stdin)
	ch := make(chan string)

	go logger.LisenForLog(ch)

	fmt.Println("EnterSomething")
	for i := 0; i < 5; i++ {
		fmt.Print("-> ")
		Input, _ := reader.ReadString('\n')
		ch <- Input
		time.Sleep(time.Second)
	}
	// for i := 1; i <= 10; i++ {//این یک حلقه تو در تو است
	// 	fmt.Println("i is", i, )
	// 	for j := 1; j <= 3; j++ {
	// 		fmt.Println("j:", j, )
	// 	}
	// 	fmt.Println()
	// }
}
