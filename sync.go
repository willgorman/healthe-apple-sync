package main

import (
	"sort"
	"strings"
	"time"

	log "github.com/sirupsen/logrus"
	"github.com/willgorman/healthe-apple-sync/data"
	"github.com/willgorman/healthe-apple-sync/healthe"
)

type SyncSettings struct {
	startDate time.Time
	endDate   time.Time
	dryRun    bool
}

func SyncSteps(steps data.DailySteps, settings SyncSettings, store healthe.StepStore) error {
	sortedDates := steps.SortedKeys()
	start := settings.startDate.Format("2006-01-02")
	end := settings.endDate.Format("2006-01-02")

	filterDates(sortedDates, start, end)

	return nil
}

func filterDates(dates []string, start, end string) []string {
	chosenStart := sort.Search(len(dates), func(i int) bool { return strings.Compare(start, dates[i]) <= 0 })
	chosenEnd := sort.Search(len(dates), func(i int) bool { return strings.Compare(dates[i], end) > 0 })
	log.Infof("start: %d end: %d", chosenStart, chosenEnd)
	return dates[chosenStart:chosenEnd]
}
