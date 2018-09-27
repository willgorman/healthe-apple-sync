package main

import (
	"flag"

	log "github.com/sirupsen/logrus"

	"github.com/willgorman/healthe-apple-sync/data"
)

func main() {
	var dataFilePath, user, password, start, end string
	var dryRun bool
	flag.StringVar(&dataFilePath, "data", "", "Path to the HealthKit export data (XML)")
	flag.StringVar(&user, "user", "", "Healthe Username")
	flag.StringVar(&password, "password", "", "Healthe Password")
	flag.StringVar(&start, "start_date", "", "TODO: format")
	flag.StringVar(&end, "end_date", "", "TODO: style")
	flag.BoolVar(&dryRun, "dryrun", false, "Display sync results without applying")

	if dataFilePath == "" {
		log.Fatal("dataFilePath is required")
	}

	if user == "" {
		log.Fatal("user is required")
	}

	if password == "" {
		log.Fatal("password is required")
	}

	data.ParseHealthKitExportXML("")
}
