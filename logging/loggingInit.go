package logging

import (
	"log"
	"os"
)

var Logger *log.Logger

func init() {
	// f, err := os.OpenFile("/logs/error.log", os.O_RDWR|os.O_CREATE|os.O_APPEND, 0666)
	f, err := os.OpenFile("/logs/error.log",
		os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)
	if err != nil {
		log.Println(err)
	}
	defer f.Close()

	Logger = log.New(f, "Error", log.LstdFlags)
}
