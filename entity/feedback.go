package entity

import "fmt"

type Feedback struct {
	hit  int
	blow int
}

func NewFeedback(hit int, blow int) *Feedback {
	return &Feedback{hit, blow}
}

func (f Feedback) String() string {
	return fmt.Sprintf("%dH%dB", f.hit, f.blow)
}

func (f *Feedback) IncHit() {
	f.hit++
}

func (f *Feedback) IncBlow() {
	f.blow++
}
