package main

import (
	"fmt"
	_ "github.com/go-chi/chi/v5"
)

type Message struct {
	userName    string
	messageText string
	timeStamp   string
}

func main() {
	fmt.Print("hello")
}
