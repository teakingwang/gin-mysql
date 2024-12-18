package main

import (
	"github.com/teakingwang/gin-mysql/cmd/app"
	"math/rand"
	"os"
	"runtime"
	"time"
)

func main() {
	rand.Seed(time.Now().UTC().UnixNano())
	if len(os.Getenv("GOMAXPROCS")) == 0 {
		runtime.GOMAXPROCS(runtime.NumCPU())
	}
	command := app.NewServerCommand()

	if err := command.Execute(); err != nil {
		os.Exit(1)
	}
}
