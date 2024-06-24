package main

import (
	"fmt"
	"net/http"
)

func main() {
	HandleRequest()
}

// ResponseWriter - тип данных (структура), для обращения к определенной странице и для показания чего либа пользователю на данной странице(текст или html страницу и тд)
// Requext - для отслеживания подключения на странице
func Homepage(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Go is super easy!") // Выведем наш стринг на страницу w
}

func Contactspage(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "My contacts!")
}

// Моя функция
func HandleRequest() {
	//функция HandleFunc принимает 2 параметра. Первый это url адрес который мы хотим отслеживать. Второй это метод
	http.HandleFunc("/", Homepage)
	http.HandleFunc("/contacts/", Contactspage)
	// ListenAndServe принимает 2 параметра. Первый это порт, по которому мы будем слушать сервер. Во втором параметре передаются настройки сервера.
	http.ListenAndServe(":8080", nil)
}
