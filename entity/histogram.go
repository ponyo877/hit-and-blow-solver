package entity

import (
	"math"
)

type Histogram struct {
	digit int
	mp    map[Feedback]int
}

func NewHistgram(e Numbers, c []Numbers) *Histogram {
	mp := map[Feedback]int{}
	for _, candidate := range c {
		f := candidate.Feedback(e)
		mp[*f]++
	}
	return &Histogram{len(e), mp}
}

func (h *Histogram) Feedbacks() []Feedback {
	fs := []Feedback{}
	for f := range h.mp {
		fs = append(fs, f)
	}
	return fs
}

func (h *Histogram) Entropy() float64 {
	var e float64
	for f, val := range h.mp {
		buf := 1.0
		if f == *NewFeedback(h.digit, 0) {
			buf = 0.95
		}
		absv := math.Abs(float64(val))
		e += buf * absv * math.Log(1+absv)
	}
	return e
}
