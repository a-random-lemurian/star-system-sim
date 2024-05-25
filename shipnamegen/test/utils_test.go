package shipnamegen_test

import (
	"runtime"
	"path/filepath"
)

func getPackagePath() string {
	_, filename, _, _ := runtime.Caller(0)
	return filepath.Dir(filename)
}

func getTestPhraseSet() string {
	return filepath.Join(getPackagePath(), "test.json")
}
