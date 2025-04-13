package main

import (
	"Melodex/TUI"
	"fmt"
	"os"
)

func main() {
	TUI.Application()
	greeter()
}

func greeter() {
	file, err := os.ReadFile("Assets/Greeter.txt")
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println("\n" + string(file) + "\n")
}
