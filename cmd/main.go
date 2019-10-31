package main

import (
	"flag"
	"log"

	"github.com/ivanmatyash/check-minibus-seats/pkg/places"
)

func main() {
	var (
		dateFlag     = flag.String("date", "", "Requested date in the dd.mm.yyyy format.")
		routeFlag    = flag.Uint("route", 2, "Requested route number.")
		intervalFlag = flag.Uint("interval", 5, "Check interval (in seconds).")
	)
	flag.Parse()

	if err := places.ValidateDate(dateFlag); err != nil {
		log.Fatal(err.Error())
	}

	if err := places.CheckPlaces(*dateFlag, *routeFlag, *intervalFlag); err != nil {
		log.Fatal(err.Error())
	}

}
