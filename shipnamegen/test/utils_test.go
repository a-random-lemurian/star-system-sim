package shipnamegen_test

import (
	"path/filepath"
	"runtime"
)

func getPackagePath() string {
	_, filename, _, _ := runtime.Caller(0)
	return filepath.Dir(filename)
}

func getTestPhraseSet() string {
	return filepath.Join(getPackagePath(), "test.json")
}
