package main

import (
	"sort"
	"strings"
	"time"

	"github.com/hashicorp/go-multierror"
	log "github.com/sirupsen/logrus"
	"github.com/willgorman/healthe-apple-sync/data"
	"github.com/willgorman/healthe-apple-sync/healthe"
)

const DATE_FORMAT = "2006-01-02"

type SyncSettings struct {
	startDate time.Time
	endDate   time.Time
	dryRun    bool
}

func SyncSteps(steps data.DailySteps, settings SyncSettings, store healthe.StepStore) error {
	var result error
	sortedDates := steps.SortedKeys()
	start := settings.startDate.Format(DATE_FORMAT)
	end := settings.endDate.Format(DATE_FORMAT)

	dateRange := filterDates(sortedDates, start, end)
	for _, date := range dateRange {
		dateTime, _ := time.Parse(DATE_FORMAT, date)
		stepsOnDate, err := store.GetSteps(dateTime)
		if err != nil {
			log.Errorf("Failed to get steps on %v.  [%v]", date, err)
			result = multierror.Append(result, err)
			continue
		}

		if stepsOnDate > 0 {
			log.Infof("Skipping %v because steps are already recorded", date)
			continue
		}

		if settings.dryRun {
			log.Infof("[dryRun] Uploading %v steps for %v", steps.StepsOnDate(date), date)
		} else {
			err = store.PostSteps(dateTime, steps.StepsOnDate(date))
			if err != nil {
				log.Errorf("Failed to upload steps for %v. Error: %v", date, err)
				result = multierror.Append(result, err)
			} else {
				log.Infof("Uploaded %v steps for %v", steps.StepsOnDate(date), date)
			}
		}

	}

	return result
}

func filterDates(dates []string, start, end string) []string {
	chosenStart := sort.Search(len(dates), func(i int) bool { return strings.Compare(start, dates[i]) <= 0 })
	chosenEnd := sort.Search(len(dates), func(i int) bool { return strings.Compare(dates[i], end) > 0 })
	log.Infof("start: %d end: %d", chosenStart, chosenEnd)
	return dates[chosenStart:chosenEnd]
}
