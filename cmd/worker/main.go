package main

import (
	"fiber-asynq-app/database"
	"fiber-asynq-app/jobs"
	"log"
)

func main() {
	log.Println("Starting worker...")

	database.ConnectDB()

	jobs.StartWorker()
}
