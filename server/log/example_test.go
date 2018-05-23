package log_test

import (
	"github.com/rockamring/goproject/server/log"
	l"log"
)

func Example() {
	name := "go"

	log.Debug("hello world %v", name)
	log.Release("hello world %v", name)
	log.Error("hello world %v", name)

	logger, err := log.New("release", "", l.LstdFlags)
	if err != nil {
		return
	}
	defer logger.Close()

	logger.Debug("hello world")
	logger.Release("hello world")
}
