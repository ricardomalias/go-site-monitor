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
			printLog()
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
	version, _ := ioutil.ReadFile("./resources/version.txt")
	banner, _ := ioutil.ReadFile("./resources/logo.txt")

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
	_, err := fmt.Scan(&option)
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

		registerLog(url, response.StatusCode == 200)

		time.Sleep(time.Duration(DelayMinutes) * time.Second)
		fmt.Println(" ")
	}

	fmt.Println(" ")
}

func readUrlFile() []string {
	file, err := os.Open("./resources/urls.txt")
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

func registerLog(url string, status bool) {
	file, err := os.OpenFile("./resources/log.txt", os.O_RDWR|os.O_CREATE|os.O_APPEND, 0666)

	if err != nil {
		fmt.Println(err)
	}

	timeFormat := time.Now().Format("2006-01-02 15:04:05")
	statusText := "online"
	if status == false {
		statusText = "offline"
	}

	_, err = file.WriteString(timeFormat + " - site: " + url + " is " + statusText + "\n")
	if err != nil {
		return
	}

	err = file.Close()
	if err != nil {
		return
	}
}

func printLog() {
	file, err := ioutil.ReadFile("./resources/log.txt")
	if err != nil {
		return
	}

	fmt.Println("Displaying logs below:")
	fmt.Println(" ")
	fmt.Println(string(file))
}
