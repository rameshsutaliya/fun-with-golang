package main

import (
	"bufio"
	"fmt"
	"fun-with-golang/pocs/lru/service"
	"os"
	"strings"
)

func main() {
	reader := bufio.NewReader(os.Stdin)
	s := service.NewService(false)
	fmt.Println("These are the available commands use user perform")
	fmt.Println("add\tremove\tget\tput\tgetAll\tevictAll")
	for {
		fmt.Println("Please enter the command")
		input, err := reader.ReadString('\n')
		if err != nil {
			fmt.Println("Error reading input:", err)
		}
		input = strings.TrimSpace(input)
		switch input {
		case "add":
			key, err := reader.ReadString('\n')
			if err != nil {
				fmt.Println("Error reading input:", err)
			}
			key = strings.TrimSpace(key)
			val, err := reader.ReadString('\n')
			if err != nil {
				fmt.Println("Error reading input:", err)
			}
			val = strings.TrimSpace(val)
			if s.Post(key, val) {
				fmt.Println("Added successfully")
			} else {
				fmt.Println("Unable to Add")
			}
		case "remove":
			key, err := reader.ReadString('\n')
			if err != nil {
				fmt.Println("Error reading input:", err)
			}
			key = strings.TrimSpace(key)
			if s.RemoveFromCache(key) {
				fmt.Println("Removed Successfully")
			} else {
				fmt.Println("There is no such key in cache")
			}
		case "get":
			key, err := reader.ReadString('\n')
			if err != nil {
				fmt.Println("Error reading input:", err)
			}
			key = strings.TrimSpace(key)
			val, ok := s.Get(key)
			if ok {
				fmt.Println("Received value from the cache is: ", val)
			} else {
				fmt.Println("There is no such key in the cache")
			}
		case "put":
			key, err := reader.ReadString('\n')
			if err != nil {
				fmt.Println("Error reading input:", err)
			}
			key = strings.TrimSpace(key)
			val, err := reader.ReadString('\n')
			if err != nil {
				fmt.Println("Error reading input:", err)
			}
			val = strings.TrimSpace(val)
			if s.Put(key, val) {
				fmt.Println("Updated successfully")
			} else {
				fmt.Println("Unable to update")
			}
		case "getAll":
			fmt.Println(s.GetAllCache())
		case "evictAll":
			s.EraseCache()
		}

		if input == "\x1b" {
			fmt.Println("Exiting...")
			return
		}
	}
}
