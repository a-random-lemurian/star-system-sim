package shipnamegen

// Enumerate the total number of possible strings generated by a Phrase.
//
// name string - Name of phrase to enumerate.
func (ps *PhraseSet) Enumerate(name string) int64 {
	phrase := ps.GetPhrase(name)
	possibilities := int64(1)

	for _, fragment := range phrase.Fragments {
		if fragment.Phrase != nil {
			for _, phrase := range fragment.Phrase {
				possibilities *= ps.Enumerate(phrase)
			}
		}
		if fragment.Words != nil {
			possibilities *= int64(len(fragment.Words))
		}
	}

	return possibilities
}
