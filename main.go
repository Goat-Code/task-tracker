package main

import (
	"encoding/json"
	"fmt"
	"os"
	"strconv"
	"time"
)

type Task struct {
	ID          int       `json:"id"`
	Description string    `json:"description"`
	Status      string    `json:"status"`
	CreatedAt   time.Time `json:"created_at"`
	UpdatedAt   time.Time `json:"updated_at"`
}

func openFile() *os.File {
	file, err := os.Open("items.json")
	if err != nil {
		fmt.Println("Error opening file: ", err)
		os.Exit(1)
	}
	return file
}

func getLastId() int {
	var tasks []Task
	file := openFile()
	
	decoder := json.NewDecoder(file)
	err = decoder.Decode(&tasks)
	if err != nil {
		fmt.Println("Error decoding file: ", err)
		os.Exit(1)
	}
	return tasks[len(tasks)-1].ID
}

func add(item string) {
	var task Task
	task.ID = getLastId() + 1
	task.Description = item
	task.Status = "todo"
	task.CreatedAt = time.Now()
	task.UpdatedAt = time.Now()

	file := openFile()
	decoder := json.NewDecoder(file)
	var tasks []Task
	err := decoder.Decode(&tasks)
	if err != nil {
		fmt.Println("Error decoding file: ", err)
		os.Exit(1)
	}
	tasks = append(tasks, task)

	
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