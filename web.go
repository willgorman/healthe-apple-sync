package web

import (
	"encoding/json"
	"fmt"
	"html"
	"net/url"
	"strconv"
	"time"

	"github.com/headzoo/surf/agent"
	"github.com/headzoo/surf/browser"
	"gopkg.in/headzoo/surf.v1"

	log "github.com/sirupsen/logrus"
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
	fmt.Println(board.Body())
	return board
}

func getSteps(board *browser.Browser, date *time.Time) (int, error) {
	params := url.Values{
		"action": []string{"LOAD_DAYSTEPS"},
		"bid":    []string{"5764"}, //<input id="UserID" type="hidden" value="5764" />
		"date":   []string{date.Format("1/2/2006")},
	}
	fmt.Println("trying")
	err := board.OpenForm("https://healtheatcernerportal.cerner.com/dt/nutr/PedometerEntryAjax.asp", params)
	if err != nil {
		return 0, err
	}

	// Hmm.  URL returns JSON but without the right content type.  (Not sure it'd be handled anyway even with since Surf is HTML focused)
	// Can't use selectors but maybe parse manually after converting html entities

	steps, err := parseSteps(html.UnescapeString(board.Body()))
	if err != nil {
		return 0, nil
	}

	log.Info(strconv.Itoa(steps) + " steps on " + date.String())

	return steps, nil
}

func postSteps(board *browser.Browser, date *time.Time, steps int) error {
	err := board.Open("https://healtheatcernerportal.cerner.com/dt/nutr/pedometerentry.asp")
	if err != nil {
		return err
	}

	form, err := board.Form("form#submitForm")
	if err != nil {
		return err
	}

	form.Input("Date", date.Format("1/2/2006"))
	form.Input("FormMode", "UPDATE")
	form.Input("Stride", "2.5")
	form.Input("steps", strconv.Itoa(steps))
	form.Input("StepHourSelect", "01:00 AM")

	err = form.Submit()
	if err != nil {
		return err
	}

	return nil
}

func parseSteps(jsonData string) (int, error) {
	rawData := map[string]interface{}{}
	err := json.Unmarshal([]byte(jsonData), &rawData)
	if err != nil {
		log.Warn(err)
		return 0, err
	}
	return strconv.Atoi(rawData["daysteps"].(string))
}
