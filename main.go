package main

import (
	log "github.com/sirupsen/logrus"

	"github.com/t-eckert/elevate/cmd"
)

func main() {
	log.SetLevel(log.InfoLevel)
	cmd.Execute()
}
