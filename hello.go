package main

import (
	"fmt"
	"io/ioutil"
	"net/http"
	"os"
)

var URL = "https://forleven.com"

func main() {
	displayIntroduction()

	for {
		displayMenu()
		option := readOption()

		switch option {
		case 1:
			startMonitor()
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
}

func displayIntroduction() {
	name := "Abigobaldo"
	version := 1.2
	banner, _ := ioutil.ReadFile("./logo.txt")

	fmt.Println(string(banner))
	fmt.Println("Run ", name, " Run!")
	fmt.Println("This program is on version: ", version)
}

func displayMenu() {
	fmt.Println("1- Start monitoring")
	fmt.Println("2- Show logs")
	fmt.Println("0- Exit")
}

func readOption() int {
	var option int
	_, err := fmt.Scan(&option)
	if err != nil {
		return 0
	}

	return option
}

func startMonitor() *http.Response {
	fmt.Println("Monitoring...")
	response, err := http.Get(URL)
	if err != nil {
		return nil
	}

	if response.StatusCode == 200 {
		fmt.Println("Site:", URL, "is online")
	} else {
		fmt.Println("Site:", URL, "is offline")
	}

	return response
}
