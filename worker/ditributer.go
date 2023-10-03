package worker

import (
	"context"

	"github.com/hibiken/asynq"
)

type TaskDistributor interface {
	DistributeTaskSendVerifyEmail(ctx context.Context, payload *PayloadendVerifyEmail, opt ...asynq.Option) error
}

type RedisTaskDistributer struct {
	client *asynq.Client
}

func NewRedisTaskDistributer(redisOpt asynq.RedisClientOpt) TaskDistributor {
	client := asynq.NewClient(redisOpt)
	return &RedisTaskDistributer{
		client: client,
	}
}
