package pubsub

import (
	"fmt"
	"log"
	"strings"
)

const apiVersion = 1

type Scope string

const (
	ScopeEvent   Scope = "evt"
	ScopeLogging Scope = "log"
	ScopeService Scope = "svc"
	ScopeSystem  Scope = "sys"
)

type Action string

const (
	ActionCreate Action = "create"
	ActionDelete Action = "delete"
	ActionUpdate Action = "update"
)

// SubjectId represents a specific entity identifier or a NATS wildcard (* or >)
type SubjectId string

const (
	Any SubjectId = "*" // Matches any single token
	All SubjectId = ">" // Matches all subsequent tokens
)

type SubjectBuilder struct {
	tokens       []string
	isTerminated bool
}

func NewSubjectBuilder(scope Scope) *SubjectBuilder {
	return &SubjectBuilder{
		tokens: []string{fmt.Sprintf("v%d", apiVersion), string(scope)},
	}
}

func (b *SubjectBuilder) addEntity(entity string, id SubjectId) *SubjectBuilder {
	if b.isTerminated {
		log.Printf("SubjectBuilder: tried to add to a terminated subject")
		return b
	}
	b.tokens = append(b.tokens, fmt.Sprintf("%s.%s", entity, id))
	if id == All {
		b.isTerminated = true
	}
	return b
}

func (b *SubjectBuilder) Guild(id SubjectId) *SubjectBuilder {
	b.addEntity("guild", id)
	return b
}

func (b *SubjectBuilder) Channel(id SubjectId) *SubjectBuilder {
	b.addEntity("channel", id)
	return b
}

func (b *SubjectBuilder) Message(id SubjectId) *SubjectBuilder {
	b.addEntity("message", id)
	return b
}

func (b *SubjectBuilder) User(id SubjectId) *SubjectBuilder {
	b.addEntity("user", id)
	return b
}

func (b *SubjectBuilder) Role(id SubjectId) *SubjectBuilder {
	b.addEntity("role", id)
	return b
}

func (b *SubjectBuilder) Upload(id SubjectId) *SubjectBuilder {
	b.addEntity("upload", id)
	return b
}

func (b *SubjectBuilder) Relationship(id SubjectId) *SubjectBuilder {
	b.addEntity("relationship", id)
	return b
}

func (b *SubjectBuilder) Any(id SubjectId) *SubjectBuilder {
	b.addEntity("*", id)
	return b
}

func (b *SubjectBuilder) All() *SubjectBuilder {
	b.tokens = append(b.tokens, ">")
	b.isTerminated = true
	return b
}

func (b *SubjectBuilder) AddAction(action Action) *SubjectBuilder {
	if b.isTerminated {
		log.Printf("SubjectBuilder: tried to add to a terminated subject")
		return b
	}
	b.tokens = append(b.tokens, string(action))
	b.isTerminated = true
	return b
}

func (b *SubjectBuilder) Build() string {
	return strings.Join(b.tokens, ".")
}
