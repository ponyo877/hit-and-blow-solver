package entity

import "fmt"

type Feedback struct {
	Hit  int
	Blow int
}

func NewFeedback(hit int, blow int) *Feedback {
	return &Feedback{hit, blow}
}

func (f Feedback) String() string {
	return fmt.Sprintf("%dH%dB", f.Hit, f.Blow)
}

func (f *Feedback) IncHit() {
	f.Hit++
}

func (f *Feedback) IncBlow() {
	f.Blow++
}
