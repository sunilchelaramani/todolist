package main

import (
	"bufio"
	"fmt"
	"os"
)

type Todo struct {
	Title     string
	Completed bool
}

func main() {
	// Create a new to-do list.
	todoList := []Todo{}

	// Read the to-do list from the file.
	if err := readToDoListFromFile("todo.txt", &todoList); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

	// Print the to-do list.
	fmt.Println("To-do list:")
	for _, todo := range todoList {
		fmt.Println(todo.Title, todo.Completed)
	}

	// Add a new to-do item.
	scanner := bufio.NewScanner(os.Stdin)
	for {
		fmt.Printf("Enter a new to-do item or 'quit' to exit: ")
		scanner.Scan()
		todoItem := scanner.Text()
		if todoItem == "quit" {
			break
		} else if todoItem == "" {
			continue
		} else {
			todo := Todo{
				Title:     todoItem,
				Completed: false,
			}
			todoList = append(todoList, todo)
		}
	}

	// Write the to-do list to the file.
	if err := writeToDoListToFile("todo.txt", todoList); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}

func readToDoListFromFile(filename string, todoList *[]Todo) error {
	file, err := os.Open(filename)
	if err != nil {
		return err
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		line := scanner.Text()
		todo := Todo{
			Title:     line,
			Completed: false,
		}

		*todoList = append(*todoList, todo)
	}

	return nil
}

func writeToDoListToFile(filename string, todoList []Todo) error {
	file, err := os.Create(filename)
	if err != nil {
		return err
	}
	defer file.Close()

	for _, todo := range todoList {
		line := fmt.Sprintf("%s\n", todo.Title)
		file.WriteString(line)
	}

	return nil
}
