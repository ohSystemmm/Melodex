package main

import (
	"fmt"
	"os"

	tui "Melodex/tui"
)

func main() {
	tui.Application()
	greeter()
}

func greeter() {
	file, err := os.ReadFile("assets/usBTW.txt")
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println("\n" + string(file) + "\n")
}
