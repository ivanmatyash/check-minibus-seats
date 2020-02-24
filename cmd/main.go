package main

import (
	"flag"
	"log"

	"github.com/ivanmatyash/check-minibus-seats/pkg/places"
)

func main() {
	var (
		dateFlag       = flag.String("date", "", "Requested date in the dd.mm.yyyy format.")
		routeFlag      = flag.Uint("route", 2, "Requested route number (1 Minsk - Mozyr, 2 Mozyr - Minsk).")
		intervalFlag   = flag.Uint("interval", 5, "Check interval (in seconds).")
		idCityFromFlag = flag.Uint("city-from", 1, "ID city from (1 - Minsk, 3 - Mozyr).")
		idCityToFlag   = flag.Uint("city-to", 3, "ID city from.")
		timeFromFlag   = flag.Uint("time-from", 0, "Time from")
		timeToFlag     = flag.Uint("time-to", 23, "Time to")
	)
	flag.Parse()

	if err := places.ValidateDate(dateFlag); err != nil {
		log.Fatal(err.Error())
	}

	if err := places.CheckPlaces(*dateFlag, *routeFlag, *intervalFlag, *idCityFromFlag, *idCityToFlag, *timeFromFlag, *timeToFlag); err != nil {
		log.Fatal(err.Error())
	}

}
