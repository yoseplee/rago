package main

import (
	"fmt"

	"github.com/yoseplee/rago/infra/logger"
	v2 "github.com/yoseplee/rago/v2"
)

func main() {
	defer logger.SyncLogger()
	fmt.Println("hello world")
}

func ingest(ingester v2.Ingester) {
	if err := ingester.Ingest(); err != nil {
		panic(err)
	}
}
