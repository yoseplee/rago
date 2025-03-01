package main

import (
	"github.com/yoseplee/rago"
	"github.com/yoseplee/rago/infra"
)

func main() {
	defer infra.Logger.Sync()
	retrieve()
}

func retrieve() {
	retriever := rago.DefaultRetriever{}
	if err := retriever.Retrieve(); err != nil {
		panic(err)
	}
}

func ingest() {
	ingester := rago.NewDefaultIngester(rago.JSONLoader{FilePath: "sample_shop_items_all.json"})
	if err := ingester.Ingest(); err != nil {
		panic(err)
	}
}
