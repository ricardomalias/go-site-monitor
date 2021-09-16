package main

import (
	"bufio"
	"fmt"
	"io"
	"io/ioutil"
	"net/http"
	"os"
	"strings"
	"time"
)

const DelayMinutes int = 5

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
	version, _ := ioutil.ReadFile("./version.txt")
	banner, _ := ioutil.ReadFile("./logo.txt")

	fmt.Println(string(banner))
	fmt.Println("Be welcome to site monitoring !")
	fmt.Println("This program is on version:", string(version))
	fmt.Println(" ")
}

func displayMenu() {
	fmt.Println("1- Start monitoring")
	fmt.Println("2- Show logs")
	fmt.Println("0- Exit")
}

func readOption() int {
	var option int
	option, err := fmt.Scan(&option)
	if err != nil {
		return 0
	}

	return option
}

func startMonitor() {
	fmt.Println("Monitoring...")

	urls := readUrlFile()

	for _, url := range urls {
		response, err := http.Get(url)

		if err != nil {
			fmt.Println("error on monitoring", err)
		}

		if response.StatusCode == 200 {
			fmt.Println("Site:", url, "is online")
		} else {
			fmt.Println("Site:", url, "is offline")
		}

		time.Sleep(time.Duration(DelayMinutes) * time.Second)
		fmt.Println(" ")
	}

	fmt.Println(" ")
}

func readUrlFile() []string {
	file, err := os.Open("./urls.txt")
	var urls []string

	if err != nil {
		fmt.Println("error:", err)
	}

	reader := bufio.NewReader(file)
	for {
		line, err := reader.ReadString('\n')
		line = strings.TrimSpace(line)

		urls = append(urls, line)

		if err == io.EOF {
			break
		}
	}

	closeErr := file.Close()
	if closeErr != nil {
		return nil
	}

	return urls
}
