package main

import (
	"encoding/json"
	"fmt"
	"github.com/go-chi/chi/v5"
	"io/ioutil"
	"net/http"
	"strconv"
)

type Message struct {
	UserName    string
	MessageText string
	TimeStamp   string
}

func main() {

	router := chi.NewRouter()
	fmt.Println("hello")
	messages := make([]Message, 0)

	router.Post("/api/Messanger", func(writer http.ResponseWriter, request *http.Request) {
		message := Message{}
		raw, _ := ioutil.ReadAll(request.Body)
		json.Unmarshal(raw, &message)
		fmt.Println(message)
		messages = append(messages, message)
	})

	router.Get("/api/Messanger/{id}", func(writer http.ResponseWriter, request *http.Request) {
		id := chi.URLParam(request, "id")
		idInt, _ := strconv.Atoi(id)
		if idInt >= 0 && idInt < len(messages) {
			message := messages[idInt]
			data, _ := json.Marshal(message)
			writer.Write(data)
		}
	})
	http.ListenAndServe(":5000", router)
}
