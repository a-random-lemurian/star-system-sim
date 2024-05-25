package shipnamegen_test

import (
	"fmt"
	"lemuria/spaceport/shipnamegen"
	"path/filepath"
	"runtime"
	"testing"
)

type testCase struct {
	phrase   string
	expected int
}

var testCases = []testCase{
	{phrase: "3 possibilities", expected: 3},
	{phrase: "3 possibilities with embedded phrase", expected: 3},
}

func getPackagePath() string {
	_, filename, _, _ := runtime.Caller(0)
	return filepath.Dir(filename)
}

func getTestPhraseSet() string {
	return filepath.Join(getPackagePath(), "test.json")
}

func TestEnumerate(t *testing.T) {
	phrase := shipnamegen.OpenPhraseFile(getTestPhraseSet())

	for _, tcase := range testCases {
		permuts := phrase.Enumerate(tcase.phrase);
		phraseTestResultString := fmt.Sprintf("phrase '%v': got %v - expected %v",
			tcase.phrase, permuts, tcase.expected)
		if permuts != int64(tcase.expected) {
			t.Fatalf("bad  -- %v", phraseTestResultString)
		} else {
			t.Logf(  "good -- %v", phraseTestResultString)
		}
	}
}
