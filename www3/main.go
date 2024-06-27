package main

import (
	"fmt"
	"html/template"
	"net/http"
)

type User struct {
	Name                  string
	Age                   uint16
	Money                 int16
	Avg_grades, Happiness float64
	Hobbies               []string
}

func main() {
	HandleRequest()
}

func (u User) GetAllInfo() string {
	return fmt.Sprintf("User name is: %s. He is: %d he has money equal: %d", u.Name, u.Age, u.Money)
}

func (u *User) SetNewName(NewName string) {
	u.Name = NewName
}

func Homepage(w http.ResponseWriter, r *http.Request) {
	Bob := User{Name: "Bob", Age: 25, Money: -50, Avg_grades: 4.2, Happiness: 0.8, Hobbies: []string{"Football", "Scate", "Dance"}}
	//fmt.Fprint(w, "<b>Main Text</b>")
	tmpl, _ := template.ParseFiles("templates/Homepage.html")
	tmpl.Execute(w, Bob)
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
// 1. tmpl, _ := template.ParseFiles("templates/Homepage.html") - Загружает шаблон из html файла и сохраняет его в переменную tmpl (т.е. Парсит)
// 2. tmpl.Execute(w, Bob) - данные структуры Bob передаются в tmpl, внутри tmpl уже есть ключи такие как name, age и тд, и они сохраняются в переменную
// w что является http.ResponseWrited который передают готовые данные пользователю на странице сайта.
