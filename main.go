package main

import (
	"binanceTemporal/activities"
	"binanceTemporal/client"
	"binanceTemporal/sqlite"
	workers "binanceTemporal/worker"
	"log"
)

func main() {
	DB, err := sqlite.New("data.db")
	if err != nil {
		log.Fatal(err)
	}
	activities.SetDB(DB)
	client, err := client.New()
	if err != nil {
		log.Fatal(err)
	}
	go workers.Start()
	go client.GetBinanceData()

	select {} // Block forever

}
