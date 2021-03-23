import "strings"

func devow(s string) string {
	return strings.Map(func(r rune) rune {
		if r == 'a' || r == 'e' || r == 'i' || r == 'o' || r == 'u' {
			return '*'
		}
		return r
	}, s)
}

func spellchecker(wordlist []string, queries []string) (list []string) {
	set := make(map[string]bool)
	lowCap, noVowels := make(map[string]string), make(map[string]string)

	for _, w := range wordlist {
		set[w] = true

		lowW := strings.ToLower(w)
		if _, ok := lowCap[lowW]; !ok {
			lowCap[lowW] = w
		}

		noVow := devow(lowW)
		if _, ok := noVowels[noVow]; !ok {
			noVowels[noVow] = w
		}
	}
	for _, q := range queries {
		lowQ := strings.ToLower(q)
		noVow := devow(lowQ)
		if set[q] {
			list = append(list, q)
		} else if _, ok := lowCap[lowQ]; ok {
			list = append(list, lowCap[lowQ])
		} else if _, ok := noVowels[noVow]; ok {
			list = append(list, noVowels[noVow])
		} else {
			list = append(list, "")
		}
	}
	return list
}