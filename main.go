package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"

	"example.com/GoNotesApp/note"
	"example.com/GoNotesApp/todo"
)

type saver interface {
	Save() error
}

type outputtable interface {
	saver
	Display()
}

func main() {
	newNote, err := getNewNote()
	if err != nil {
		fmt.Println("Received error: ", err)
		return
	}

	todoTask, err := getNewTodo()
	if err != nil {
		fmt.Println("Received error: ", err)
		return
	}

	err = outputData(todoTask)
	if err != nil {
		return
	}

	err = outputData(newNote)
}

func outputData(data outputtable) error {
	data.Display()
	err := saveData(data)
	return err
}

func saveData(data saver) error {
	err := data.Save()

	if err != nil {
		fmt.Printf("Saving data failed with error %v\n", err)
		return err
	}
	fmt.Println("Data saved successfully!")
	return nil
}

func getNewTodo() (todo.Todo, error) {
	text := getUserInput("Enter a todo task: ")
	return todo.New(text)
}

func getNewNote() (newNote note.Note, err error) {
	title := getUserInput("Enter note title: ")
	content := getUserInput("Enter note content: ")

	newNote, err = note.New(title, content)
	return
}

func getUserInput(prompt string) (value string) {
	fmt.Print(prompt)

	reader := bufio.NewReader(os.Stdin)
	value, err := reader.ReadString('\n')

	if err != nil {
		return ""
	}

	value = strings.TrimSuffix(value, "\n")
	value = strings.TrimSuffix(value, "\r")
	return
}
