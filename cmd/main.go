package main

import (
	"github.com/yoseplee/rago/config"
	"github.com/yoseplee/rago/v1"
)

func main() {

	retrieve()
}

func retrieve() {
	retriever := v1.DefaultRetriever{}
	if err := retriever.Retrieve(); err != nil {
		panic(err)
	}
}

func ingest() {
	ingester := v1.NewDefaultIngester(v1.JSONLoader{FilePath: config.Config.SampleFilePath})
	if err := ingester.Ingest(); err != nil {
		panic(err)
	}
}
