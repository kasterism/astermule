package handlers

import (
	"net/http"
	"strconv"

	"github.com/sirupsen/logrus"
)

const (
	formatBase = 10
)

var (
	logger *logrus.Logger
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
		logger.Fatalln(err, "URL cannot listen")
		return ErrURLExisted
	}
	return nil
}

func entryHandler(w http.ResponseWriter, r *http.Request) {
}
