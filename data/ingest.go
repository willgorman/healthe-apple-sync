package data

import (
	"bufio"
	"encoding/xml"
	"os"
	"strings"
)

type DailySteps struct {
	steps map[string]int
}

type StepCount struct {
	CreationDate string `xml:"creationDate,attr"`
	StartDate    string `xml:"startDate,attr"`
	EndDate      string `xml:"endDate,attr"`
	Value        int    `xml:"value,attr"`
}

func ParseHealthKitExportXML(filePath string) (*DailySteps, error) {
	file, err := os.Open(filePath)
	if err != nil {
		return nil, err
	}
	fileReader := bufio.NewReader(file)
	decoder := xml.NewDecoder(fileReader)
	steps := DailySteps{
		steps: map[string]int{},
	}

	for {
		// Read tokens from the XML document in a stream.
		t, _ := decoder.Token()
		if t == nil {
			break
		}
		// Inspect the type of the token just read.
		switch se := t.(type) {
		case xml.StartElement:
			// If we just read a StartElement token
			// ...and its name is "page"
			for _, attr := range se.Attr {

				if attr.Value == "HKQuantityTypeIdentifierStepCount" {
					var sc StepCount
					// decode a whole chunk of following XML into the
					// variable p which is a Page (se above)
					decoder.DecodeElement(&sc, &se)
					// Do some stuff with the page.
					steps.steps[sc.getDay()] += sc.Value
				}
			}
		}
	}

	return &steps, nil
}

func (sc *StepCount) getDay() string {
	parts := strings.Split(sc.EndDate, " ")
	return parts[0]
}
