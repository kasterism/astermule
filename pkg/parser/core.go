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
	Status Status `json:"status"`
	Data   string `json:"data"`
}

// TODO: Define Status
type Status struct {
	Health bool
}

func NewMessage(health bool, data string) *Message {
	return &Message{
		Status: Status{
			Health: health,
		},
		Data: data,
	}
}

func (in *Message) DeepMergeInto(out *Message) {
	if !in.Status.Health {
		out.Status.Health = false
		return
	}

	if in.Data == "" {
		out.Status.Health = true
		return
	}

	if out.Data == "" {
		out.Data = in.Data
		out.Status.Health = true
		return
	}

	inData, err := in.Unmarshal()
	if err != nil {
		logger.Errorln("Unmarshal fail:", err)
	}

	outData, err := out.Unmarshal()
	if err != nil {
		logger.Errorln("Unmarshal fail:", err)
	}

	inDataMap, ok := inData.(map[string]interface{})
	if !ok {
		out.Status.Health = false
		return
	}

	outDataMap, ok := outData.(map[string]interface{})
	if !ok {
		out.Status.Health = false
		return
	}

	// Shallow merge
	for k, v := range inDataMap {
		if _, ok := outDataMap[k]; !ok {
			outDataMap[k] = v
		}
	}

	res, err := json.Marshal(outDataMap)
	if err != nil {
		out.Status.Health = false
		return
	}

	out.Status.Health = true
	out.Data = string(res)
}

func (m Message) Marshal() ([]byte, error) {
	return json.Marshal(m)
}

func (m Message) Unmarshal() (interface{}, error) {
	var data interface{}
	err := json.Unmarshal([]byte(m.Data), &data)
	if err != nil {
		return nil, err
	}
	return data, nil
}
