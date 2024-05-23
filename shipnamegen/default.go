package shipnamegen

import (
	"fmt"
	"log"
	"os"
	"path/filepath"
	"runtime"
)

func getPackagePath() string {
	_, filename, _, _ := runtime.Caller(0)
	return filepath.Dir(filename)
}

func DefaultPhraseSet() PhraseSet {
	defaultFilename := filepath.Join(getPackagePath(), "default.json")
	data, err := os.ReadFile(defaultFilename)
	if err != nil {
		log.Fatal(fmt.Sprintf("Failed to read file %s", defaultFilename))
	}

	var phrases PhraseSet
	phrases, err = jsonToPhraseSet(data)

	return phrases
}
