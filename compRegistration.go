package main

import "errors"

type CompId uint64

var currId CompId = 0
var registerdComps = make(map[string]CompId)

func nextCompId(label string) (CompId, error) {
	if _, ok := registerdComps[label]; ok {
		return 0, errors.New("Component already registered")
	}

	next := currId
	currId++
	if next > currId {
		panic("Component ID's have overflown")
	}

	return next, nil
}
