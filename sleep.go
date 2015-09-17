package main

import (
  "fmt"
  "time"
)

func sleep() {
	fmt.Println("Sleep begins")
	<-time.After(time.Second*5)
	fmt.Println("Sleep Ends")
}

func main() {
    sleep()
	fmt.Println("Sleep time is over")
}