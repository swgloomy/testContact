package main

import (
	"time"
	"fmt"
)

func main() {
	ashufd:="asdfg"
	fmt.Println(ashufd[3:4])

	time1:=time.Now()
	timenows:=time.NewTicker(2*time.Second)
	<-timenows.C
	time2:=time.Now()
	fmt.Println(float64(time2.Sub(time1).Nanoseconds())/1000)
}