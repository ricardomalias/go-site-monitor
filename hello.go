package main

import (
	"fmt"
	"os"
	"reflect"
)

func main() {
	displayIntroduction()
	displayMenu()
	option := readOption()

	switch option {
	case 1:
		fmt.Println("Monitoring...")
	case 2:
		fmt.Println("Displaying logs")
	case 0:
		fmt.Println("Exiting...")
		os.Exit(0)
	default:
		fmt.Println("Unknown option")
		os.Exit(-1)
	}
}

func displayIntroduction() {
	name := "Abigobaldo"
	version := 1.2

	fmt.Println("Run ", name, " Run!")
	fmt.Println("This program is on version: ", version)
	fmt.Println(reflect.TypeOf(version))
}

func displayMenu() {
	fmt.Println("1- Start monitoring")
	fmt.Println("2- Show logs")
	fmt.Println("0- Exit")
}

func readOption() int {
	var option int
	fmt.Scan(&option)

	return option
}
