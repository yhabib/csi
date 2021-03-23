package protein

import (
	"errors"
)

var (
	// ErrStop occurs when there is a STOP codon
	ErrStop error = errors.New("stop")
	// ErrInvalidBase occurs where this not a valid codon
	ErrInvalidBase error = errors.New("not valid codon")
)

var translations = map[string]string{
	"AUG": "Methionine",
	"UUU": "Phenylalanine",
	"UUC": "Phenylalanine",
	"UUA": "Leucine",
	"UUG": "Leucine",
	"UCU": "Serine",
	"UCC": "Serine",
	"UCA": "Serine",
	"UCG": "Serine",
	"UAU": "Tyrosine",
	"UAC": "Tyrosine",
	"UGU": "Cysteine",
	"UGC": "Cysteine",
	"UGG": "Tryptophan",
	"UAA": "STOP",
	"UAG": "STOP",
	"UGA": "STOP",
}

// FromCodon
func FromCodon(s string) (string, error) {
	v, ok := translations[s]
	if !ok {
		return "", ErrInvalidBase
	}
	if v == "STOP" {
		return "", ErrStop
	}
	return v, nil
}

// FromRNA
func FromRNA(s string) (proteins []string, err error) {
	for i := 0; i < len(s)-2; i += 3 {
		protein, err := FromCodon(s[i : i+3])
		if err == ErrInvalidBase {
			return proteins, err
		}
		if err == ErrStop {
			break
		}
		proteins = append(proteins, protein)
	}

	return
}
