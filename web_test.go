package web

import (
	"fmt"
	"testing"
)

func TestGet(t *testing.T) {
	fmt.Println("Here we go")
	browser := login("foo", "bar")
	err := browser.Open("https://healtheatcernerportal.cerner.com/dt/nutr/pedometerentry.asp")
	if err != nil {
		panic(err)
	}

	form, err := browser.Form("form#submitForm")
	if err != nil {
		panic(err)
	}

	form.Input("Date", "8/20/2018")
	form.Input("FormMode", "UPDATE")
	form.Input("Stride", "2.5")
	form.Input("steps", "6993")
	form.Input("StepHourSelect", "01:00 AM")

	err = form.Submit()
	if err != nil {
		panic(err)
	}

	fmt.Println(form.Method())

	// fmt.Println(browser.StatusCode())
	fmt.Println(browser.Url())
	fmt.Println(browser.Body())
}
