package shipnamegen

import (
	"fmt"
	"log"
	"runtime"
)

var defaultPhraseSetFile string = "default.json"

func fatal(errmsg string) {
	var buf [8192]byte
	strace := runtime.Stack(buf[:], false)
	log.Fatal(fmt.Sprintf("err: %s \n %s", errmsg, strace))
}

func (ps *PhraseSet) GetPhrase(name string) Phrase {
	return ps.Phrases[name]
}
