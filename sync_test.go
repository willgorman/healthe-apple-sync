package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestFilterDates(t *testing.T) {
	dates := []string{"2018-01-02", "2018-01-04", "2018-01-06", "2018-01-08"}

	actual := filterDates(dates, "2018-01-01", "2018-01-10")
	assert.Equal(t, dates, actual)

	actual = filterDates(dates, "2018-01-02", "2018-01-08")
	assert.Equal(t, dates, actual)

	actual = filterDates(dates, "2018-01-04", "2018-01-06")
	assert.Equal(t, []string{"2018-01-04", "2018-01-06"}, actual)

	actual = filterDates(dates, "2018-01-03", "2018-01-07")
	assert.Equal(t, []string{"2018-01-04", "2018-01-06"}, actual)
}
