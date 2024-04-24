package entity

import (
	"slices"
)

type Numbers []int

var allChache []Numbers = []Numbers{}

func (n Numbers) Feedback(estimate Numbers) *Feedback {
	f := NewFeedback(0, 0)
	for i, num := range estimate {
		if n[i] == num {
			f.IncHit()
			continue
		}
		if slices.Contains(n, num) {
			f.IncBlow()
		}
	}
	return f
}

func (n Numbers) Equals(other Numbers) bool {
	if len(n) != len(other) {
		return false
	}
	for i := range n {
		if n[i] != other[i] {
			return false
		}
	}
	return true
}

type NumbersList []Numbers

func (n NumbersList) Contains(target Numbers) bool {
	for _, num := range n {
		if num.Equals(target) {
			return true
		}
	}
	return false
}
