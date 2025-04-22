package main

import (
	"fmt"
	"os"

	"Melodex/src/backend/music"
	tui "Melodex/src/tui"
)

func main() {
	music.Init()
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
