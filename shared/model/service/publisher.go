package service

import "context"

type PublishMessageService interface {
	PublishMessage(ctx context.Context, delayInMS int, topic string, obj any) error
}
