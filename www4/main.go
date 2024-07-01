package main

import (
	"fmt"
	"database/sql"

	_ "github.com/go-sql-driver/mysql"
)

func main() {
	fmt.Println("Работа с Mysql")
}





db, err := sql.Open("mysql", "root:yourpassword@tcp(127.0.0.1:3306)/mydatabase")


err = db.Ping



createTableQuery := `
    CREATE TABLE IF NOT EXISTS users (
        id INT AUTO_INCREMENT,
        username VARCHAR(255) NOT NULL,
        password VARCHAR(255) NOT NULL,
        PRIMARY KEY (id)
    );`




_, err = db.Exec(createTableQuery)