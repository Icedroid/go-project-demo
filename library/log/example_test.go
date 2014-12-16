package log_test

import (
	"go-project-demo/library/log"
)

func Example() {
	name := "Leaf"

	log.Debug("My name is %v", name)
	log.Info("My name is %v", name)
	log.Error("My name is %v", name)
	// log.Fatal("My name is %v", name)

	logger, err := log.New("release", "./file.log")
	if err != nil {
		return
	}
	defer logger.Close()

	logger.Debug("will not print")
	logger.Info("My name is %v", name)

	log.Export(logger)

	log.Debug("will not print")
	log.Info("My name is %v", name)
}
