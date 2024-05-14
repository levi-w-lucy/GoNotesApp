package note

import (
	"encoding/json"
	"errors"
	"fmt"
	"os"
	"strings"
	"time"
)

type Note struct {
	title     string
	content   string
	createdAt time.Time
}

func (note Note) Display() {
	fmt.Printf("Your note titled %v has the following content:\n\n%v\n\n", note.title, note.content)
}

func (note Note) Save() (err error) {
	fileName := strings.ReplaceAll(note.title, " ", "_")
	fileName = strings.ToLower(fileName)

	json, err := json.Marshal(note)
	if err != nil {
		return err
	}

	err = os.WriteFile(fileName, json, 0644)
	return
}

func New(title string, content string) (Note, error) {
	if title == "" || content == "" {
		return Note{}, errors.New("title or Content was empty. Please enter a valid value")
	}

	return Note{
		title:     title,
		content:   content,
		createdAt: time.Now(),
	}, nil
}

func (note *Note) updateTitle(title string) {
	note.title = title
}

func (note *Note) updateContent(content string) {
	note.content = content
}
