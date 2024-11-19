package main

import (
	"fmt"
	"os"
	"strconv"
)

func add(item string) {
	fmt.Println("Adding new item: ", item)
}

func update(id int) {
	fmt.Println("Updating item with id: ", id)
}

func remove(id int) {
	fmt.Println("Removing item with id: ", id)
}

func main() {
	if len(os.Args) < 2 {
		fmt.Println("Please provide a command")
		os.Exit(1)
	}
	cmd := os.Args[1]
	item := os.Args[2]

	switch cmd {
	case "add":
		if item == "" {
			fmt.Println("Please provide item")
			os.Exit(1)
		}
		add(item)
	case "update":
		if item == "" {
			fmt.Println("Please provide item id")
			os.Exit(1)
		}
		itemID, err := strconv.Atoi(item)
		if err != nil {
			fmt.Println("Please provide a valid item id")
			os.Exit(1)
		}
		update(itemID)
	case "remove":
		if item == "" {
			fmt.Println("Please provide item id")
			os.Exit(1)
		}
		itemID, err := strconv.Atoi(item)
		if err != nil {
			fmt.Println("Please provide a valid item id")
			os.Exit(1)
		}
		remove(itemID)
	}
}	