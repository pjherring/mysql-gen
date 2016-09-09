package util

import (
	"log"
	"strings"
)

func CompareLines(expected, actual string) bool {
	expectedSplit := strings.Split(strings.TrimSpace(expected), "\n")
	actualSplit := strings.Split(strings.TrimSpace(actual), "\n")

	for idx, p := range expectedSplit {

		//did we run out of lines
		if len(actualSplit) <= idx {
			log.Printf("actual is out of lines at %d", idx)
			return false
		}

		pTrim := strings.TrimSpace(p)
		aTrim := strings.TrimSpace(actualSplit[idx])

		//are the lines equal ignoring space
		if pTrim != aTrim {
			log.Printf("\"%s\" != \"%s\" at line %d", pTrim, aTrim, idx)
			return false
		}

	}

	return true
}
