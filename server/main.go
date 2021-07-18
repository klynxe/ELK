package main

import (
	"log"
	"os"
	"time"

	"github.com/sirupsen/logrus"
	"go.elastic.co/ecslogrus"
)

func main() {
	file, err := os.OpenFile("./logs/server_log.log", os.O_WRONLY|os.O_CREATE|os.O_APPEND, 0755)
	if err != nil {
		log.Fatal(err)
	}
	var log = logrus.New()
	log.SetFormatter(&ecslogrus.Formatter{})
	log.SetOutput(file)
	logEntry := log.WithField("my_field", "value")
	for {
		time.Sleep(5 * time.Second)
		logEntry.Info("info")
	}
}
