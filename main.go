package main

import (
	"flag"

	"github.com/willgorman/healthe-apple-sync/data"
)

func main() {
	var dataFilePath string
	flag.StringVar(&dataFilePath, "data", "", "Path to the HealthKit export data (XML)")

	data.ParseHealthKitExportXML("")
}
