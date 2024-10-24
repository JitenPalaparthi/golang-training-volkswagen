package main

import (
	"encoding/json"
	"flag"
	"fmt"
	"io"
	"log"
	"net/http"
	"os"
)

var (
	PORT      string
	chStudent chan Student
)

func init() {
	chStudent = make(chan Student, 5)
}

func main() {
	//slice := os.Args
	PORT = os.Getenv("PORT")
	if PORT == "" {
		flag.StringVar(&PORT, "port", "8081", "-port=8081")
		flag.Parse() // parse all flags
	}

	str := `hello world! "how are you doing"`
	println(str)

	http.HandleFunc("/student", func(w http.ResponseWriter, r *http.Request) {
		if r.Method == http.MethodPost {
			bytes, err := io.ReadAll(r.Body)
			if err != nil {
				http.Error(w, err.Error(), 400)
				return
			}
			s := new(Student)

			println(string(bytes), "\n")

			err = json.Unmarshal(bytes, s)
			if err != nil {
				http.Error(w, err.Error(), 400)
				return
			}
			chStudent <- *s
			fmt.Println(s)

		} else {
			http.Error(w, "method not implemented", http.StatusMethodNotAllowed)
			return
		}

	})

	println("Server is up and running on the port:", PORT)
	go createFileChan(chStudent)
	if err := http.ListenAndServe(":"+PORT, nil); err != nil {
		log.Fatal(err)
	}

}

type Student struct {
	Id    int    `json:"id"`
	Name  string `json:"name"`
	Email string `json:"email"`
}

func createFile(name string, data []byte) error {
	f, err := os.Create(name)
	defer f.Close()
	if err != nil {
		return err
	}
	_, err = f.Write(data)
	if err != nil {
		return err
	}
	return nil
}

func createFileChan(chStudent <-chan Student) {

	for ch := range chStudent {
		bytes, err := json.Marshal(ch)
		if err != nil {
			log.Println(err)
		} else {
			f, err := os.Create(fmt.Sprintf("students/%v", ch.Id))
			defer f.Close()
			if err != nil {
				log.Println(err)
			}
			_, err = f.Write(bytes)
			if err != nil {
				log.Println(err)
			}
		}
	}
}

/*
curl --location 'localhost:60051/student' \
--header 'Content-Type: application/json' \
--data-raw '{
    "id":3,
    "name":"Priya",
    "email":"Priya@outlook.com"
}'
*/
