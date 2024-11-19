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
	//make json file if file doesnt exist
	_, err := os.Stat("items.json")
	if os.IsNotExist(err) {
		file, err := os.Create("items.json")
		if err != nil {
			fmt.Println("Error creating file: ", err)
			os.Exit(1)
		}
		defer file.Close()
	}

	//file open
	file, err := os.OpenFile("items.json", os.O_RDWR, 0644)
	if err != nil {
		fmt.Println("Error opening file: ", err)
		os.Exit(1)
	}
	defer file.Close()

	





	if len(os.Args) < 2 {
		fmt.Println("Please provide a command")
		os.Exit(1)
	}
	cmd := os.Args[1]
	item := os.Args[2]

	switch cmd {
	case "add":
		missingItemCheck(item)
		add(item)
	case "update":
		missingItemCheck(item)
		itemID := checkValidItemId(item)
		update(itemID)
	case "remove":
		missingItemCheck(item)
		itemID := checkValidItemId(item)
		remove(itemID)
	}
}

func checkValidItemId(item string) int {
	itemID, err := strconv.Atoi(item)
	if err != nil {
		fmt.Println("Please provide a valid item id")
		os.Exit(1)
	}
	return itemID
}

func missingItemCheck(item string) {
	if item == "" {
		fmt.Println("Please provide item")
		os.Exit(1)
	}
}	