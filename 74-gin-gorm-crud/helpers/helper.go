package helpers

import (
	"demo/models"
	"encoding/json"
	"log"
	"os"
	"runtime"
)

func DoAudit(filename string, chAudit <-chan models.Audit) {
	f, err := os.OpenFile(filename, os.O_APPEND|os.O_WRONLY, 0644)
	if err != nil {
		log.Println(err)
		runtime.Goexit()
	}
	defer f.Close()

	for ch := range chAudit {
		bytes, err := json.Marshal(ch)
		if err != nil {
			log.Println(err)
		} else {
			_, err = f.WriteString(string(bytes) + "\n")
			if err != nil {
				log.Println(err)
			}
		}
	}
}
