package main

import (
	"errors"
)

var message = ""

func MustStoreMessage(s string)  {

	err := StoreMessage(s)

	if err != nil {
		panic(err)
	}

}

func StoreMessage(m string) error {
	if m == "" {
		return errors.New("no message")
	}

	message = m

	return nil
}

func main() {
	MustStoreMessage("Hello!")
}
