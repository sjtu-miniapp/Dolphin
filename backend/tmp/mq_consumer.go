package main

import (
	"fmt"
	"time"
)

func main() {
	str := "2008-01-23T21:12:45"
	t, _ := time.Parse("2006-01-02T15:04:05", str)
	fmt.Println(t)
}