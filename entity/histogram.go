package entity

import "math"

type Histogram struct {
	mp map[Feedback]int
}

func NewHistgram(e Numbers, c []Numbers) *Histogram {
	mp := map[Feedback]int{}
	for _, candidate := range c {
		f := candidate.Feedback(e)
		mp[*f]++
	}
	return &Histogram{mp}
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
	for _, val := range h.mp {
		absv := math.Abs(float64(val))
		e += absv * math.Log(1+absv)
	}
	return e
}
