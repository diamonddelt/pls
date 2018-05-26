package main

import (
	"fmt"
	"log"
	"os/exec"
	"strings"
)

func main() {
	cmd := exec.Command("cmd", "/C", "echo", "%PATH%")
	raw, err := cmd.Output()
	if err != nil {
		log.Fatal(err)
	}

	paths := strings.Split(string(raw), ";")
	for _, v := range paths {
		fmt.Println(v)
	}
}
