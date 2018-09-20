package web

import (
	"fmt"
	"os"
	"testing"
	"time"

	_ "github.com/joho/godotenv/autoload"
)

func TestGet(t *testing.T) {
	fmt.Println("Here we go")
	browser := login(os.Getenv("HEALTHE_USER"), os.Getenv("HEALTHE_PASSWORD"))
	err := browser.Open("https://healtheatcernerportal.cerner.com/dt/nutr/pedometerentry.asp")
	if err != nil {
		panic(err)
	}

	fmt.Println(browser.Body())

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

func TestGetSteps(t *testing.T) {
	browser := login(os.Getenv("HEALTHE_USER"), os.Getenv("HEALTHE_PASSWORD"))
	fmt.Println("logged in")
	date, _ := time.Parse("1/2/2006", "9/10/2017")
	steps, _ := getSteps(browser, &date)
	fmt.Printf("%d steps\n", steps)
}

func TestTime(t *testing.T) {
	// n := time.Now()
	then, _ := time.Parse("1/2/2006", "9/10/2018")
	p := fmt.Println
	p(then.Format("1/2/06"))
}
