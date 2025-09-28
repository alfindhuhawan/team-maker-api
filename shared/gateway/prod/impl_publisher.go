package prod

import (
	"context"
	"team-maker-api/shared/driver"
	"team-maker-api/shared/infrastructure/logger"
	"team-maker-api/shared/infrastructure/messaging"
	"team-maker-api/shared/model/payload"
)

type MessagePublisherImpl struct {
	Log       logger.Logger
	Publisher messaging.Publisher
	AppData   driver.ApplicationData
}

func (r *MessagePublisherImpl) PublishMessage(ctx context.Context, delayInMS int, topic string, obj any) error {
	err := PublishMessageMain(ctx, r.Publisher, r.AppData, delayInMS, topic, obj)
	if err != nil {
		r.Log.Error(ctx, "rabbitmq error", err.Error())
	}
	return err
}

func PublishMessageMain(ctx context.Context, publisher messaging.Publisher, appData driver.ApplicationData, delayInMS int, topic string, obj any) error {
	err := publisher.Publish(topic, delayInMS, payload.Payload{
		Data:      obj,
		Publisher: appData,
		TraceID:   logger.GetTraceID(ctx),
	})

	return err
}
