package shipnamegen_test

import (
	"lemuria/spaceport/shipnamegen"
	"strings"
	"testing"
)

func TestRecursion(t *testing.T) {
	phrase := shipnamegen.OpenPhraseFile(getTestPhraseSet())

	defer func() {
		if r := recover(); r != nil {
			expectedPrefix := "Excessive recursion!"
			panicMsg := r.(string)
			if (!strings.HasPrefix(panicMsg, expectedPrefix)) {
				t.Errorf("Unexpected panic: %s", panicMsg)
			} else {
				t.Logf("Panic occurs as expected.\n")
			}
		}
	}()

	phrase.GenerateString("bad recursion")
}
