package main

import (
	"fmt"
	"os"
	"strings"
)

func main() {
	// go build echo3.go
	// ./echo3 echo Hello!
	fmt.Println(strings.Join(os.Args[1:], " "))
}
