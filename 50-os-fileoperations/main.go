package main

import (
	"os"
)

const (
	FILENAME = "data/data.txt"
)

func main() {
	f, err := os.OpenFile(FILENAME, os.O_APPEND|os.O_WRONLY, 0644)
	if err != nil {
		println(err.Error())
	}
	newData := "\nThis is demo for VolksWagen folks!"

	n, err := f.WriteString(newData)
	if err != nil {
		println(err.Error())
		return
	}
	println("data succssfully appended to file.", n, "bytes are written")
	f.Close()
}

func fatal(message string) {
	println(message)
	os.Exit(1)
}
