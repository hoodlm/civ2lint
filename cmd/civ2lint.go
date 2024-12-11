package main

import (
	"flag"

	"go.uber.org/zap"

	"github.com/leonsp-ai/civ2lint/lib"
)

func main() {
	var c lib.Config

	flag.StringVar(&c.Path, "path", ".", "Path to the game, mod, or scenario directory")
	flag.Parse()

	logger, _ := zap.NewProduction()
	defer logger.Sync() // flushes buffer, if any
	sugar := logger.Sugar()

	sugar.Info("Logger initialized")

	cl := lib.New(c, sugar)
	err := cl.Lint()
	if err != nil {
		panic(err)
	}
}
