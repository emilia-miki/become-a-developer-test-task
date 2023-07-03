package uniqueness

import (
	"errors"
	"math"
)

type uniqueness struct {
	ordinal  int
	isUnique bool
}

func newUniqueness(ordinal int) *uniqueness {
	return &uniqueness{
		ordinal:  ordinal,
		isUnique: true,
	}
}

type uniquenessMap struct {
	data    map[rune]*uniqueness
	counter int
}

func NewMap() *uniquenessMap {
	var m uniquenessMap
	m.data = make(map[rune]*uniqueness)
	return &m
}

func (um *uniquenessMap) Check(r rune) {
	u, ok := um.data[r]
	if !ok {
		um.data[r] = newUniqueness(um.counter)
		um.counter += 1
	} else {
		u.isUnique = false
	}
}

func (um *uniquenessMap) GetFirstUniqueRune() (rune, error) {
	if len(um.data) == 0 {
		return 0, errors.New("The map is empty")
	}

	var minRune rune
	minOrdinal := math.MaxInt
	for k, v := range um.data {
		if !v.isUnique {
			continue
		}

		if v.ordinal < minOrdinal {
			minOrdinal = v.ordinal
			minRune = k
		}
	}

	if minOrdinal == math.MaxInt {
		return 0, errors.New("All the runes are not unique")
	}

	return minRune, nil
}
