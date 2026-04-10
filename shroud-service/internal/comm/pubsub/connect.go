package pubsub

import (
	"github.com/nats-io/nats.go"
	"github.com/nats-io/nats.go/jetstream"
)

var Connection *nats.Conn
var JetStream *nats.JetStream

func Connect(uri string) error {
	Conn, err := nats.Connect(uri)
	JetStream, err := jetstream.New(Conn)
	Connection = Conn
	return err
}
