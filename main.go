package main

import (
	"flag"
	"time"

	"github.com/willgorman/healthe-apple-sync/data"

	log "github.com/sirupsen/logrus"
	"github.com/willgorman/healthe-apple-sync/healthe"
)

func main() {
	var dataFilePath, user, password, start, end string
	var dryRun bool
	flag.StringVar(&dataFilePath, "data", "", "Path to the HealthKit export data (XML)")
	flag.StringVar(&user, "user", "", "Healthe Username")
	flag.StringVar(&password, "password", "", "Healthe Password")
	flag.StringVar(&start, "start", "1900-01-01", "TODO: format")
	flag.StringVar(&end, "end", "2200-01-01", "TODO: style")
	flag.BoolVar(&dryRun, "dryrun", false, "Display sync results without applying")
	flag.Parse()

	if dataFilePath == "" {
		flag.Usage()
		log.Fatal("data is required")
	}

	if user == "" {
		log.Fatal("user is required")
	}

	if password == "" {
		log.Fatal("password is required")
	}

	startDate, err := time.Parse(DATE_FORMAT, start)
	if err != nil {
		log.Fatalf("Invalid format for -start.  Expected YYYY-MM-DD but was %v", start)
	}

	endDate, err := time.Parse(DATE_FORMAT, end)
	if err != nil {
		log.Fatalf("Invalid format for -end.  Expected YYYY-MM-DD but was %v", end)
	}

	settings := SyncSettings{
		startDate: startDate,
		endDate:   endDate,
		dryRun:    dryRun,
	}

	store, err := healthe.Login(user, password)
	if err != nil {
		log.Fatalf("Could not create Healthe data source: %v", err)
	}

	steps, err := data.ParseHealthKitExportXML(dataFilePath)
	if err != nil {
		log.Fatalf("Could not parse HealthKit data: %v", err)
	}

	err = SyncSteps(*steps, settings, store)
	if err != nil {
		log.Fatal(err)
	}

	return
}
