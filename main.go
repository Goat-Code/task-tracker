package main

import (
	"encoding/json"
	"errors"
	"fmt"
	"io"
	"os"
	"strconv"
)

type Task struct {
    ID   int    `json:"id"`
    Name string `json:"name"`
}

func main() {

    if len(os.Args) < 2 {
        fmt.Println("Please provide a command (add, update, remove)")
        os.Exit(1)
    }
    cmd := os.Args[1]
    
    var tasks []Task
    file, err := os.OpenFile("items.json", os.O_RDWR|os.O_CREATE, 0755)
    if err != nil {
        fmt.Println("Error opening file: ", err)
        os.Exit(1)
    }
    defer file.Close()

    decoder := json.NewDecoder(file)
    err = decoder.Decode(&tasks)
    if err != nil && !errors.Is(err, io.EOF) {
        fmt.Println("Error decoding file: ", err)
        os.Exit(1)
    }

    switch cmd {
    case "add":
		if len(os.Args) < 3 {
			fmt.Println("Please provide an item in the format: <item_name> for 'add' command")
			os.Exit(1)
		}
		item := os.Args[2]
        tasks = add(item, tasks)
    case "update":
		if len(os.Args) < 4 {
			fmt.Println("Please provide an item id and name in the format: <item_id> <item_name> for 'update' command")
			os.Exit(1)
		}
		var name string
		item := os.Args[2]
		name = os.Args[3]
        itemID, err := strconv.Atoi(item)
        if err != nil {
            fmt.Println("Invalid item id must be an integer")
            os.Exit(1)
        }
        tasks = update(itemID, tasks, name)
    case "remove":
		if len(os.Args) < 3 {
			fmt.Println("Please provide an item id in the format: <item_id> for 'remove' command")
			os.Exit(1)
		}
		item := os.Args[2]
        itemID, err := strconv.Atoi(item)
        if err != nil {
            fmt.Println("Invalid item id, must be an integer")
            os.Exit(1)
        }
        tasks = remove(itemID, tasks)
	case "list":
		list(tasks)
    default:
        fmt.Println("Unknown command")
        os.Exit(1)
    }
    file.Truncate(0)
    file.Seek(0, 0)
    encoder := json.NewEncoder(file)
    err = encoder.Encode(tasks)
    if err != nil {
        fmt.Println("Error encoding tasks to file: ", err)
        os.Exit(1)
    }
}

func getLastID(tasks []Task) int {
    if len(tasks) == 0 {
        return 0
    }

    return tasks[len(tasks)-1].ID
}

func add(item string, tasks []Task) []Task {
    newTask := Task{
        ID:   getLastID(tasks) + 1,
        Name: item,
    }
    tasks = append(tasks, newTask)
    return tasks
}

func update(id int, tasks []Task, name string) []Task {
    for i, task := range tasks {
        if task.ID == id {
            tasks[i].Name = name
            return tasks
        }
    }
    fmt.Println("Item with id not found: ", id)
    return tasks
}

func remove(id int, tasks []Task) []Task {
    fmt.Println("Removing item with id: ", id)
    for i, task := range tasks {
        if task.ID == id {
            return append(tasks[:i], tasks[i+1:]...)
        }
    }
    fmt.Println("Item with id not found: ", id)
    return tasks
}

func list(tasks []Task) {
	for _, task := range tasks {
		fmt.Printf("%d: %s\n", task.ID, task.Name)
	}
}