package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"

	"example.com/GoNotesApp/note"
)

func main() {
	newNote, err := getNewNote()
	if err != nil {
		fmt.Println("Received error: ", err)
	}

	newNote.Display()
	err = newNote.Save()

	if err != nil {
		fmt.Printf("Saving the note failed with error %v\n", err)
		return
	}

	fmt.Println("Successfully saved note!")
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
