package command

import (
	"errors"
	log "github.com/sirupsen/logrus"
	"math/rand"
	"os"
	"os/signal"
	"runtime"
	"syscall"
	"time"
)

var (
	GMain   MainInterface
	GSignal chan os.Signal
)

type MainInterface interface {
	Initialize() error
	RunLoop()
	Destroy()
}

func Run(inst MainInterface) {
	if inst == nil {
		panic(errors.New("inst is nil, exit"))
		return
	}

	rand.Seed(time.Now().UTC().UnixNano())
	runtime.GOMAXPROCS(runtime.NumCPU())

	log.Info("instance initialize...")
	err := inst.Initialize()
	log.Info("inited")
	if err != nil {
		panic(err)
		return
	}

	// global
	GMain = inst

	log.Info("instance run_loop...")
	go inst.RunLoop()

	GSignal = make(chan os.Signal, 1)
	signal.Notify(GSignal, syscall.SIGHUP, syscall.SIGQUIT, syscall.SIGTERM, syscall.SIGINT)
	for {
		s := <-GSignal
		log.Infof("get a signal %s", s.String())
		switch s {
		case syscall.SIGQUIT, syscall.SIGTERM, syscall.SIGINT:
			log.Info("instance exit...")
			inst.Destroy()
			time.Sleep(time.Second)
			return
		case syscall.SIGHUP:
		default:
			return
		}
	}
}
