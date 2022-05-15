package main

import (
	"fmt"
	"os"
	"time"
)

func text() {
	fmt.Print("\033[H\033[2J")
	fmt.Print("Loading.")
	time.Sleep(1 * time.Second)
	fmt.Print("\033[H\033[2J")
	fmt.Print("Loading..")
	time.Sleep(1 * time.Second)
	fmt.Print("\033[H\033[2J")
	fmt.Print("Loading...")
	time.Sleep(1 * time.Second)
	fmt.Print("\033[H\033[2J")
	fmt.Print("Loading.")
	time.Sleep(1 * time.Second)
	fmt.Print("\033[H\033[2J")
	fmt.Print("Loading..")
	time.Sleep(1 * time.Second)
	fmt.Print("\033[H\033[2J")
	fmt.Print("Loading...")
	time.Sleep(1 * time.Second)
	fmt.Print("\033[H\033[2J")
	fmt.Print("Loading.")
	time.Sleep(1 * time.Second)
	fmt.Print("\033[H\033[2J")
	fmt.Print("Loading..")
	time.Sleep(1 * time.Second)
	fmt.Print("\033[H\033[2J")
	fmt.Print("Loading...")
	time.Sleep(3 * time.Second)
	os.Exit(0)
}

func ascii() {
	fmt.Print("\033[H\033[2J")
	fmt.Println(" _                 _ _             ")
	time.Sleep(1 * time.Second)
	fmt.Println("| |               | (_)            ")
	time.Sleep(1 * time.Second)
	fmt.Println("| | ___   __ _  __| |_ _ __   __ _ ")
	time.Sleep(1 * time.Second)
	fmt.Println("| |/ _ \\ / _` |/ _` | | '_ \\ / _` |")
	time.Sleep(1 * time.Second)
	fmt.Println("| | (_) | (_| | (_| | | | | | (_| |")
	time.Sleep(1 * time.Second)
	fmt.Println("|_|\\___/ \\__,_|\\__,_|_|_| |_|\\__, |")
	time.Sleep(1 * time.Second)
	fmt.Println("                              __/ |")
	time.Sleep(1 * time.Second)
	fmt.Println("                             |___/")
	time.Sleep(3 * time.Second)
}

func main() {
	fmt.Println("[1] ASCII")
	fmt.Println("[2] Text\n")

	var option int

	fmt.Println("Pick an option 1-2: ")
	fmt.Scanln(&option)

	if option != 1 {
		text()
	}
	if option != 2 {
		ascii()
	}
}