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
