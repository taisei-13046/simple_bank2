package worker

import (
	"context"
	"encoding/json"
	"fmt"

	"github.com/hibiken/asynq"
	"github.com/rs/zerolog/log"
)

const TaskSendVerifyEmail = "task:send_verify_email"

type PayloadendVerifyEmail struct {
	Username string `json:"username"`
}

func (distributer *RedisTaskDistributer) DistributeTaskSendVerifyEmail(ctx context.Context, payload *PayloadendVerifyEmail, opt ...asynq.Option) error {
	jsonPayload, err := json.Marshal(payload)
	if err != nil {
		return fmt.Errorf("cannot marshal payload: %w", err)
	}
	task := asynq.NewTask(TaskSendVerifyEmail, jsonPayload, opt...)
	info, err := distributer.client.EnqueueContext(ctx, task)
	if err != nil {
		return fmt.Errorf("cannot enqueue task: %w", err)
	}

	log.Info().Msgf("task enqueued: %v", info)

	return nil
}

func (processor *RedisTaskProcessor) ProcessTaskSendVerifyEmail(ctx context.Context, task *asynq.Task) error {
	var payload PayloadendVerifyEmail
	if err := json.Unmarshal(task.Payload(), &payload); err != nil {
		return fmt.Errorf("cannot unmarshal payload: %w", err)
	}

	user, err := processor.store.GetUser(ctx, payload.Username)
	if err != nil {
		return fmt.Errorf("cannot get user: %w", err)
	}

	log.Info().Msgf("payload: %v", user)

	return nil
}
