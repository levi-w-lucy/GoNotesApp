package todo

import (
	"encoding/json"
	"errors"
	"fmt"
	"os"
)

type Todo struct {
	Text string `json:"text"`
}

func (todo Todo) Display() {
	fmt.Println(todo.Text)
}

func (todo Todo) Save() (err error) {
	fileName := "todo.json"

	json, err := json.Marshal(todo)
	if err != nil {
		return err
	}

	err = os.WriteFile(fileName, json, 0644)
	return
}

func New(content string) (Todo, error) {
	if content == "" {
		return Todo{}, errors.New("title or Content was empty. Please enter a valid value")
	}

	return Todo{
		Text: content,
	}, nil
}
