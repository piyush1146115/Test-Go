package main

import (
	"errors"
	"fmt"
	"time"
)

type DateError struct {
	Message string
	Date    time.Time
}

func (de DateError) Error() string {
	return fmt.Sprintf("%s: %s", de.Date.String(), de.Message)
}

func NewError(message string) DateError {
	return DateError{
		Message: message,
		Date:    time.Now(),
	}
}

// PrintError can take in an DateError since DateError fulfills the error interface
func PrintError(err error) {
	fmt.Println(err.Error())
}
func main() {
	de := NewError("Auch, I failed")
	regularErr := errors.New("Another error")
	PrintError(de)
	PrintError(regularErr)
}
