package parser

import (
	"github.com/kasterism/astermule/pkg/dag"
	"github.com/sirupsen/logrus"
)

type Message struct {
	status string
	Data   string `json:"data"`
}

type ControlPlane struct {
	Fs    []func()
	Entry []chan<- Message
	Exit  []<-chan Message
}

type Parser interface {
	Parse(*dag.DAG) ControlPlane
}

var (
	logger *logrus.Logger
)

func SetLogger(log *logrus.Logger) {
	logger = log
}
