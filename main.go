package main

import (
	"encoding/json"
	"errors"
	"fmt"
	"io"
	"os"
	"strconv"
	"time"
)

type Task struct {
    ID        int       `json:"id"`
    Name      string    `json:"name"`
    Status    string    `json:"status"`
    CreatedAt string `json:"createdAt"`
    UpdatedAt string `json:"updatedAt"`
}

func main() {
    // Check if a command is provided
    if len(os.Args) < 2 {
        fmt.Println("Please provide a command (add, update, remove)")
        os.Exit(1)
    }
    // Open the file
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
    // Check the command
    cmd := os.Args[1]
    switch cmd {
    case "add":
		if len(os.Args) != 3 {
            fmt.Println("Please provide an item in the format: <item_name> for 'add' command")
			os.Exit(1)
		}
		item := os.Args[2]
        tasks = add(item, tasks)
    case "list":
        list(tasks)
    case "update":
		if len(os.Args) != 3 {
			fmt.Println("Please provide an item id and name in the format: <item_id> <item_name> for 'update' command")
			os.Exit(1)
		}
		var name string
		item := os.Args[2]
		name = os.Args[3]
        itemId, err := strconv.Atoi(item)
        if err != nil {
            fmt.Println("Invalid item id must be an integer")
            os.Exit(1)
        }
        tasks = update(itemId, tasks, name)
    case "mark-in-progress":
        if len(os.Args) != 3 {
            fmt.Println("Please provide an item id in the format: <item_id> for 'mark-in-progress' command")
            os.Exit(1)
        }
        item := os.Args[2]
        itemId, err := strconv.Atoi(item)
        if err != nil {
            fmt.Println("Invalid item id, must be an integer")
            os.Exit(1)
        }
        markInProgress(tasks, itemId)
    case "mark-done":
        if len(os.Args) != 3 {
            fmt.Println("Please provide an item id in the format: <item_id> for 'mark-done' command")
            os.Exit(1)
        }
        item := os.Args[2]
        itemId, err := strconv.Atoi(item)
        if err != nil {
            fmt.Println("Invalid item id, must be an integer")
            os.Exit(1)
        }
        markDone(tasks, itemId)
    case "remove":
		if len(os.Args) != 3 {
			fmt.Println("Please provide an item id in the format: <item_id> for 'remove' command")
			os.Exit(1)
		}
		item := os.Args[2]
        itemId, err := strconv.Atoi(item)
        if err != nil {
            fmt.Println("Invalid item id, must be an integer")
            os.Exit(1)
        }
        tasks = remove(itemId, tasks)
    default:
        fmt.Println("Unknown command")
        os.Exit(1)
    }
    // Save the tasks to the file
    file.Truncate(0)
    file.Seek(0, 0)
    encoder := json.NewEncoder(file)
    err = encoder.Encode(tasks)
    if err != nil {
        fmt.Println("Error encoding tasks to file: ", err)
        os.Exit(1)
    }
}

func add(item string, tasks []Task) []Task {
    newTask := Task{
        ID:   getLastId(tasks) + 1,
        Name: item,
        Status: "todo",
        CreatedAt: time.Now().Format("RCF1123"),
        UpdatedAt: time.Now().Format("RCF1123"),
    }
    tasks = append(tasks, newTask)
    return tasks
}

func list(tasks []Task) {
    if len(tasks) == 0 {
        fmt.Println("No items in the list")
        return
    }
    if len(os.Args) > 3 {
        fmt.Println("Too many arguments for 'list' command, please specify one of the following: done, todo, in-progress")
        os.Exit(1)
    }
    var filter string
    if len(os.Args) > 2 {
        filter = os.Args[2]
    }
    switch filter {
        case "done":
            for _, task := range tasks {
                if task.Status == "done" {
                    fmt.Printf("%d: %s - %s - %s - %s\n", task.ID, task.Name, task.Status, task.CreatedAt, task.UpdatedAt)
                }
            }
        case "todo":
            for _, task := range tasks {
            if task.Status == "todo" {
                    fmt.Printf("%d: %s - %s - %s - %s\n", task.ID, task.Name, task.Status, task.CreatedAt, task.UpdatedAt)
            }
        }
        case "in-progress":
            for _, task := range tasks {
                if task.Status == "in-progress" {
                    fmt.Printf("%d: %s - %s - %s - %s\n", task.ID, task.Name, task.Status, task.CreatedAt, task.UpdatedAt)
                }
            }
        
        default:
            for _, task := range tasks {
                fmt.Printf("%d: %s - %s - %s - %s\n", task.ID, task.Name, task.Status, task.CreatedAt, task.UpdatedAt)
            }
        }
}


func markDone(tasks []Task, itemId int) []Task {
    for i, task := range tasks {
		if task.ID == itemId {
			tasks[i].Status = "done"
			return tasks
		}
	}
	fmt.Println("Item with id not found: ", itemId)
	return tasks
}

func markInProgress(tasks []Task, itemId int) []Task {
	for i, task := range tasks {
		if task.ID == itemId {
			tasks[i].Status = "in-progress"
			return tasks
		}
	}
	fmt.Println("Item with id not found: ", itemId)
	return tasks
}

func getLastId(tasks []Task) int {
    if len(tasks) == 0 {
        return 0
    }

    return tasks[len(tasks)-1].ID
}


func update(id int, tasks []Task, name string) []Task {
    for i, task := range tasks {
        if task.ID == id {
            tasks[i].Name = name
            tasks[i].UpdatedAt = time.Now().Format("RFC1123")
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
};
