package prototype

import (
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
)

func TestKeywords_Clone(t *testing.T) {
	updateAt, _ := time.Parse("2006", "2020")
	words := Keywords{
		"testA": &Keyword{
			word:  "testA",
			visit: 1,
			updAt: &updateAt,
		},
		"testB": &Keyword{
			word:  "testB",
			visit: 2,
			updAt: &updateAt,
		},
		"testC": &Keyword{
			word:  "testC",
			visit: 3,
			updAt: &updateAt,
		},
	}

	now := time.Now()
	updatedWords := []*Keyword{
		{
			word:  "testB",
			visit: 10,
			updAt: &now,
		},
	}

	got := words.Clone(updatedWords)
	assert.Equal(t, words["testA"], got["testA"])
	assert.NotEqual(t, words["testB"], got["testB"])
	assert.NotEqual(t, updatedWords[0], got["testB"])
	assert.Equal(t, words["testC"], got["testC"])
}
