package main

import (
	"sort"
)

func main() {
	type a struct {
		v int
	}
	b := []*a{&a{1}, &a{2}}
	sort.Slice(b,
		func(i, j int) bool {
			return b[i].v > b[j].v
		})
	print(b[0].v)
}
