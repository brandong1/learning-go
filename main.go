package main

import (
	"fmt"
	f "fmt"
	"runtime"
)

func main() {
	fmt.Println("Hi! I want to be a Gopher!")
	f.Println("Hi! I want to be a Pickle!")
	f.Println(runtime.NumCPU())
}
