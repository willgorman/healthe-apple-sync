package data

import (
	"fmt"
	"testing"
)

func Test(t *testing.T) {
	a, err := parseHealthKitExportXML("/Users/Will/Downloads/apple_health_export/export.xml")
	if err != nil {
		panic(err)
	}
	// fmt.Println(a)
	fmt.Println(len(a.steps))

	for k, _ := range a.steps {
		fmt.Println(k)
	}
}
