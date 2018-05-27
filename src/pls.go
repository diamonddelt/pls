package main

import (
	"fmt"
	"strings"
	"os"
)

func main() {
	for _, v := range strings.Split(os.Getenv("PATH"), ";") {
		fmt.Println(v)
	}
}
