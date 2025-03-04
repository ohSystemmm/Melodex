package main

import (
	"Melodex/Backend"
	"Melodex/TUI"
	"fmt"
	"os"
)

func main() {
	greeter()
	Backend.Run()
	TUI.Application()

}

func greeter() {
	file, err := os.ReadFile("Assets/Greeter.txt")
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println("\n" + string(file) + "\n")
}
