package main

import (
	"log"
	"os"
	"os/signal"
	"services/internal/comm/pubsub"
	"services/internal/config"
	"syscall"
)

func main() {
	environment, err := config.Load[config.AmbassadorEnvironment]()
	if err != nil {
		panic(err)
	}

	err = pubsub.Connect(environment.NatsUri)
	if err != nil {
		panic(err)
	}

	log.Println("Shroud Ambassador service started")

	sigchan := make(chan os.Signal, 1)
	signal.Notify(sigchan, syscall.SIGINT, syscall.SIGTERM)
	<-sigchan
	log.Println("Shroud Ambassador service shutting down...")

	log.Println("NATS connection draining")
	err = pubsub.Connection.Drain()
	if err != nil {
		panic(err)
	}
	log.Println("NATS connection drained")

	pubsub.Connection.Close()
	log.Println("NATS connection closed")

	log.Println("Shroud Ambassador service shut down")
}
