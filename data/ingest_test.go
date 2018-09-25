package data

import (
	"fmt"
	"testing"
)

func Test(t *testing.T) {
	a, err := ParseHealthKitExportXML("/Users/Will/Downloads/apple_health_export/export.xml")
	if err != nil {
		panic(err)
	}
	// fmt.Println(a)
	fmt.Println(len(a.steps))

	for k, v := range a.steps {
		fmt.Printf("%s: %d\n", k, v)
	}
}
