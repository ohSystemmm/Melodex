package main

import (
	"Melodex/Backend"
	"Melodex/TUI"
	"fmt"
	"os"
)

func main() {
	Backend.Run()
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
