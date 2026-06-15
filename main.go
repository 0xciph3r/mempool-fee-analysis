package main

import (
	"log"

	"github.com/0xciph3r/mempool-fee/app"
	mempoolfee "github.com/0xciph3r/mempool-fee/mempool-fee"
)

func main() {

	err := app.App().Listen("0.0.0.0:4000")

	if err != nil {
		log.Fatal(err)
	}

	fee := mempoolfee.GetBestFee()

	log.Println(fee)
}
