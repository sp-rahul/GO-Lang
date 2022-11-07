package main

import (
	"errors"
	"fmt"
	"strconv"

	"google.golang.org/protobuf/internal/encoding/text"
)

func main() {
	result, err := returnError(returnError: true)
	if err != nil {
		fmt.Println(err)
	} else {
		fmt.Println(result)
	}

}
type specialError struct {
	errorMessage string
	errorCode int
}

func (s specialError) Error() string {
	return s.errorMessage + " " + strconv.Itoa(s.errorCode)
}

func returnError(returnError bool) (string, specialError) {

	return errors.New{errorMessage: "Special error", errorCode: 123}
}
type error interface {
	Error() string
}