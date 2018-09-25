package healthe

import (
	"fmt"
	"log"
	"os"
	"testing"
	"time"

	_ "github.com/joho/godotenv/autoload"
)

func TestPostSteps(t *testing.T) {
	fmt.Println("Here we go")
	browser, err := Login(os.Getenv("HEALTHE_USER"), os.Getenv("HEALTHE_PASSWORD"))
	if err != nil {
		t.Fatal(err)
	}
	err = browser.PostSteps(time.Date(2018, time.August, 20, 0, 0, 0, 0, time.UTC), 6993)
	if err != nil {
		t.Fatal(err)
	}
	// err = browser.Open("https://healtheatcernerportal.cerner.com/dt/nutr/pedometerentry.asp")
	// if err != nil {
	// 	panic(err)
	// }

	// fmt.Println(browser.Body())

	// form, err := browser.Form("form#submitForm")
	// if err != nil {
	// 	panic(err)
	// }

	// form.Input("Date", "8/20/2018")
	// form.Input("FormMode", "UPDATE")
	// form.Input("Stride", "2.5")
	// form.Input("steps", "6993")
	// form.Input("StepHourSelect", "01:00 AM")

	// err = form.Submit()
	// if err != nil {
	// 	panic(err)
	// }

	// fmt.Println(form.Method())

	// // fmt.Println(browser.StatusCode())
	// fmt.Println(browser.Url())
	// fmt.Println(browser.Body())
}

func TestGetSteps(t *testing.T) {
	browser, err := Login(os.Getenv("HEALTHE_USER"), os.Getenv("HEALTHE_PASSWORD"))
	log.Printf("User %s", os.Getenv("HEALTHE_USER"))
	if err != nil {
		t.Fatal(err)
	}
	fmt.Println("logged in")
	date, _ := time.Parse("1/2/2006", "9/10/2017")
	steps, err := browser.GetSteps(date)
	if err != nil {
		panic(err)
	}
	fmt.Printf("%d steps\n", steps)
}

func TestTime(t *testing.T) {
	// n := time.Now()
	then, _ := time.Parse("1/2/2006", "9/10/2018")
	p := fmt.Println
	p(then.Format("1/2/06"))
}
