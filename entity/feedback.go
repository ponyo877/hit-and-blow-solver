package entity

import "fmt"

type Feedback struct {
	Hit  int
	Blow int
}

func NewFeedback(h int, b int) *Feedback {
	return &Feedback{h, b}
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
