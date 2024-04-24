package entity

import (
	"math"
)

type Histogram struct {
	solver       *Solver
	feedbackbMap map[Feedback]int
}

func NewHistgram(s *Solver, e Numbers, c []Numbers) *Histogram {
	mp := map[Feedback]int{}
	for _, candidate := range c {
		f := candidate.Feedback(e)
		mp[*f]++
	}
	return &Histogram{s, mp}
}

func (h *Histogram) Feedbacks() []Feedback {
	fs := []Feedback{}
	for f := range h.feedbackbMap {
		fs = append(fs, f)
	}
	return fs
}

func (h *Histogram) Entropy() float64 {
	var e float64
	for f, val := range h.feedbackbMap {
		buf := 1.0
		if f == *NewFeedback(h.solver.Digit(), 0) {
			buf = 0.95
		}
		absv := math.Abs(float64(val))
		e += buf * absv * math.Log(1+absv)
	}
	return e
}
