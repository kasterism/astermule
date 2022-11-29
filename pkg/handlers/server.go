package handlers

import (
	"errors"
	"net/http"
	"strconv"

	"github.com/kasterism/astermule/pkg/parser"
	"github.com/sirupsen/logrus"
)

const (
	formatBase = 10
)

var (
	logger       *logrus.Entry
	controlPlane *parser.ControlPlane

	ErrURLExisted = errors.New("url is already used")
)

func SetLogger(log *logrus.Entry) {
	logger = log
}

func StartServer(cp *parser.ControlPlane, address string, port uint, target string) error {
	http.HandleFunc(target, launchHandler)
	controlPlane = cp

	launchAllThread()

	listenAddr := address + ":" + strconv.FormatUint(uint64(port), formatBase)
	logger.Infoln("Start listening in", listenAddr)
	err := http.ListenAndServe(listenAddr, nil)
	if err != nil {
		logger.Fatalln("URL cannot listen:", err)
		return ErrURLExisted
	}
	return nil
}

// Notice! we don't care what the user sends! this handler is just a trigger that starts the process!
func launchHandler(w http.ResponseWriter, _ *http.Request) {
	beforeServerStart()
	w.Write(afterServerStart())
}

func launchAllThread() {
	for _, f := range controlPlane.Fs {
		go f()
	}
}

func beforeServerStart() {
	for i := range controlPlane.Entry {
		controlPlane.Entry[i] <- *parser.NewMessage(true, nil)
	}
}

func afterServerStart() []byte {
	res := parser.NewMessage(true, nil)
	for i := range controlPlane.Exit {
		msg := <-controlPlane.Exit[i]
		msg.DeepMergeInto(res)
	}
	data, err := res.Parse()
	if err != nil {
		logger.Errorln("Result message parse error:", err)
		errMsg := parser.NewMessage(false, nil)
		errData, _ := errMsg.Parse()
		return errData
	}
	return data
}
