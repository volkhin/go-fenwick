package main

import (
	"fmt"

	"github.com/volkhin/go-fenwick"
)

func main() {
	f := fenwick.New(10)
	f.Add(1, 4)
	f.Add(3, 123)
	fmt.Println(f.Sum(1, 10)) // 127
}
