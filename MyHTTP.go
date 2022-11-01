package main

import (
	"log"
	"net/http"
)

type Handler struct{}

func (h *Handler) ServeHTTP(w http.ResponseWriter, req *http.Request) {
	_, err := w.Write([]byte("Я работаю"))
	if err != nil {
		log.Println(err)
	} else {
		log.Println("Запрос обработан")
	}

}

func main() {
	log.Println(http.ListenAndServe("127.0.0.1:9090", &Handler{}))
}
