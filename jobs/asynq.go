package jobs

import (
	"context"
	"encoding/json"
	"fiber-asynq-app/database"
	"log"

	"github.com/google/uuid"
	"github.com/hibiken/asynq"
)

const TaskCreateUser = "task:create_user"

type CreateUserPayload struct {
	Name  string `json:"name"`
	Email string `json:"email"`
}

func EnqueueCreateUserTask(name, email string) error {
	payload, err := json.Marshal(CreateUserPayload{Name: name, Email: email})
	if err != nil {
		return err
	}

	client := asynq.NewClient(asynq.RedisClientOpt{Addr: "localhost:6379"})
	defer client.Close()

	task := asynq.NewTask(TaskCreateUser, payload)
	_, err = client.Enqueue(task)
	return err
}

func ProcessCreateUserTask(ctx context.Context, task *asynq.Task) error {
	var payload CreateUserPayload
	if err := json.Unmarshal(task.Payload(), &payload); err != nil {
		return err
	}

	_, err := database.DB.Exec(ctx, "INSERT INTO users (id, name, email) VALUES ($1, $2, $3)", uuid.New(), payload.Name, payload.Email)
	if err != nil {
		log.Println("Failed to insert user:", err)
		return err
	}

	log.Println("User created:", payload.Email)

	return nil
}
