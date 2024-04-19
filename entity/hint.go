package entity

type Hint struct {
	numbers  Numbers
	feedback *Feedback
}

func NewHint(n Numbers, f *Feedback) *Hint {
	return &Hint{n, f}
}
