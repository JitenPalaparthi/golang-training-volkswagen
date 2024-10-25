package main

import (
	"database/sql"
	"fmt"
	"log"

	_ "github.com/go-sql-driver/mysql"
)

func main() {
	dsn := `admin:admin@tcp(127.0.0.1:3306)/demodb`
	db, err := sql.Open("mysql", dsn)
	if err != nil {
		log.Fatal(err)
	}
	defer db.Close()

	err = db.Ping()
	if err != nil {
		log.Println(err)
	} else {
		log.Println("Successfully connected to the database")
	}

	usersTable := `
	CREATE TABLE IF NOT EXISTS users(
	id INT AUTO_INCREMENT,
	name VARCHAR(100) NOT NULL,
	email VARCHAR(100) NOT NULL,
	created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    PRIMARY KEY (id));`
	_, err = db.Exec(usersTable)
	if err != nil {
		log.Println(err)
	} else {
		log.Println("Table either existed or created")
	}

	insertUser := `INSERT INTO users(name,email)VALUES(?,?)`

	r1, err := db.Exec(insertUser, "Jiten", "Jitenp@outlook.com")
	if err != nil {
		log.Println("error in inserting record", err.Error())
	} else {
		log.Print("Record successfully inserted ")
		log.Println(r1.LastInsertId())
	}

	deleteUser := `DELETE from users where id=?`

	r1, err = db.Exec(deleteUser, 1)
	if err != nil {
		log.Println("error in inserting record", err.Error())
	} else {
		if n, _ := r1.RowsAffected(); n > 0 {
			log.Print("Record successfully deleted ")
		}
	}
	seletUsers := `SELECT id,name,email,created_at FROM users`
	rows, err := db.Query(seletUsers)
	if err != nil {
		log.Println(err)
	}
	defer rows.Close()

	for rows.Next() {
		var (
			id          int
			name, email string
			createdat   string
		)
		if err := rows.Scan(&id, &name, &email, &createdat); err != nil {
			log.Println(err.Error())
		} else {
			fmt.Printf("User-> Id:%d Name:%v Email: %v Created-At %v\n", id, name, email, createdat)
		}

	}

}

// standard pacakge -> sql
// thirdparty-pacakge ORM -> GORM
// third-payrty pacakges SQLX
