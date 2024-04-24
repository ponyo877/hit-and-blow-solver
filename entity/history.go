package entity

import "reflect"

type History struct {
	solver  *Solver
	history []*Hint
}

func NewHistory(s *Solver) *History {
	return &History{s, []*Hint{}}
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
	for _, numbers := range h.solver.AllPatterns() {
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
	return NewHistgram(h.solver, e, h.Candidate()).Feedbacks()
}
