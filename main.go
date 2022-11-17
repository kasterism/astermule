package main

import (
	"flag"
	"fmt"
)

func main() {
	var (
		address string
		port    uint
		dag     string
	)

	flag.StringVar(&address, "address", "0.0.0.0", "The boot address of launching astermule.")
	flag.UintVar(&port, "port", 80, "The boot port of launching astermule.")
	flag.StringVar(&dag, "dag", "{}", "Describe the directed acyclic graph that astermule needs to collect(JSON format).")

	fmt.Println(address)
}
