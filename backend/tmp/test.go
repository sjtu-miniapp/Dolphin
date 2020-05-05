package main

import "fmt"

func main() {
	e := fmt.Errorf("hello")
	fmt.Println(e.Error() == "hello")
}
