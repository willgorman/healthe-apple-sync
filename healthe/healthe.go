package healthe

import (
	"encoding/json"
	"errors"
	"html"
	"net/url"
	"strconv"
	"strings"
	"time"

	"github.com/headzoo/surf/agent"
	"github.com/headzoo/surf/browser"
	"gopkg.in/headzoo/surf.v1"

	log "github.com/sirupsen/logrus"
)

type StepStore interface {
	GetSteps(date time.Time) (int, error)
	PostSteps(date time.Time, steps int) error
}

type healtheStepStore struct {
	browser.Browser
}

func Login(user, password string) (StepStore, error) {
	board := surf.NewBrowser()
	board.SetUserAgent(agent.Chrome())
	err := board.Open("https://healtheatcernerportal.cerner.com/dt/nutr/pedometerentry.asp")

	if err != nil {
		panic(err)
	}

	board.Click("a.oAuth")
	form, err := board.Form("form.signin--form")
	if err != nil {
		return nil, err
	}

	form.Input("login_username", user)
	form.Input("login_password", password)
	err = form.Submit()
	if err != nil {
		return nil, err
	}

	// Login errors return a 200 but with an html error panel
	alertError := board.Find(".alert-error")
	if alertError.Length() > 0 {
		log.Errorf("User/password error ")
		return nil, errors.New(strings.Trim(alertError.Text(), " \n"))
	}
	// fmt.Println(board.Body())
	return healtheStepStore{*board}, nil
}

func (store healtheStepStore) GetSteps(date time.Time) (int, error) {
	params := url.Values{
		"action": []string{"LOAD_DAYSTEPS"},
		"bid":    []string{"5764"}, //<input id="UserID" type="hidden" value="5764" />
		"date":   []string{date.Format("1/2/2006")},
	}

	err := store.OpenForm("https://healtheatcernerportal.cerner.com/dt/nutr/PedometerEntryAjax.asp", params)
	if err != nil {
		log.Errorf("Failed to open form: %s", store.Body())
		return 0, err
	}

	// Hmm.  URL returns JSON but without the right content type.  (Not sure it'd be handled anyway even with since Surf is HTML focused)
	// Can't use selectors but maybe parse manually after converting html entities

	steps, err := parseSteps(html.UnescapeString(store.Body()))
	if err != nil {
		log.Errorf("Failed to parse steps from: %s", store.Body())
		return 0, err
	}

	log.Info(strconv.Itoa(steps) + " steps on " + date.String())

	return steps, nil
}

func (store healtheStepStore) PostSteps(date time.Time, steps int) error {
	err := store.Open("https://healtheatcernerportal.cerner.com/dt/nutr/pedometerentry.asp")
	if err != nil {
		return err
	}

	form, err := store.Form("form#submitForm")
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
