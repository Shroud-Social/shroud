package v1

import (
	"encoding/json"
	"fmt"
	"log"
	"services/internal/domain/realm/upload"

	"github.com/nats-io/nats.go"
)

func NewUpload(m *nats.Msg) {
	var receipt upload.Receipt
	err := json.Unmarshal(m.Data, &receipt)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(receipt)
}

func DeleteUpload(m *nats.Msg) {

}

func NewMessage(m *nats.Msg) {

}

func UpdateMessage(m *nats.Msg) {

}

func DeleteMessage(m *nats.Msg) {

}
