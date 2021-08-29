package prototype

import (
	"time"
)

type Cloneable interface {
	Clone() Cloneable
}

type Keyword struct {
	word  string
	visit int
	updAt *time.Time
}

func (k *Keyword) Clone() *Keyword {
	upd := time.Unix(k.updAt.Unix(), 0)
	newKwd := &Keyword{
		word:  k.word,
		visit: k.visit,
		updAt: &upd,
	}
	return newKwd
}

type Keywords map[string]*Keyword

func (words Keywords) Clone(updatedWords []*Keyword) Keywords {
	newKeywords := Keywords{}

	for k, v := range words {
		newKeywords[k] = v
	}
	for _, word := range updatedWords {
		newKeywords[word.word] = word.Clone()
	}

	return newKeywords
}
