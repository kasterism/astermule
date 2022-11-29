package parser

import (
	"encoding/json"

	"github.com/kasterism/astermule/pkg/dag"
	"github.com/sirupsen/logrus"
)

var (
	logger *logrus.Entry
)

func SetLogger(log *logrus.Entry) {
	logger = log
}

type ControlPlane struct {
	Fs    []func()
	Entry []chan<- Message
	Exit  []<-chan Message
}

type Parser interface {
	Parse(*dag.DAG) ControlPlane
}

type Message struct {
	Status Status      `json:"status"`
	Data   interface{} `json:"data"`
}

// TODO: Define Status
type Status struct {
	Health bool
}

func NewMessage(health bool, data interface{}) *Message {
	return &Message{
		Status: Status{
			Health: health,
		},
		Data: data,
	}
}

func (in *Message) DeepMergeInto(out *Message) {
	// TODO: Merge Json
}

func (m Message) Parse() ([]byte, error) {
	return json.Marshal(m)
}
