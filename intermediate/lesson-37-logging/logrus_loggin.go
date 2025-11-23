package main

import "github.com/sirupsen/logrus"

func main() {
	log := logrus.New()

	// set log level
	log.SetLevel(logrus.InfoLevel)

	// Set log format
	log.SetFormatter(&logrus.JSONFormatter{})

	log.Info("This is an info message")
	log.Warn("This is s a warninig message.")
	log.Error("This is an error message. ")

	log.WithFields(logrus.Fields{
		"username": "John Doe",
		"method":   "GET",
	}).Info("User logged in.")
}
