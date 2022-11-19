package main

import (
	"encoding/json"
	"flag"
	"fmt"
	"os"
	"os/signal"
	"syscall"

	"github.com/kasterism/astermule/pkg/dag"
	"github.com/kasterism/astermule/pkg/handlers"
	"github.com/sirupsen/logrus"
)

var (
	logger = logrus.New()
)

func setSignal() {
	count := 1
	c := make(chan os.Signal, count)
	signal.Notify(c, os.Interrupt, syscall.SIGTERM)
	go func() {
		<-c
		logger.Infoln("Signal interrupt")
		quitJob()
		os.Exit(1)
	}()
}

func main() {
	var (
		address string
		port    uint
		target  string
		dagStr  string
	)

	flag.StringVar(&address, "address", "0.0.0.0", "The boot address of launching astermule.")
	flag.UintVar(&port, "port", 8080, "The boot port of launching astermule.")
	flag.StringVar(&target, "target", "/", "Path of the http service.")
	flag.StringVar(&dagStr, "dag", "{}", "Describe the directed acyclic graph that astermule needs to collect(JSON format).")

	flag.Parse()

	defer func() {
		logger.Infoln("Coredump clean...")
		quitJob()
	}()

	setSignal()
	setLogger()

	graph := dag.NewDAG()
	err := json.Unmarshal([]byte(dagStr), graph)
	if err != nil {
		logger.Fatalln("The dag is not canonical and cannot be resolved")
		return
	}
	err = graph.Preflight()
	if err != nil {
		logger.Errorln("Preflight errors:", err)
		return
	}

	err = handlers.StartServer(address, port, target)
	if err != nil {
		return
	}

	fmt.Println(graph)
}

func setLogger() {
	const logKey = "package"

	handlers.SetLogger(logger.WithField(logKey, "handlers").Logger)
	handlers.SetLogger(logger.WithField(logKey, "dag").Logger)
}

func quitJob() {
	logger.Infoln("Quit...")
}
