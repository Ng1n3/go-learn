package main

import (
	"log"
	"os"
)

func main() {
	log.Println("This is a log message.")

	log.SetPrefix("INFO: ")
	log.Println("This is an info message")

	// Log Flags
	log.SetFlags(log.Ldate | log.Ltime | log.Lshortfile)
	log.Println("This is a log message with date, time.")

	log.SetFlags(log.Ldate | log.Ltime | log.Llongfile)
	log.Println("This is a log message with date, time, file.")

	infoLogger.Println("This is an info message")
	warnLogger.Println("This is a warning message")
	errorLogger.Println("This is an error message")

	// writing to a file and not terminal
	file, err := os.OpenFile("app.log", os.O_CREATE|os.O_WRONLY|os.O_APPEND, 0666)
	if err != nil {
		log.Fatalln("failed to  open log file: ", err)
	}

	defer file.Close()

	debugLogger := log.New(file, "INFO: ", log.Ldate|log.Ltime|log.Lshortfile)
	debugLogger.Println("This is a debug message.")

	infoLogger = log.New(file, "INFO: ", log.Ldate|log.Ltime|log.Lshortfile)
	warnLogger = log.New(file, "WARN: ", log.Ldate|log.Ltime|log.Lshortfile)
	errorLogger = log.New(file, "ERROR: ", log.Ldate|log.Ltime|log.Lshortfile)

	infoLogger.Println("This is a log message.")
	warnLogger.Println("This is a warning message.")
	errorLogger.Println("This is an error message.")
}

var (
	infoLogger  = log.New(os.Stdout, "INFO: ", log.Ldate|log.Ltime|log.Lshortfile)
	warnLogger  = log.New(os.Stdout, "WARN: ", log.Ldate|log.Ltime|log.Lshortfile)
	errorLogger = log.New(os.Stdout, "ERROR: ", log.Ldate|log.Ltime|log.Lshortfile)
)
