package main

import (
	"fmt"
	"sort"

	"github.com/a-randon-lemurian/star-system-sim/shipnamegen"
)

type phrasePermuts struct {
	phrase  string
	permuts int64
}

type byPhraseName []phrasePermuts

func (p byPhraseName) Len() int           { return len(p) }
func (p byPhraseName) Less(i, j int) bool { return p[i].phrase < p[j].phrase }
func (p byPhraseName) Swap(i, j int)      { p[i], p[j] = p[j], p[i] }

func main() {
	dps := shipnamegen.DefaultPhraseSet()
	fmt.Printf("--- TOTAL POSSIBILITIES ---\n")
	fmt.Printf("%-20s   %-20s", "phrase", "possibilities")
	fmt.Printf("\n------------------------------------\n")

	permuts := make([]phrasePermuts, 0)

	for phrase := range dps.Phrases {
		permuts = append(permuts, phrasePermuts{
			phrase:  phrase,
			permuts: dps.Enumerate(phrase)})
	}

	sort.Sort(byPhraseName(permuts))

	for _, p := range permuts {
		fmt.Printf("%-20s   %-20d\n", p.phrase, p.permuts)
	}
}
