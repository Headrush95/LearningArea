package main28

import (
	"bufio"
	"errors"
	"fmt"
	"log"
	"os"
)

var limitErrorMessage = "limit exceeded"

type limitError struct {
	message    string
	limit      int
	lastString string
}

func (err *limitError) Error() string {
	return fmt.Sprintf("%s, limit: %d, last string: %s\n", err.message, err.limit, err.lastString)
}

func ReadFile(src string) error {
	file, err := os.OpenFile(src, os.O_RDONLY, 0666)
	if err != nil {
		log.Fatal(err)
	}
	defer func() {
		err := file.Close()
		if err != nil {
			log.Fatal(err)
		}
	}() // close file

	scn := bufio.NewScanner(file)
	limit := 10
	countLines := 0
	line := ""
	for scn.Scan() {
		countLines++
		if countLines > limit {
			err := limitError{limitErrorMessage, limit, line}
			return &err
		}
		line = scn.Text()

	}

	if err := scn.Err(); err != nil {
		log.Fatal(err)
	}

	return nil
}

func main() {
	err := ReadFile("./2.8/in.txt")

	if _, ok := err.(*limitError); ok {
		fmt.Println("string count exceed limit, please read another file, err: ", err.Error())
		newError := fmt.Errorf("New error: %w", err)
		fmt.Println(newError)
		fmt.Println(errors.Unwrap(err))
	}

	//fmt.Printf("Total strings: %d\n", countLines)
}
