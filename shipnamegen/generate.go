package shipnamegen

import (
	"fmt"
	"math/rand"
)

type stringGenRecursion struct {
	phraseCounts map[string]int
	totalRecursion int
	phrasePath []string
}

func (sgr *stringGenRecursion) createRecursionPath() string {
	str := ""
	for _, step := range sgr.phrasePath {
		str += "`"+ step + "`" + "->"
	}
	return str
}

func (sgr *stringGenRecursion) dieOfExcessiveRecursion() {
	panic(fmt.Sprintf("Excessive recursion!\nRecursion path: %v",
		sgr.createRecursionPath()))
}

func (sgr *stringGenRecursion) RegisterRecursion(name string) {
	if sgr.phraseCounts == nil {
		sgr.phraseCounts = make(map[string]int)
	}

	sgr.totalRecursion++
	sgr.phraseCounts[name]++
	sgr.phrasePath = append(sgr.phrasePath, name)

	// Recursion!
	if sgr.phraseCounts[name] >= 5 || sgr.totalRecursion > 200 {
		sgr.dieOfExcessiveRecursion()
	}
}

// Internal function, that keeps track of past phrases checked.
// External code starts the process at GenerateString(), which
// calls this function.
func (ps *PhraseSet) generateString(name string, sgr *stringGenRecursion) string {
	sgr.RegisterRecursion(name)
	var out string = ""
	for _, v := range ps.GetPhrase(name).Fragments {
		if v.Words != nil {
			out += v.Words[rand.Intn(len(v.Words))]
		}
		if v.Phrase != nil {
			out += ps.generateString(v.Phrase[rand.Intn(len(v.Phrase))], sgr)
		}
	}

	return out
}

// TODO: weighted phrases

// Generate a random string from a phrase.
func (ps *PhraseSet) GenerateString(name string) string {
	return ps.generateString(name, &stringGenRecursion{})
}
