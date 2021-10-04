package command

import (
	_ "cnjob/internal/logger"
	log "github.com/sirupsen/logrus"
	"testing"
)

type exampleInstance struct {
	state int
	m     func()
}

func (e *exampleInstance) Initialize() error {
	log.Info("null instance initialize")
	return nil
}

func (e *exampleInstance) RunLoop() {
	log.Info("null run_loop...")
	//os.Exit(1)
}

func (e *exampleInstance) Destroy() {
	log.Info("null destroy...")
}

func TestRun(t *testing.T) {
	Run(&exampleInstance{})
}
