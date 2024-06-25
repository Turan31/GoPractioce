package main

import (
	"fmt"
	"net/http"
)

type User struct {
	name                  string
	age                   uint16
	money                 int16
	avg_grades, happiness float64
}

func main() {
	HandleRequest()
}

func Homepage(w http.ResponseWriter, r *http.Request) {
	Bob := User{name: "Bob", age: 25, money: 50, avg_grades: 4.2, happiness: 0.8}
	Bob.SetNewName("Alex")
	fmt.Fprint(w, Bob.GetAllInfo())
}

func Contactspage(w http.ResponseWriter, r *http.Request) {
	fmt.Fprint(w, "My contacts!")
}

func HandleRequest() {
	//функция HandleFunc принимает 2 параметра. Первый это url адрес который мы хотим отслеживать. Второй это метод
	http.HandleFunc("/", Homepage)
	http.HandleFunc("/contacts/", Contactspage)
	// ListenAndServe принимает 2 параметра. Первый это порт, по которому мы будем слушать сервер. Во втором параметре передаются настройки сервера.
	http.ListenAndServe(":8080", nil)
}

func (u User) GetAllInfo() string {
	return fmt.Sprintf("User name is: %s. He is: %d he has money equal: %d", u.name, u.age, u.money)
}

func (u *User) SetNewName(NewName string) {
	u.name = NewName
}
