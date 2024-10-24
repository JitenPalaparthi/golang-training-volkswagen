package helpers

import (
	"demo/models"
	"encoding/json"
	"fmt"
	"log"
	"os"
)

func CreateFileChan(chStudent <-chan models.Student) {
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
