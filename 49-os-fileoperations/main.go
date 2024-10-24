package main

import (
	"io"
	"os"
)

const (
	FILENAME = "data/data.txt"
)

func main() {

	if err := os.Remove(FILENAME); err != nil {
		println("could not find the file or could not delete", err.Error())
	} else {
		println("successfully deleted the file")
	}

	f, err := os.Create(FILENAME)
	if err != nil {
		fatal(err.Error())
	}
	defer f.Close()

	data := "Hello world"

	n, err := f.Write([]byte(data))

	if err != nil {
		println(err.Error())
		return
	}
	println("data succssfully write to file.", n, "bytes are written")

	f1, err := os.Open(FILENAME)
	if err != nil {
		println(err.Error())
		return
	}
	defer f1.Close()
	println("Reading from the file,", FILENAME)
	bytes, err := io.ReadAll(f1)
	if err != nil {
		println(err.Error())
		return
	}
	println(string(bytes))
}

func fatal(message string) {
	println(message)
	os.Exit(1)
}
