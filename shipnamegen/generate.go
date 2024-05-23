package shipnamegen

import (
	"math/rand"
)

// TODO: prevent infinite recursion

func (ps PhraseSet) GenerateString(name string) string {
	var out string = ""
	for _, v := range ps.GetPhrase(name).Fragments {
		if v.Words != nil {
			out += v.Words[rand.Intn(len(v.Words))]
		}
		if v.Phrase != nil {
			out += ps.GenerateString(v.Phrase[rand.Intn(len(v.Phrase))])
		}
	}

	return out
}
