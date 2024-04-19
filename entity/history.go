package entity

import "reflect"

type History struct {
	digit   int
	history []*Hint
}

func NewHistory(d int) *History {
	return &History{d, []*Hint{}}
}

func (h *History) Digit() int {
	return h.digit
}

func (h *History) IsEmpty() bool {
	return len(h.history) == 0
}

func (h *History) Push(hint *Hint) {
	h.history = append(h.history, hint)
}

func (h *History) Pop() {
	h.history = h.history[:len(h.history)-1]
}

func (h *History) Candidate() NumbersList {
	candidate := NumbersList{}
	for _, numbers := range AllNumbers(h.digit) {
		var isBreak bool
		for _, hint := range h.history {
			if !reflect.DeepEqual(hint.feedback, numbers.Feedback(hint.numbers)) {
				isBreak = true
				break
			}
		}
		if !isBreak {
			candidate = append(candidate, numbers)
		}
	}
	return candidate
}

func (h *History) FeedbackSelect(e Numbers) []Feedback {
	return NewHistgram(e, h.Candidate()).Feedbacks()
}
