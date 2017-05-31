package cmd

import (
	log "github.com/sirupsen/logrus"
)

func failOnError(err error, errorDescription string) {
	if err != nil {
		log.WithError(err).Fatal(errorDescription)
	}
}
