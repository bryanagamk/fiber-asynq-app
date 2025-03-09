package jobs

import (
	"log"

	"github.com/hibiken/asynq"
)

func StartWorker() {
	server := asynq.NewServer(
		asynq.RedisClientOpt{Addr: "localhost:6379"},
		asynq.Config{
			Concurrency: 50,
		},
	)

	mux := asynq.NewServeMux()
	mux.HandleFunc(TaskCreateUser, ProcessCreateUserTask)

	if err := server.Run(mux); err != nil {
		log.Fatalf("Could not start server: %v\n", err)
	}
}
