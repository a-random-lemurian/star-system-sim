package shipnamegen

import (
	"encoding/json"
	"fmt"
	"log"
)

func makeStringSlice(interf interface{}) []string {
	processed := interf.([]interface{})
	var slice []string
	for _, word := range processed {
		slice = append(slice, word.(string))
	}
	return slice
}

func jsonParsePhrase(phrase interface{}) []PhraseFragment {
	phraseFrags, ok := phrase.([]interface{})
	if !ok {
		fatal("JSON error: Failed to parse phrase")
	}

	var pfSlice []PhraseFragment

	for _, v := range phraseFrags {
		var pf PhraseFragment
		m := v.(map[string]interface{})

		if m["word"] != nil && m["phrase"] != nil {
			fatal("JSON error: You may not specify both a word and phrase at the same time")
		}

		if m["word"] != nil {
			pf.Words = makeStringSlice(m["word"])
		}
		if m["phrase"] != nil {
			pf.Phrase = makeStringSlice(m["phrase"])
		}

		pfSlice = append(pfSlice, pf)
	}

	return pfSlice
}

/*
Convert a JSON phrase set file into a usable PhraseSet.
*/
func jsonToPhraseSet(jsonText []byte) (PhraseSet, error) {
	var data map[string]interface{}
	err := json.Unmarshal(jsonText, &data)
	if err != nil {
		log.Fatal(fmt.Sprintf("JSON parse error"))
	}

	var pset PhraseSet
	pset.Phrases = make(map[string]Phrase)

	rawPhrases, ok := data["phrases"].(map[string]interface{})
	if !ok {
		fatal("JSON error: the key \"phrases\" does not exist!")
	}
	for phrName, phrVal := range rawPhrases {
		var phrase Phrase
		phrase.Name = phrName
		phrase.Fragments = jsonParsePhrase(phrVal)
		pset.Phrases[phrName] = phrase
	}

	return pset, nil
}
