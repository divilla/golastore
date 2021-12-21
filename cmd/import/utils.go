package main

import (
	"github.com/rs/zerolog/log"
)

func fatal(err error) {
	if err != nil {
		panic(err)
	}
}

func warn(err error) {
	if err != nil {
		log.Warn().
			Err(err).
			Send()
	}
}

func info(err error) {
	if err != nil {
		log.Info().
			Err(err).
			Send()
	}
}
