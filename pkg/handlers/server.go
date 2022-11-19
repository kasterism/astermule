package handlers

import (
	"errors"
	"net/http"
	"strconv"

	"github.com/sirupsen/logrus"
)

const (
	formatBase = 10
)

var (
	logger *logrus.Logger

	ErrURLExisted = errors.New("url is already used")
)

func SetLogger(log *logrus.Logger) {
	logger = log
}

func StartServer(address string, port uint, target string) error {
	http.HandleFunc(target, entryHandler)
	listenAddr := address + ":" + strconv.FormatUint(uint64(port), formatBase)
	logger.Infoln("Start listening in", listenAddr)
	err := http.ListenAndServe(listenAddr, nil)
	if err != nil {
		logger.Fatalln("URL cannot listen:", err)
		return ErrURLExisted
	}
	return nil
}

func entryHandler(w http.ResponseWriter, r *http.Request) {
}
