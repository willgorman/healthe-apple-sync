package web

import (
	"errors"
	"time"

	"github.com/headzoo/surf/agent"
	"github.com/headzoo/surf/browser"
	"gopkg.in/headzoo/surf.v1"
)

func login(user, password string) *browser.Browser {
	board := surf.NewBrowser()
	board.SetUserAgent(agent.Chrome())
	err := board.Open("https://healtheatcernerportal.cerner.com/dt/nutr/pedometerentry.asp")

	if err != nil {
		panic(err)
	}

	board.Click("a.oAuth")
	form, err := board.Form("form.signin--form")
	if err != nil {
		panic(err)
	}

	form.Input("login_username", user)
	form.Input("login_password", password)
	err = form.Submit()
	return board
}

func getSteps(board *browser.Browser, date *time.Time) int {
	return 0
}

func postSteps(board *browser.Browser, date *time.Time, steps int) error {
	err := board.Open("https://healtheatcernerportal.cerner.com/dt/nutr/pedometerentry.asp")
	if err != nil {
		panic(err)
	}

	// board.Form()

	return errors.New("what")
}
