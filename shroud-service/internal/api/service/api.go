package service

import (
	"fmt"
	servicev1 "services/internal/api/service/v1"
	"services/internal/comm/pubsub"

	"github.com/nats-io/nats.go"
)

const (
	ApiVersion1 = "v1"

	service              = ApiVersion1 + ".service"
	messages             = service + ".message"
	SubjectMessageNew    = messages + ".new"
	SubjectMessageUpdate = messages + ".update"
	SubjectMessageDelete = messages + ".delete"

	uploads             = service + ".uploads"
	SubjectUploadNew    = uploads + ".new"
	SubjectUploadDelete = uploads + ".delete"
)

const (
	event = ApiVersion1 + ".event"
)

func SubjectMessageCreated(guildId, channelId string) string {
	return fmt.Sprintf("%s.guild.%s.channel.%s.message.created", event, guildId, channelId)
}

func SubjectMessageUpdated(guildId, channelId string) string {
	return fmt.Sprintf("%s.guild.%s.channel.%s.message.updated", event, guildId, channelId)
}

func SubjectMessageDeleted(guildId, channelId string) string {
	return fmt.Sprintf("%s.guild.%s.channel.%s.message.deleted", event, guildId, channelId)
}

var (
	SubscriptionsDispatcher = map[string]nats.MsgHandler{
		SubjectUploadNew:     servicev1.NewUpload,
		SubjectUploadDelete:  servicev1.DeleteUpload,
		SubjectMessageNew:    servicev1.NewMessage,
		SubjectMessageUpdate: servicev1.UpdateMessage,
		SubjectMessageDelete: servicev1.DeleteMessage,
	}
	SubscriptionsScribe = map[string]nats.MsgHandler{}
)

func Subscribe(group map[string]nats.MsgHandler) error {
	for s, h := range group {
		_, err := pubsub.Connection.Subscribe(s, h)
		if err != nil {
			return err
		}
	}
	return nil
}
