package parser

import (
	"github.com/kasterism/astermule/pkg/clients/httpclient"
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
	c := ControlPlane{}
	s.Init(d)
	s.makeChannelGroup(d)
	c.Entry, c.Exit = s.scanChannelGroup(d)
	c.Fs = s.makeFunc(d)
	return c
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

func (s *SimpleParser) scanChannelGroup(d *dag.DAG) ([]chan<- Message, []<-chan Message) {
	entry := make([]chan<- Message, 0)
	exit := make([]<-chan Message, 0)
	for _, v := range s.ChanGroup {
		if len(v.ReadCh) == 0 {
			ch := make(chan Message)
			entry = append(entry, ch)
			v.ReadCh = append(v.ReadCh, ch)
		}

		if len(v.WriteCh) == 0 {
			ch := make(chan Message)
			exit = append(exit, ch)
			v.WriteCh = append(v.WriteCh, ch)
		}
	}
	return entry, exit
}

func (s *SimpleParser) makeFunc(d *dag.DAG) []func() {
	fs := make([]func(), 0)
	for i := range d.Nodes {
		node := d.Nodes[i]
		chGrp := s.ChanGroup[node.Name]
		f := func() {
			for {
				logger.Infoln("func register:", node.Name)
				msgs := make([]Message, 0)
				for _, readCh := range chGrp.ReadCh {
					msg := <-readCh
					msgs = append(msgs, msg)
				}

				logger.Infoln("func launch:", node.Name)

				// TODO: Check error
				mergeMsg := &Message{}
				for i := range msgs {
					msgs[i].DeepMergeInto(mergeMsg)
				}

				// Prepare sendMsg
				sendMsg := &Message{}
				sendMsg.Status.Health = true

				// Call http client
				logger.Infoln("send msg to", node.URL)
				res, err := httpclient.Send(node.Action, node.URL, mergeMsg.Data)
				if err != nil {
					logger.Errorln("httpclient error:", err)
					sendMsg.Status.Health = false
				} else {
					logger.Infoln("receive respense:", res)
					sendMsg.Data = res
				}

				for _, writeCh := range chGrp.WriteCh {
					writeCh <- *sendMsg
				}

				logger.Infoln("func end:", node.Name)
			}
		}
		fs = append(fs, f)
	}
	return fs
}
