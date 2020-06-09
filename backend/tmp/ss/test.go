package main

import (
	"fmt"
	"time"
)

func main() {
	//now := time.Now()
	s, _ := time.Parse("2006-01-02T15:04:05", "2016-01-02T12:12:12")
	fmt.Print(s)
}
