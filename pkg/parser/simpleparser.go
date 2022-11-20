package parser

import (
	"github.com/kasterism/astermule/pkg/dag"
)

type SimpleParser struct {
	ChanGroup map[string]*ChannelGroup
}

type ChannelGroup struct {
	ReadCh  []<-chan Message
	WriteCh []chan<- Message
}

func NewSimpleParser() *SimpleParser {
	return &SimpleParser{
		ChanGroup: make(map[string]*ChannelGroup),
	}
}

func (s *SimpleParser) Parse(d *dag.DAG) ControlPlane {
	s.Init(d)
	s.makeChannelGroup(d)

}

func (s *SimpleParser) Init(d *dag.DAG) {
	for _, node := range d.Nodes {
		s.ChanGroup[node.Name] = &ChannelGroup{
			ReadCh:  make([]<-chan Message, 0),
			WriteCh: make([]chan<- Message, 0),
		}
	}
}

func (s *SimpleParser) makeChannelGroup(d *dag.DAG) {
	for _, node := range d.Nodes {
		for _, dep := range node.Dependencies {
			ch := make(chan Message)
			s.ChanGroup[dep].WriteCh = append(s.ChanGroup[dep].WriteCh, ch)
			s.ChanGroup[node.Name].ReadCh = append(s.ChanGroup[node.Name].ReadCh, ch)
		}
	}
}
